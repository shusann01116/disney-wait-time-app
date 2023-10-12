package repository

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/disney-wait-time-app/app/collector_go/core"
	"github.com/rs/zerolog/log"
)

var logger = log.With().Str("pkg", "core").Logger()

type Table struct {
	DynamoDBClient *dynamodb.Client
	TableName      string
}

func (t Table) AddFacility(ctx context.Context, f *core.Facility) error {
	logger.Info().Msgf("AddFacility: %v", f)

	if t.DynamoDBClient == nil {
		return errors.New(
			"DynamoDBClient is nil, set DynamoDBClient before calling AddFacility\nor use NewTable() to create Table instance",
		)
	}

	if f.ID != f.FacilityIndexId {
		return &core.ErrInvalidFacility{Facility: *f}
	}

	item, err := attributevalue.MarshalMap(*f)
	if err != nil {
		return err
	}

	_, err = t.DynamoDBClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &t.TableName,
		Item:      item,
	})
	if err != nil {
		return err
	}
	return nil
}

func (t Table) AddStandbyBatch(
	ctx context.Context,
	standby []core.Standby,
	maxItems int,
) (int, error) {
	var err error
	var item map[string]types.AttributeValue
	written := 0
	batchSize := 25 // DynamoDB allows a maximum batch size of 25 items.
	start := 0
	end := start + batchSize
	for start < maxItems && start < len(standby) {
		var writeReqs []types.WriteRequest
		if end > len(standby) {
			end = len(standby)
		}
		for _, s := range standby[start:end] {
			item, err = attributevalue.MarshalMap(s)
			if err != nil {
				log.Printf(
					"Couldn't marshal standby %v for batch writing. Here's why: %v\n",
					s.ID,
					err,
				)
			} else {
				writeReqs = append(
					writeReqs,
					types.WriteRequest{PutRequest: &types.PutRequest{Item: item}},
				)
			}
		}
		_, err = t.DynamoDBClient.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
			RequestItems: map[string][]types.WriteRequest{t.TableName: writeReqs},
		})
		if err != nil {
			log.Printf(
				"Couldn't add a batch of standby to %v. Here's why: %v\n",
				t.TableName,
				err,
			)
		} else {
			written += len(writeReqs)
		}
		start = end
		end += batchSize
	}

	return written, err
}

func (t Table) AddStandby(ctx context.Context, standby core.Standby) error {
	log.Info().Msgf("AddStandby: %v", standby)

	item, err := attributevalue.MarshalMap(standby)
	if err != nil {
		return err
	}

	log.Debug().Msgf("item: %v", item)
	_, err = t.DynamoDBClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &t.TableName,
		Item:      item,
	})
	if err != nil {
		return err
	}
	return nil
}

func (t Table) UpdateLatestStandby(
	ctx context.Context,
	standby core.Standby,
) (map[string]interface{}, error) {
	update := expression.Set(expression.Name("Data"), expression.Value(standby.ID))
	update.Set(expression.Name("StandbyDuration"), expression.Value(standby.StandbyDuration))
	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		return nil, err
	}

	output, err := t.DynamoDBClient.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:                 &t.TableName,
		Key:                       standby.GetKey(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
		ReturnValues:              types.ReturnValueUpdatedNew,
	})
	if err != nil {
		return nil, err
	}

	var attMap map[string]interface{}
	err = attributevalue.UnmarshalMap(output.Attributes, &attMap)
	if err != nil {
		return nil, err
	}

	return attMap, nil
}
