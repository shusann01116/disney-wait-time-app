package core

import (
	"fmt"
	"reflect"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type OperatingHour struct {
	ID         OperatingHourId `dynamodbav:"SK"`
	FacilityID FacilityId      `dynamodbav:"PK"`
	From       time.Time       `dynamodbav:",unixtime,omitempty"`
	To         time.Time       `dynamodbav:",unixtime,omitempty"`
}

type OperatingHourId time.Time

func (o OperatingHourId) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return &types.AttributeValueMemberS{
		Value: fmt.Sprintf("OH#%v", time.Time(o).Format(time.DateOnly)),
	}, nil
}

func (o *OperatingHourId) UnmarshalDynamoDBAttributeValue(av types.AttributeValue) error {
	tv, ok := av.(*types.AttributeValueMemberS)
	if !ok {
		return &UnmarshalTypeError{
			Value: fmt.Sprintf("%T", av),
			Type:  reflect.TypeOf((*OperatingHourId)(nil)),
		}
	}

	// Split OH#<date> to <date>
	var val string
	_, err := fmt.Sscanf(tv.Value, "OH#%s", val)
	if err != nil {
		return err
	}

	t, err := time.Parse(time.DateOnly, val)
	if err != nil {
		return err
	}

	*o = OperatingHourId(t)

	return nil
}
