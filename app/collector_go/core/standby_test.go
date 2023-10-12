package core

import (
	"fmt"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/go-cmp/cmp"
)

func TestMarshalStandby(t *testing.T) {
	now := time.Now().UTC()
	cases := map[string]struct {
		input    Standby
		expected types.AttributeValue
	}{
		"normal": {
			input: Standby{
				ID:              StandbyId(now),
				FacilityID:      FacilityId(1),
				Status:          Status{Name: "a"},
				UpdateAt:        now,
				StandbyDuration: 5,
			},
			expected: &types.AttributeValueMemberM{
				Value: map[string]types.AttributeValue{
					"PK": &types.AttributeValueMemberS{
						Value: "FC#1",
					},
					"SK": &types.AttributeValueMemberS{
						Value: fmt.Sprintf("ST#%v", now.Format(time.DateOnly)),
					},
					"Data": &types.AttributeValueMemberS{
						Value: "a",
					},
					"UpdateAt": &types.AttributeValueMemberN{
						Value: fmt.Sprintf("%d", now.Unix()),
					},
					"StandbyDuration": &types.AttributeValueMemberN{
						Value: "5",
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
