// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: communication/communication.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProvisionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ProvisionRequest) Reset() {
	*x = ProvisionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_communication_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProvisionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProvisionRequest) ProtoMessage() {}

func (x *ProvisionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_communication_communication_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProvisionRequest.ProtoReflect.Descriptor instead.
func (*ProvisionRequest) Descriptor() ([]byte, []int) {
	return file_communication_communication_proto_rawDescGZIP(), []int{0}
}

func (x *ProvisionRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ProvisionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *structpb.Struct `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ProvisionResponse) Reset() {
	*x = ProvisionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_communication_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProvisionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProvisionResponse) ProtoMessage() {}

func (x *ProvisionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_communication_communication_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProvisionResponse.ProtoReflect.Descriptor instead.
func (*ProvisionResponse) Descriptor() ([]byte, []int) {
	return file_communication_communication_proto_rawDescGZIP(), []int{1}
}

func (x *ProvisionResponse) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

type ConnectorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ConnectorRequest) Reset() {
	*x = ConnectorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_communication_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectorRequest) ProtoMessage() {}

func (x *ConnectorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_communication_communication_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectorRequest.ProtoReflect.Descriptor instead.
func (*ConnectorRequest) Descriptor() ([]byte, []int) {
	return file_communication_communication_proto_rawDescGZIP(), []int{2}
}

func (x *ConnectorRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ConnectorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *structpb.Struct `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ConnectorResponse) Reset() {
	*x = ConnectorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_communication_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectorResponse) ProtoMessage() {}

func (x *ConnectorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_communication_communication_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectorResponse.ProtoReflect.Descriptor instead.
func (*ConnectorResponse) Descriptor() ([]byte, []int) {
	return file_communication_communication_proto_rawDescGZIP(), []int{3}
}

func (x *ConnectorResponse) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_communication_communication_proto protoreflect.FileDescriptor

var file_communication_communication_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x40, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2c, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x40, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x4f, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x0d, 0x43, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x4f, 0x0a, 0x09, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x42, 0x0a, 0x0d, 0x43, 0x61, 0x6c, 0x6c, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_communication_communication_proto_rawDescOnce sync.Once
	file_communication_communication_proto_rawDescData = file_communication_communication_proto_rawDesc
)

func file_communication_communication_proto_rawDescGZIP() []byte {
	file_communication_communication_proto_rawDescOnce.Do(func() {
		file_communication_communication_proto_rawDescData = protoimpl.X.CompressGZIP(file_communication_communication_proto_rawDescData)
	})
	return file_communication_communication_proto_rawDescData
}

var file_communication_communication_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_communication_communication_proto_goTypes = []interface{}{
	(*ProvisionRequest)(nil),  // 0: proto.ProvisionRequest
	(*ProvisionResponse)(nil), // 1: proto.ProvisionResponse
	(*ConnectorRequest)(nil),  // 2: proto.ConnectorRequest
	(*ConnectorResponse)(nil), // 3: proto.ConnectorResponse
	(*structpb.Struct)(nil),   // 4: google.protobuf.Struct
}
var file_communication_communication_proto_depIdxs = []int32{
	4, // 0: proto.ProvisionResponse.data:type_name -> google.protobuf.Struct
	4, // 1: proto.ConnectorResponse.data:type_name -> google.protobuf.Struct
	2, // 2: proto.Provision.CallConnector:input_type -> proto.ConnectorRequest
	0, // 3: proto.Connector.CallProvision:input_type -> proto.ProvisionRequest
	3, // 4: proto.Provision.CallConnector:output_type -> proto.ConnectorResponse
	1, // 5: proto.Connector.CallProvision:output_type -> proto.ProvisionResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_communication_communication_proto_init() }
func file_communication_communication_proto_init() {
	if File_communication_communication_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_communication_communication_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProvisionRequest); i {
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
		file_communication_communication_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProvisionResponse); i {
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
		file_communication_communication_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectorRequest); i {
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
		file_communication_communication_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectorResponse); i {
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
			RawDescriptor: file_communication_communication_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_communication_communication_proto_goTypes,
		DependencyIndexes: file_communication_communication_proto_depIdxs,
		MessageInfos:      file_communication_communication_proto_msgTypes,
	}.Build()
	File_communication_communication_proto = out.File
	file_communication_communication_proto_rawDesc = nil
	file_communication_communication_proto_goTypes = nil
	file_communication_communication_proto_depIdxs = nil
}