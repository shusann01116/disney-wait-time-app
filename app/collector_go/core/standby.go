package core

import (
	"fmt"
	"reflect"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Standby struct {
	ID         StandbyId  `dynamodbav:"SK"`
	FacilityID FacilityId `dynamodbav:"PK"`
	Status     Status     `dynamodbav:"Data"`
	UpdateAt   time.Time  `dynamodbav:",unixtime,omitempty"`
	// StandbyDuration is the duration of standby in minutes
	StandbyDuration time.Duration `dynamodbav:"StandbyDuration,omitempty"`
}

func NewStandby(
	updateAt time.Time,
	facilityId FacilityId,
	status Status,
	standbyDuration time.Duration,
) Standby {
	return Standby{
		ID:              StandbyId(updateAt),
		FacilityID:      facilityId,
		Status:          status,
		UpdateAt:        updateAt,
		StandbyDuration: standbyDuration,
	}
}

func (s Standby) GetKey() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"PK": &types.AttributeValueMemberS{
			Value: fmt.Sprintf("FC#%d", s.FacilityID),
		},
		"SK": &types.AttributeValueMemberS{
			Value: *aws.String("LatestStatus"),
		},
	}
}

type StandbyId time.Time

func (s StandbyId) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return &types.AttributeValueMemberS{
		Value: fmt.Sprintf("ST#%v", time.Time(s).Format(time.DateOnly)),
	}, nil
}

func (s *StandbyId) UnmarshalDynamoDBAttributeValue(av types.AttributeValue) error {
	tv, ok := av.(*types.AttributeValueMemberS)
	if !ok {
		return &UnmarshalTypeError{
			Value: fmt.Sprintf("%T", av),
			Type:  reflect.TypeOf((*StandbyId)(nil)),
		}
	}

	// Split ST#<date> to <date>
	var val string
	_, err := fmt.Sscanf(tv.Value, "ST#%s", &val)
	if err != nil {
		return &UnmarshalTypeError{
			Value: fmt.Sprintf("%T", av),
			Type:  reflect.TypeOf((*StandbyId)(nil)),
		}
	}

	t, err := time.Parse(time.DateOnly, val)
	if err != nil {
		return err
	}

	*s = StandbyId(t)

	return nil
}

type Status struct {
	Name string
}

func (s Status) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return &types.AttributeValueMemberS{
		Value: s.Name,
	}, nil
}

func (s *Status) UnmarshalDynamoDBAttributeValue(av types.AttributeValue) error {
	tv, ok := av.(*types.AttributeValueMemberS)
	if !ok {
		return &UnmarshalTypeError{
			Value: fmt.Sprintf("%T", av),
			Type:  reflect.TypeOf((*Status)(nil)),
		}
	}

	*s = Status{Name: tv.Value}

	return nil
}
