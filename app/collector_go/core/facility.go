package core

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Facility struct {
	ID              FacilityId      `dynamodbav:"PK"`
	FacilityIndexId FacilityId      `dynamodbav:"FacilityIndexId"`
	SortKey         FacilitySortKey `dynamodbav:"SK"`
	Name            string          `dynamodbav:"Data"`
	KanaName        string          `dynamodbav:"FacilityKanaName,omitempty"`
}

type (
	FacilityId      uint
	FacilitySortKey string
)

func NewFacility(id uint, name string, kanaName string) *Facility {
	return &Facility{
		ID:              FacilityId(id),
		FacilityIndexId: FacilityId(id),
		SortKey:         FacilitySortKey("SK"),
		Name:            name,
		KanaName:        kanaName,
	}
}

func (f FacilityId) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return &types.AttributeValueMemberS{
		Value: fmt.Sprintf("FC#%d", f),
	}, nil
}

func (f *FacilityId) UnmarshalDynamoDBAttributeValue(av types.AttributeValue) error {
	tv, ok := av.(*types.AttributeValueMemberS)
	if !ok {
		return &UnmarshalTypeError{
			Value: fmt.Sprintf("%T", av),
			Type:  reflect.TypeOf((*FacilityId)(nil)),
		}
	}

	// Split FC#<id> to <id>
	var val string
	_, err := fmt.Sscanf(tv.Value, "FC#%s", &val)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return err
	}

	*f = FacilityId(id)

	return nil
}

func (f FacilitySortKey) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return &types.AttributeValueMemberS{
		Value: "Name",
	}, nil
}

func (f *FacilitySortKey) UnmarshalDynamoDBAttributeValue(av types.AttributeValue) error {
	*f = FacilitySortKey("SK")
	return nil
}
