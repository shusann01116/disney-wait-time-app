package core

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func getIgnoreAVUnexportedOptions() cmp.Options {
	return cmp.Options{
		cmpopts.IgnoreUnexported(types.AttributeValueMemberM{}),
		cmpopts.IgnoreUnexported(types.AttributeValueMemberN{}),
		cmpopts.IgnoreUnexported(types.AttributeValueMemberNS{}),
		cmpopts.IgnoreUnexported(types.AttributeValueMemberBOOL{}),
		cmpopts.IgnoreUnexported(types.AttributeValueMemberB{}),
		cmpopts.IgnoreUnexported(types.AttributeValueMemberBS{}),
		cmpopts.IgnoreUnexported(types.AttributeValueMemberL{}),
		cmpopts.IgnoreUnexported(types.AttributeValueMemberS{}),
		cmpopts.IgnoreUnexported(types.AttributeValueMemberSS{}),
		cmpopts.IgnoreUnexported(types.AttributeValueMemberNULL{}),
	}
}
