// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.4
// source: OpenApiCommonModelMessages.proto

package openapi

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

// --- INTENSIVE COMMANDS 1 - 49
// --- COMMON API 50 - 69
type ProtoPayloadType int32

const (
	// common intensive
	ProtoPayloadType_PROTO_MESSAGE ProtoPayloadType = 5
	// common commands
	ProtoPayloadType_ERROR_RES       ProtoPayloadType = 50
	ProtoPayloadType_HEARTBEAT_EVENT ProtoPayloadType = 51
)

// Enum value maps for ProtoPayloadType.
var (
	ProtoPayloadType_name = map[int32]string{
		5:  "PROTO_MESSAGE",
		50: "ERROR_RES",
		51: "HEARTBEAT_EVENT",
	}
	ProtoPayloadType_value = map[string]int32{
		"PROTO_MESSAGE":   5,
		"ERROR_RES":       50,
		"HEARTBEAT_EVENT": 51,
	}
)

func (x ProtoPayloadType) Enum() *ProtoPayloadType {
	p := new(ProtoPayloadType)
	*p = x
	return p
}

func (x ProtoPayloadType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProtoPayloadType) Descriptor() protoreflect.EnumDescriptor {
	return file_OpenApiCommonModelMessages_proto_enumTypes[0].Descriptor()
}

func (ProtoPayloadType) Type() protoreflect.EnumType {
	return &file_OpenApiCommonModelMessages_proto_enumTypes[0]
}

func (x ProtoPayloadType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ProtoPayloadType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ProtoPayloadType(num)
	return nil
}

// Deprecated: Use ProtoPayloadType.Descriptor instead.
func (ProtoPayloadType) EnumDescriptor() ([]byte, []int) {
	return file_OpenApiCommonModelMessages_proto_rawDescGZIP(), []int{0}
}

// COMMON error codes 1 - 99
type ProtoErrorCode int32

const (
	ProtoErrorCode_UNKNOWN_ERROR           ProtoErrorCode = 1  // Generic error.
	ProtoErrorCode_UNSUPPORTED_MESSAGE     ProtoErrorCode = 2  // Message is not supported. Wrong message.
	ProtoErrorCode_INVALID_REQUEST         ProtoErrorCode = 3  // Generic error.  Usually used when input value is not correct.
	ProtoErrorCode_TIMEOUT_ERROR           ProtoErrorCode = 5  // Deal execution is reached timeout and rejected.
	ProtoErrorCode_ENTITY_NOT_FOUND        ProtoErrorCode = 6  // Generic error for requests by id.
	ProtoErrorCode_CANT_ROUTE_REQUEST      ProtoErrorCode = 7  // Connection to Server is lost or not supported.
	ProtoErrorCode_FRAME_TOO_LONG          ProtoErrorCode = 8  // Message is too large.
	ProtoErrorCode_MARKET_CLOSED           ProtoErrorCode = 9  // Market is closed.
	ProtoErrorCode_CONCURRENT_MODIFICATION ProtoErrorCode = 10 // Order is blocked (e.g. under execution) and change cannot be applied.
	ProtoErrorCode_BLOCKED_PAYLOAD_TYPE    ProtoErrorCode = 11 // Message is blocked by server or rate limit is reached.
)

// Enum value maps for ProtoErrorCode.
var (
	ProtoErrorCode_name = map[int32]string{
		1:  "UNKNOWN_ERROR",
		2:  "UNSUPPORTED_MESSAGE",
		3:  "INVALID_REQUEST",
		5:  "TIMEOUT_ERROR",
		6:  "ENTITY_NOT_FOUND",
		7:  "CANT_ROUTE_REQUEST",
		8:  "FRAME_TOO_LONG",
		9:  "MARKET_CLOSED",
		10: "CONCURRENT_MODIFICATION",
		11: "BLOCKED_PAYLOAD_TYPE",
	}
	ProtoErrorCode_value = map[string]int32{
		"UNKNOWN_ERROR":           1,
		"UNSUPPORTED_MESSAGE":     2,
		"INVALID_REQUEST":         3,
		"TIMEOUT_ERROR":           5,
		"ENTITY_NOT_FOUND":        6,
		"CANT_ROUTE_REQUEST":      7,
		"FRAME_TOO_LONG":          8,
		"MARKET_CLOSED":           9,
		"CONCURRENT_MODIFICATION": 10,
		"BLOCKED_PAYLOAD_TYPE":    11,
	}
)

func (x ProtoErrorCode) Enum() *ProtoErrorCode {
	p := new(ProtoErrorCode)
	*p = x
	return p
}

func (x ProtoErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProtoErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_OpenApiCommonModelMessages_proto_enumTypes[1].Descriptor()
}

func (ProtoErrorCode) Type() protoreflect.EnumType {
	return &file_OpenApiCommonModelMessages_proto_enumTypes[1]
}

func (x ProtoErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ProtoErrorCode) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ProtoErrorCode(num)
	return nil
}

// Deprecated: Use ProtoErrorCode.Descriptor instead.
func (ProtoErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_OpenApiCommonModelMessages_proto_rawDescGZIP(), []int{1}
}

var File_OpenApiCommonModelMessages_proto protoreflect.FileDescriptor

var file_OpenApiCommonModelMessages_proto_rawDesc = []byte{
	0x0a, 0x20, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2a, 0x49, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x5f,
	0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x5f, 0x52, 0x45, 0x53, 0x10, 0x32, 0x12, 0x13, 0x0a, 0x0f, 0x48, 0x45, 0x41, 0x52,
	0x54, 0x42, 0x45, 0x41, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x10, 0x33, 0x2a, 0xf0, 0x01,
	0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54,
	0x45, 0x44, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10,
	0x03, 0x12, 0x11, 0x0a, 0x0d, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x06, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x41,
	0x4e, 0x54, 0x5f, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54,
	0x10, 0x07, 0x12, 0x12, 0x0a, 0x0e, 0x46, 0x52, 0x41, 0x4d, 0x45, 0x5f, 0x54, 0x4f, 0x4f, 0x5f,
	0x4c, 0x4f, 0x4e, 0x47, 0x10, 0x08, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54,
	0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10, 0x09, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4f, 0x4e,
	0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x49, 0x43, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0a, 0x12, 0x18, 0x0a, 0x14, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x45,
	0x44, 0x5f, 0x50, 0x41, 0x59, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x0b,
	0x42, 0x57, 0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x1c, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x50, 0x01, 0x5a, 0x08, 0x2f, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0xa0, 0x01, 0x01,
}

var (
	file_OpenApiCommonModelMessages_proto_rawDescOnce sync.Once
	file_OpenApiCommonModelMessages_proto_rawDescData = file_OpenApiCommonModelMessages_proto_rawDesc
)

func file_OpenApiCommonModelMessages_proto_rawDescGZIP() []byte {
	file_OpenApiCommonModelMessages_proto_rawDescOnce.Do(func() {
		file_OpenApiCommonModelMessages_proto_rawDescData = protoimpl.X.CompressGZIP(file_OpenApiCommonModelMessages_proto_rawDescData)
	})
	return file_OpenApiCommonModelMessages_proto_rawDescData
}

var file_OpenApiCommonModelMessages_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_OpenApiCommonModelMessages_proto_goTypes = []interface{}{
	(ProtoPayloadType)(0), // 0: ProtoPayloadType
	(ProtoErrorCode)(0),   // 1: ProtoErrorCode
}
var file_OpenApiCommonModelMessages_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_OpenApiCommonModelMessages_proto_init() }
func file_OpenApiCommonModelMessages_proto_init() {
	if File_OpenApiCommonModelMessages_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_OpenApiCommonModelMessages_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_OpenApiCommonModelMessages_proto_goTypes,
		DependencyIndexes: file_OpenApiCommonModelMessages_proto_depIdxs,
		EnumInfos:         file_OpenApiCommonModelMessages_proto_enumTypes,
	}.Build()
	File_OpenApiCommonModelMessages_proto = out.File
	file_OpenApiCommonModelMessages_proto_rawDesc = nil
	file_OpenApiCommonModelMessages_proto_goTypes = nil
	file_OpenApiCommonModelMessages_proto_depIdxs = nil
}
