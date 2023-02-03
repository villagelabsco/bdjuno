package utils

import (
	"context"
	"fmt"
	"github.com/gogo/protobuf/proto"
	juno "github.com/villagelabsco/juno/v4/types"
	"strconv"

	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"
	"google.golang.org/grpc/metadata"
)

// RemoveDuplicateValues removes the duplicated values from the given slice
func RemoveDuplicateValues(slice []string) []string {
	keys := make(map[string]bool)
	var list []string

	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// GetHeightRequestContext adds the height to the context for queries
func GetHeightRequestContext(context context.Context, height int64) context.Context {
	return metadata.AppendToOutgoingContext(
		context,
		grpctypes.GRPCBlockHeightHeader,
		strconv.FormatInt(height, 10),
	)
}

func ProtoMsgName(msg proto.Message) string {
	return "/" + proto.MessageName(msg)
}

func FindEventAndAttr(index int, tx *juno.Tx, event proto.Message, attrKey string) (string, error) {
	evt, err := tx.FindEventByType(index, ProtoMsgName(event))
	if err != nil {
		return "", fmt.Errorf("error while finding event %s: %s", ProtoMsgName(event), err)
	}
	res, err := tx.FindAttributeByKey(evt, attrKey)
	if err != nil {
		return "", fmt.Errorf("error while finding %s attribute in evt %s: %s", attrKey, ProtoMsgName(event), err)
	}
	return res, nil
}
