package flatten

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"

	"github.com/sinsay/grpc-gateway/internal/descriptor"
	"github.com/sinsay/grpc-gateway/protoc-gen-openapiv2-fx/internal/flatten/meta"
)

type FlatInfo struct {
	Rule    *meta.FlattenRules
	Message *protogen.Message
	Field   *protogen.Field
}

var ProtoPlugin *protogen.Plugin = nil

func ExtractDefFromPlugin(msg *descriptor.Message, f *descriptor.Field) (message *protogen.Message, field *protogen.Field) {
	message, field = nil, nil

	if msg == nil || f == nil {
		return
	}

	file, ok := ProtoPlugin.FilesByPath[msg.File.FileDescriptorProto.GetName()]
	if !ok {
		return
	}

	for _, m := range file.Messages {
		if m.GoIdent.GoName == msg.GetName() {
			message = m
			break
		}
	}

	if message == nil {
		return
	}

	for _, ff := range message.Fields {
		if string(ff.Desc.Name()) == f.GetName() {
			field = ff
			break
		}
	}

	return
}

func ExtractFlatten(msg *descriptor.Message, f *descriptor.Field) (*FlatInfo, bool) {

	message, field := ExtractDefFromPlugin(msg, f)

	if message == nil || field == nil {
		return nil, false
	}

	flatten := proto.GetExtension(field.Desc.Options(), meta.E_Flatten).(bool)
	ext := proto.GetExtension(field.Desc.Options(), meta.E_FlattenRule).(*meta.FlattenRules)

	if ext == nil {
		// create default ext
		m := int32(0)
		ext = &meta.FlattenRules{Reserved: &meta.Reserved{
			Min: &m, // useless now
			Max: &m, // useless now
		}}
	} else {
		flatten = true
	}

	return &FlatInfo{
		Rule:    ext,
		Message: message,
		Field:   field,
	}, flatten
}
