package core

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/go-cmp/cmp"
)

func TestMarshalFacility(t *testing.T) {
	cases := map[string]struct {
		input    Facility
		expected types.AttributeValue
	}{
		"normal": {
			input: Facility{
				ID:              FacilityId(1),
				FacilityIndexId: FacilityId(1),
				SortKey:         "",
				Name:            "a",
				KanaName:        "kana",
			},
			expected: &types.AttributeValueMemberM{
				Value: map[string]types.AttributeValue{
					"PK": &types.AttributeValueMemberS{
						Value: "FC#1",
					},
					"SK": &types.AttributeValueMemberS{
						Value: "Name",
					},
					"Data": &types.AttributeValueMemberS{
						Value: "a",
					},
					"FacilityKanaName": &types.AttributeValueMemberS{
						Value: "kana",
					},
					"FacilityIndexId": &types.AttributeValueMemberS{
						Value: "FC#1",
					},
				},
			},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			actual, err := attributevalue.Marshal(c.input)
			if err != nil {
				t.Fatalf("failed to marshal: %v", err)
			}
			if diff := cmp.Diff(c.expected, actual, getIgnoreAVUnexportedOptions()); diff != "" {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}
		})
	}
}

func TestUnmarshalFacility(t *testing.T) {
	cases := map[string]struct {
		input    types.AttributeValue
		expected Facility
	}{
		"normal": {
			input: &types.AttributeValueMemberM{
				Value: map[string]types.AttributeValue{
					"PK": &types.AttributeValueMemberS{
						Value: fmt.Sprintf("FC#%d", uint(1)),
					},
					"SK": &types.AttributeValueMemberS{
						Value: "Name",
					},
					"Data": &types.AttributeValueMemberS{
						Value: "a",
					},
					"FacilityKanaName": &types.AttributeValueMemberS{
						Value: "kana",
					},
				},
			},
			expected: Facility{
				ID:       FacilityId(1),
				SortKey:  "SK",
				Name:     "a",
				KanaName: "kana",
			},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			var actual Facility
			err := attributevalue.Unmarshal(c.input, &actual)
			if err != nil {
				t.Fatalf("failed to unmarshal: %v", err)
			}
			if diff := cmp.Diff(c.expected, actual, getIgnoreAVUnexportedOptions()); diff != "" {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}
		})
	}
}

func TestMarshalFacilityId(t *testing.T) {
	cases := map[string]struct {
		input    FacilityId
		expected types.AttributeValue
	}{
		"normal": {
			input:    FacilityId(1),
			expected: &types.AttributeValueMemberS{Value: fmt.Sprintf("FC#%d", uint(1))},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			actual, err := attributevalue.Marshal(c.input)
			if err != nil {
				t.Fatalf("failed to marshal: %v", err)
			}
			if diff := cmp.Diff(c.expected, actual, getIgnoreAVUnexportedOptions()); diff != "" {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}
		})
	}
}

func TestUnmarshalFacilityId(t *testing.T) {
	cases := map[string]struct {
		input    types.AttributeValue
		expected FacilityId
	}{
		"normal": {
			input: &types.AttributeValueMemberS{
				Value: fmt.Sprintf("FC#%d", uint(1)),
			},
			expected: FacilityId(1),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			var actual FacilityId
			err := attributevalue.Unmarshal(c.input, &actual)
			if err != nil {
				t.Fatalf("failed to unmarshal: %v", err)
			}
			if diff := cmp.Diff(c.expected, actual, getIgnoreAVUnexportedOptions()); diff != "" {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}
		})
	}
}

func TestMarshalFacilitySortKey(t *testing.T) {
	cases := map[string]struct {
		input    FacilitySortKey
		expected types.AttributeValue
	}{
		"normal": {
			input:    FacilitySortKey("Name"),
			expected: &types.AttributeValueMemberS{Value: "Name"},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			actual, err := attributevalue.Marshal(c.input)
			if err != nil {
				t.Fatalf("failed to marshal: %v", err)
			}
			if diff := cmp.Diff(c.expected, actual, getIgnoreAVUnexportedOptions()); diff != "" {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}
		})
	}
}

func TestUnmarshalFacilitySortKey(t *testing.T) {
	cases := map[string]struct {
		input    types.AttributeValue
		expected FacilitySortKey
	}{
		"normal": {
			input:    &types.AttributeValueMemberS{Value: "Name"},
			expected: FacilitySortKey("SK"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			var actual FacilitySortKey
			err := attributevalue.Unmarshal(c.input, &actual)
			if err != nil {
				t.Fatalf("failed to unmarshal: %v", err)
			}
			if diff := cmp.Diff(c.expected, actual, getIgnoreAVUnexportedOptions()); diff != "" {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}
		})
	}
}
