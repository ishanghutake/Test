// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.5.1
// source: greetpb/greet.proto

package greet

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_greetpb_greet_proto protoreflect.FileDescriptor

var file_greetpb_greet_proto_rawDesc = []byte{
	0x0a, 0x13, 0x67, 0x72, 0x65, 0x65, 0x74, 0x70, 0x62, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x72, 0x65, 0x65, 0x74, 0x32, 0x0e, 0x0a, 0x0c,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x09, 0x5a, 0x07,
	0x2e, 0x3b, 0x67, 0x72, 0x65, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_greetpb_greet_proto_goTypes = []interface{}{}
var file_greetpb_greet_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_greetpb_greet_proto_init() }
func file_greetpb_greet_proto_init() {
	if File_greetpb_greet_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_greetpb_greet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_greetpb_greet_proto_goTypes,
		DependencyIndexes: file_greetpb_greet_proto_depIdxs,
	}.Build()
	File_greetpb_greet_proto = out.File
	file_greetpb_greet_proto_rawDesc = nil
	file_greetpb_greet_proto_goTypes = nil
	file_greetpb_greet_proto_depIdxs = nil
}
