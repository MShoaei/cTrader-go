// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.4
// source: OpenApiCommonMessages.proto

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

// * Base message that is used for all messages that are sent to/from Open API proxy of cTrader platform.
type ProtoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PayloadType *uint32 `protobuf:"varint,1,req,name=payloadType" json:"payloadType,omitempty"` // Contains id of ProtoPayloadType or other custom PayloadTypes (e.g. ProtoOAPayloadType).
	Payload     []byte  `protobuf:"bytes,2,opt,name=payload" json:"payload,omitempty"`          // Serialized protobuf message that corresponds to payloadType.
	ClientMsgId *string `protobuf:"bytes,3,opt,name=clientMsgId" json:"clientMsgId,omitempty"`  // Request message id, assigned by the client that will be returned in the response.
}

func (x *ProtoMessage) Reset() {
	*x = ProtoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpenApiCommonMessages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoMessage) ProtoMessage() {}

func (x *ProtoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_OpenApiCommonMessages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoMessage.ProtoReflect.Descriptor instead.
func (*ProtoMessage) Descriptor() ([]byte, []int) {
	return file_OpenApiCommonMessages_proto_rawDescGZIP(), []int{0}
}

func (x *ProtoMessage) GetPayloadType() uint32 {
	if x != nil && x.PayloadType != nil {
		return *x.PayloadType
	}
	return 0
}

func (x *ProtoMessage) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *ProtoMessage) GetClientMsgId() string {
	if x != nil && x.ClientMsgId != nil {
		return *x.ClientMsgId
	}
	return ""
}

// * Error response that is sent from Open API proxy when error occurs.
type ProtoErrorRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PayloadType             *ProtoPayloadType `protobuf:"varint,1,opt,name=payloadType,enum=ProtoPayloadType,def=50" json:"payloadType,omitempty"`
	ErrorCode               *string           `protobuf:"bytes,2,req,name=errorCode" json:"errorCode,omitempty"`                              // Contains name of ProtoErrorCode or other custom ErrorCodes (e.g. ProtoCHErrorCode).
	Description             *string           `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`                          // Error description.
	MaintenanceEndTimestamp *uint64           `protobuf:"varint,4,opt,name=maintenanceEndTimestamp" json:"maintenanceEndTimestamp,omitempty"` // The Unix time in milliseconds of the end of the maintenance.
}

// Default values for ProtoErrorRes fields.
const (
	Default_ProtoErrorRes_PayloadType = ProtoPayloadType_ERROR_RES
)

func (x *ProtoErrorRes) Reset() {
	*x = ProtoErrorRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpenApiCommonMessages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoErrorRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoErrorRes) ProtoMessage() {}

func (x *ProtoErrorRes) ProtoReflect() protoreflect.Message {
	mi := &file_OpenApiCommonMessages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoErrorRes.ProtoReflect.Descriptor instead.
func (*ProtoErrorRes) Descriptor() ([]byte, []int) {
	return file_OpenApiCommonMessages_proto_rawDescGZIP(), []int{1}
}

func (x *ProtoErrorRes) GetPayloadType() ProtoPayloadType {
	if x != nil && x.PayloadType != nil {
		return *x.PayloadType
	}
	return Default_ProtoErrorRes_PayloadType
}

func (x *ProtoErrorRes) GetErrorCode() string {
	if x != nil && x.ErrorCode != nil {
		return *x.ErrorCode
	}
	return ""
}

func (x *ProtoErrorRes) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *ProtoErrorRes) GetMaintenanceEndTimestamp() uint64 {
	if x != nil && x.MaintenanceEndTimestamp != nil {
		return *x.MaintenanceEndTimestamp
	}
	return 0
}

// * Event that is sent from Open API proxy and can be used as criteria that connection is healthy when no other messages are sent by cTrader platform. Open API client can send this message when he needs to keep the connection open for a period without other messages longer than 30 seconds.
type ProtoHeartbeatEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PayloadType *ProtoPayloadType `protobuf:"varint,1,opt,name=payloadType,enum=ProtoPayloadType,def=51" json:"payloadType,omitempty"`
}

// Default values for ProtoHeartbeatEvent fields.
const (
	Default_ProtoHeartbeatEvent_PayloadType = ProtoPayloadType_HEARTBEAT_EVENT
)

func (x *ProtoHeartbeatEvent) Reset() {
	*x = ProtoHeartbeatEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpenApiCommonMessages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoHeartbeatEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoHeartbeatEvent) ProtoMessage() {}

func (x *ProtoHeartbeatEvent) ProtoReflect() protoreflect.Message {
	mi := &file_OpenApiCommonMessages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoHeartbeatEvent.ProtoReflect.Descriptor instead.
func (*ProtoHeartbeatEvent) Descriptor() ([]byte, []int) {
	return file_OpenApiCommonMessages_proto_rawDescGZIP(), []int{2}
}

func (x *ProtoHeartbeatEvent) GetPayloadType() ProtoPayloadType {
	if x != nil && x.PayloadType != nil {
		return *x.PayloadType
	}
	return Default_ProtoHeartbeatEvent_PayloadType
}

var File_OpenApiCommonMessages_proto protoreflect.FileDescriptor

var file_OpenApiCommonMessages_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x4f,
	0x70, 0x65, 0x6e, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x6c, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x0d, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x22, 0xc9, 0x01,
	0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x12,
	0x3e, 0x0a, 0x0b, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x09, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52,
	0x45, 0x53, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x02,
	0x28, 0x09, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x38, 0x0a, 0x17, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x45, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x17, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x45, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x5b, 0x0a, 0x13, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x44, 0x0a, 0x0b, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x0f, 0x48, 0x45, 0x41, 0x52, 0x54, 0x42,
	0x45, 0x41, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x42, 0x4c, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x74,
	0x72, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x42, 0x17, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x50, 0x01, 0x5a, 0x08, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0xa0, 0x01, 0x01,
}

var (
	file_OpenApiCommonMessages_proto_rawDescOnce sync.Once
	file_OpenApiCommonMessages_proto_rawDescData = file_OpenApiCommonMessages_proto_rawDesc
)

func file_OpenApiCommonMessages_proto_rawDescGZIP() []byte {
	file_OpenApiCommonMessages_proto_rawDescOnce.Do(func() {
		file_OpenApiCommonMessages_proto_rawDescData = protoimpl.X.CompressGZIP(file_OpenApiCommonMessages_proto_rawDescData)
	})
	return file_OpenApiCommonMessages_proto_rawDescData
}

var file_OpenApiCommonMessages_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_OpenApiCommonMessages_proto_goTypes = []interface{}{
	(*ProtoMessage)(nil),        // 0: ProtoMessage
	(*ProtoErrorRes)(nil),       // 1: ProtoErrorRes
	(*ProtoHeartbeatEvent)(nil), // 2: ProtoHeartbeatEvent
	(ProtoPayloadType)(0),       // 3: ProtoPayloadType
}
var file_OpenApiCommonMessages_proto_depIdxs = []int32{
	3, // 0: ProtoErrorRes.payloadType:type_name -> ProtoPayloadType
	3, // 1: ProtoHeartbeatEvent.payloadType:type_name -> ProtoPayloadType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_OpenApiCommonMessages_proto_init() }
func file_OpenApiCommonMessages_proto_init() {
	if File_OpenApiCommonMessages_proto != nil {
		return
	}
	file_OpenApiCommonModelMessages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_OpenApiCommonMessages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_OpenApiCommonMessages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoErrorRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_OpenApiCommonMessages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoHeartbeatEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_OpenApiCommonMessages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_OpenApiCommonMessages_proto_goTypes,
		DependencyIndexes: file_OpenApiCommonMessages_proto_depIdxs,
		MessageInfos:      file_OpenApiCommonMessages_proto_msgTypes,
	}.Build()
	File_OpenApiCommonMessages_proto = out.File
	file_OpenApiCommonMessages_proto_rawDesc = nil
	file_OpenApiCommonMessages_proto_goTypes = nil
	file_OpenApiCommonMessages_proto_depIdxs = nil
}
