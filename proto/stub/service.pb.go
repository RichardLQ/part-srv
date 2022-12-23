// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.5
// source: service.proto

package stub

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x73, 0x74, 0x75, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0xbc, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x58, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x74, 0x50, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x74, 0x75, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x48,
	0x6f, 0x74, 0x50, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e,
	0x73, 0x74, 0x75, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x52, 0x73, 0x70, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x68, 0x6f, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x53, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x74,
	0x75, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x14, 0x2e, 0x73, 0x74, 0x75, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x52, 0x73, 0x70, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22,
	0x0d, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x3a, 0x01,
	0x2a, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x73, 0x74, 0x75, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_service_proto_goTypes = []interface{}{
	(*GetHotPartTimeReq)(nil), // 0: stub.GetHotPartTimeReq
	(*GetPartTimeReq)(nil),    // 1: stub.GetPartTimeReq
	(*GetPartTimeRsp)(nil),    // 2: stub.GetPartTimeRsp
}
var file_service_proto_depIdxs = []int32{
	0, // 0: stub.UserService.GetHotPartTime:input_type -> stub.GetHotPartTimeReq
	1, // 1: stub.UserService.GetPartTime:input_type -> stub.GetPartTimeReq
	2, // 2: stub.UserService.GetHotPartTime:output_type -> stub.GetPartTimeRsp
	2, // 3: stub.UserService.GetPartTime:output_type -> stub.GetPartTimeRsp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	file_partime_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}