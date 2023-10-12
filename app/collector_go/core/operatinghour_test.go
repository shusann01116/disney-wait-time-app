package core

import (
	"fmt"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/go-cmp/cmp"
)

func TestMarshalOperatingHour(t *testing.T) {
	now := time.Now().UTC()

	cases := map[string]struct {
		input    OperatingHour
		expected types.AttributeValue
	}{
		"empty": {
			input: OperatingHour{},
			expected: &types.AttributeValueMemberM{
				Value: map[string]types.AttributeValue{
					"PK": &types.AttributeValueMemberS{
						Value: "FC#0",
					},
					"SK": &types.AttributeValueMemberS{
						Value: "OH#0001-01-01",
					},
					"From": &types.AttributeValueMemberN{
						// Unix time min
						Value: "-62135596800",
					},
					"To": &types.AttributeValueMemberN{
						// Unix time min
						Value: "-62135596800",
					},
				},
			},
		},
		"normal": {
			input: OperatingHour{
				ID:         OperatingHourId{},
				FacilityID: FacilityId(1),
				From:       now,
				To:         now,
			},
			expected: &types.AttributeValueMemberM{
				Value: map[string]types.AttributeValue{
					"PK": &types.AttributeValueMemberS{
						Value: "FC#1",
					},
					"SK": &types.AttributeValueMemberS{
						Value: "OH#0001-01-01",
					},
					"From": &types.AttributeValueMemberN{
						// Unix time min
						Value: fmt.Sprintf("%d", now.Unix()),
					},
					"To": &types.AttributeValueMemberN{
						// Unix time min
						Value: fmt.Sprintf("%d", now.Unix()),
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
