package stubs

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/awsdocs/aws-doc-sdk-examples/gov2/testtools"
)

func StubAddItem(
	tableName string,
	item map[string]types.AttributeValue,
	raiseErr *testtools.StubError,
) testtools.Stub {
	return testtools.Stub{
		OperationName: "PutItem",
		Input:         &dynamodb.PutItemInput{TableName: aws.String(tableName), Item: item},
		Output:        &dynamodb.PutItemOutput{},
		Error:         raiseErr,
	}
}

func StubAddItemBatch(
	tableName string,
	inputRequests []types.WriteRequest,
	raiseErr *testtools.StubError,
) testtools.Stub {
	return testtools.Stub{
		OperationName: "BatchWriteItem",
		Input: &dynamodb.BatchWriteItemInput{
			RequestItems: map[string][]types.WriteRequest{
				tableName: inputRequests,
			},
		},
		Output: &dynamodb.BatchWriteItemOutput{},
		Error:  raiseErr,
	}
}

func StubUpdateLatestStandby(
	tableName string,
	key map[string]types.AttributeValue,
	data string,
	standbyDurationStr string,
	raiseErr *testtools.StubError,
) testtools.Stub {
	return testtools.Stub{
		OperationName: "UpdateItem",
		Input: &dynamodb.UpdateItemInput{
			TableName:                aws.String(tableName),
			Key:                      key,
			ExpressionAttributeNames: map[string]string{"#0": "Data", "#1": "StandbyDuration"},
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":0": &types.AttributeValueMemberS{Value: data},
				":1": &types.AttributeValueMemberN{
					Value: standbyDurationStr,
				},
			},
			UpdateExpression: aws.String("SET #0 = :0, #1 = :1\n"),
			ReturnValues:     types.ReturnValueUpdatedNew,
		},
		Output: &dynamodb.UpdateItemOutput{
			Attributes: map[string]types.AttributeValue{
				"Data": &types.AttributeValueMemberS{Value: data},
				"StandbyDuration": &types.AttributeValueMemberS{
					Value: standbyDurationStr,
				},
			},
		},
		Error: raiseErr,
	}
}
