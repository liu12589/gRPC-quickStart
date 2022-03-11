package serializer

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func ProtobufToJSON(message proto.Message) (string, error) {
	marshaler := jsonpb.Marshaler{
		// 使用默认值填入字段
		EmitDefaults: false,
		// 是否将枚举些微证书或者字符串
		EnumsAsInts: false,
		// 使用什么缩进
		Indent: "   ",
		// 是否使用原始字段
		OrigName: true,
	}
	return marshaler.MarshalToString(message)
}
