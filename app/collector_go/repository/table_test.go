package repository

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/awsdocs/aws-doc-sdk-examples/gov2/testtools"
	"github.com/disney-wait-time-app/app/collector_go/core"
	"github.com/disney-wait-time-app/app/collector_go/stubs"
	"github.com/stretchr/testify/assert"
)

func enterTest() (*testtools.AwsmStubber, *Table) {
	stubber := testtools.NewStubber()
	table := &Table{
		DynamoDBClient: dynamodb.NewFromConfig(*stubber.SdkConfig),
		TableName:      "test-table",
	}
	return stubber, table
}

func TestOperationBeforeSetClient(t *testing.T) {
	expected := errors.New(
		"DynamoDBClient is nil, set DynamoDBClient before calling AddFacility\nor use NewTable() to create Table instance",
	)
	err := Table{}.AddFacility(context.TODO(), &core.Facility{})
	assert.Error(t, expected, err)
}

func TestTable_AddFacility(t *testing.T) {
	t.Run(
		"NoError",
		func(t *testing.T) { AddFacility(nil, t) },
	)
	t.Run(
		"TestError",
		func(t *testing.T) { AddFacility(&testtools.StubError{Err: errors.New("test-error")}, t) },
	)
}

func AddFacility(raiseError *testtools.StubError, t *testing.T) {
	stubber, table := enterTest()

	facility := core.NewFacility(0, "hoge", "ホゲ")
	item, marshalErr := attributevalue.MarshalMap(facility)
	if marshalErr != nil {
		panic(marshalErr)
	}

	stubber.Add(stubs.StubAddItem(table.TableName, item, raiseError))

	err := table.AddFacility(context.Background(), facility)

	testtools.VerifyError(err, raiseError, t)
	testtools.ExitTest(stubber, t)
}

func TestTable_AddStandby(t *testing.T) {
	t.Run("NoError", func(t *testing.T) { AddStandby(nil, t) })
}

func AddStandby(raiseError *testtools.StubError, t *testing.T) {
	stubber, table := enterTest()

	standby := core.NewStandby(
		time.Now(),
		core.FacilityId(0),
		core.Status{Name: "Open"},
		10,
	)
	item, marshalErr := attributevalue.MarshalMap(standby)
	if marshalErr != nil {
		panic(marshalErr)
	}

	stubber.Add(stubs.StubAddItem(table.TableName, item, raiseError))

	err := table.AddStandby(context.Background(), standby)

	testtools.VerifyError(err, raiseError, t)
	testtools.ExitTest(stubber, t)
}

func TestTable_BatchAddStandby(t *testing.T) {
	t.Run(
		"NoError",
		func(t *testing.T) { AddStandbyBatch(nil, t) },
	)
	t.Run(
		"TestError",
		func(t *testing.T) {
			AddStandbyBatch(
				&testtools.StubError{Err: errors.New("test-error"), ContinueAfter: true},
				t,
			)
		},
	)
}

func AddStandbyBatch(raiseError *testtools.StubError, t *testing.T) {
	stubber, table := enterTest()

	var testData []core.Standby
	var inputRequests []types.WriteRequest
	for i := 0; i < 30; i++ {
		standby := core.NewStandby(
			time.Now(),
			core.FacilityId(i),
			core.Status{Name: "Open"},
			10,
		)
		testData = append(testData, standby)
		inputRequests = append(inputRequests, types.WriteRequest{
			PutRequest: &types.PutRequest{
				Item: map[string]types.AttributeValue{
					"PK": &types.AttributeValueMemberS{
						Value: *aws.String(fmt.Sprintf("FC#%v", standby.FacilityID)),
					},
					"SK": &types.AttributeValueMemberS{
						Value: *aws.String(fmt.Sprintf("ST#%v", standby.UpdateAt.Format(time.DateOnly))),
					},
					"Data": &types.AttributeValueMemberS{
						Value: *aws.String(standby.Status.Name),
					},
					"UpdateAt": &types.AttributeValueMemberN{
						Value: *aws.String(fmt.Sprint(standby.UpdateAt.Unix())),
					},
					"StandbyDuration": &types.AttributeValueMemberN{
						Value: *aws.String(fmt.Sprint(standby.StandbyDuration)),
					},
				},
			},
		})
	}

	stubber.Add(stubs.StubAddItemBatch(table.TableName, inputRequests[0:25], raiseError))
	stubber.Add(stubs.StubAddItemBatch(table.TableName, inputRequests[25:30], raiseError))

	count, err := table.AddStandbyBatch(context.Background(), testData, 200)

	testtools.VerifyError(err, raiseError, t)
	if raiseError == nil {
		if count != 30 {
			t.Errorf("Got %v items written, expected %v", 30, count)
		}
	}
	testtools.ExitTest(stubber, t)
}

func TestTable_UpdateLatestStandby(t *testing.T) {
	t.Run(
		"NoError",
		func(t *testing.T) { UpdateLatestStandby(nil, t) },
	)
	t.Run(
		"TestError",
		func(t *testing.T) { UpdateLatestStandby(&testtools.StubError{Err: errors.New("test-error")}, t) },
	)
}

func UpdateLatestStandby(raiseError *testtools.StubError, t *testing.T) {
	stubber, table := enterTest()

	now := time.Now()
	standby := core.NewStandby(
		now,
		core.FacilityId(0),
		core.Status{Name: "Open"},
		10,
	)

	sk := fmt.Sprintf("ST#%v", now.Format(time.DateOnly))
	duration := fmt.Sprint(standby.StandbyDuration)

	stubber.Add(
		stubs.StubUpdateLatestStandby(
			table.TableName,
			standby.GetKey(),
			sk,
			duration,
			raiseError,
		),
	)

	attr, err := table.UpdateLatestStandby(context.Background(), standby)

	testtools.VerifyError(err, raiseError, t)
	if raiseError == nil {
		if attr["Data"] != sk ||
			attr["StandbyDuration"] != duration {
			t.Errorf(
				"want: %v, %v(%T)\ngot: %v, %v(%T)",
				sk,
				duration,
				duration,
				attr["Data"],
				attr["StandbyDuration"],
				attr["StandbyDuration"],
			)
		}
	}
	testtools.ExitTest(stubber, t)
}
