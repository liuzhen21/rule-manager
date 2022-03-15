// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: api/helloworld/v1/error.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// @plugins=protoc-gen-go-errors
// 错误
type Error int32

const (
	// 未知类型
	// @code=UNKNOWN
	Error_RESOURCE_ERR_UNKNOWN Error = 0
	// 找不到资源
	// @code=NOT_FOUND
	Error_RESOURCE_ERR_NOT_FOUND Error = 1
	// 获取列表数据出错
	// @code=INTERNAL
	Error_RESOURCE_ERR_LIST_MYSQL Error = 2
	// 获取详情数据出错
	// @code=INTERNAL
	Error_RESOURCE_ERR_INFO_MYSQL Error = 3
)

// Enum value maps for Error.
var (
	Error_name = map[int32]string{
		0: "RESOURCE_ERR_UNKNOWN",
		1: "RESOURCE_ERR_NOT_FOUND",
		2: "RESOURCE_ERR_LIST_MYSQL",
		3: "RESOURCE_ERR_INFO_MYSQL",
	}
	Error_value = map[string]int32{
		"RESOURCE_ERR_UNKNOWN":    0,
		"RESOURCE_ERR_NOT_FOUND":  1,
		"RESOURCE_ERR_LIST_MYSQL": 2,
		"RESOURCE_ERR_INFO_MYSQL": 3,
	}
)

func (x Error) Enum() *Error {
	p := new(Error)
	*p = x
	return p
}

func (x Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Error) Descriptor() protoreflect.EnumDescriptor {
	return file_api_helloworld_v1_error_proto_enumTypes[0].Descriptor()
}

func (Error) Type() protoreflect.EnumType {
	return &file_api_helloworld_v1_error_proto_enumTypes[0]
}

func (x Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Error.Descriptor instead.
func (Error) EnumDescriptor() ([]byte, []int) {
	return file_api_helloworld_v1_error_proto_rawDescGZIP(), []int{0}
}

var File_api_helloworld_v1_error_proto protoreflect.FileDescriptor

var file_api_helloworld_v1_error_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2a, 0x77,
	0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x45, 0x53, 0x4f, 0x55,
	0x52, 0x43, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x1a, 0x0a, 0x16, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x45, 0x52,
	0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x1b, 0x0a,
	0x17, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x5f, 0x4c, 0x49,
	0x53, 0x54, 0x5f, 0x4d, 0x59, 0x53, 0x51, 0x4c, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x52, 0x45,
	0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x5f,
	0x4d, 0x59, 0x53, 0x51, 0x4c, 0x10, 0x03, 0x42, 0x71, 0x0a, 0x21, 0x64, 0x65, 0x76, 0x2e, 0x74,
	0x6b, 0x65, 0x65, 0x6c, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x4f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x3a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6b, 0x65, 0x65, 0x6c,
	0x2d, 0x69, 0x6f, 0x2f, 0x74, 0x6b, 0x65, 0x65, 0x6c, 0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_helloworld_v1_error_proto_rawDescOnce sync.Once
	file_api_helloworld_v1_error_proto_rawDescData = file_api_helloworld_v1_error_proto_rawDesc
)

func file_api_helloworld_v1_error_proto_rawDescGZIP() []byte {
	file_api_helloworld_v1_error_proto_rawDescOnce.Do(func() {
		file_api_helloworld_v1_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_helloworld_v1_error_proto_rawDescData)
	})
	return file_api_helloworld_v1_error_proto_rawDescData
}

var file_api_helloworld_v1_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_helloworld_v1_error_proto_goTypes = []interface{}{
	(Error)(0), // 0: helloworld.v1.Error
}
var file_api_helloworld_v1_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_helloworld_v1_error_proto_init() }
func file_api_helloworld_v1_error_proto_init() {
	if File_api_helloworld_v1_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_helloworld_v1_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_helloworld_v1_error_proto_goTypes,
		DependencyIndexes: file_api_helloworld_v1_error_proto_depIdxs,
		EnumInfos:         file_api_helloworld_v1_error_proto_enumTypes,
	}.Build()
	File_api_helloworld_v1_error_proto = out.File
	file_api_helloworld_v1_error_proto_rawDesc = nil
	file_api_helloworld_v1_error_proto_goTypes = nil
	file_api_helloworld_v1_error_proto_depIdxs = nil
}
