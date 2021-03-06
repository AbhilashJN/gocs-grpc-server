// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: services/map_details/map_details.proto

package map_details

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

type MapNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MapNameRequest) Reset() {
	*x = MapNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_map_details_map_details_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapNameRequest) ProtoMessage() {}

func (x *MapNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_map_details_map_details_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapNameRequest.ProtoReflect.Descriptor instead.
func (*MapNameRequest) Descriptor() ([]byte, []int) {
	return file_services_map_details_map_details_proto_rawDescGZIP(), []int{0}
}

type MapNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MapName string `protobuf:"bytes,1,opt,name=map_name,json=mapName,proto3" json:"map_name,omitempty"`
}

func (x *MapNameResponse) Reset() {
	*x = MapNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_map_details_map_details_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapNameResponse) ProtoMessage() {}

func (x *MapNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_map_details_map_details_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapNameResponse.ProtoReflect.Descriptor instead.
func (*MapNameResponse) Descriptor() ([]byte, []int) {
	return file_services_map_details_map_details_proto_rawDescGZIP(), []int{1}
}

func (x *MapNameResponse) GetMapName() string {
	if x != nil {
		return x.MapName
	}
	return ""
}

var File_services_map_details_map_details_proto protoreflect.FileDescriptor

var file_services_map_details_map_details_proto_rawDesc = []byte{
	0x0a, 0x26, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x61, 0x70, 0x5f, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2f, 0x6d, 0x61, 0x70, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x61, 0x70, 0x5f, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x10, 0x0a, 0x0e, 0x4d, 0x61, 0x70, 0x4e, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2c, 0x0a, 0x0f, 0x4d, 0x61, 0x70, 0x4e, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61,
	0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61,
	0x70, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x5c, 0x0a, 0x11, 0x4d, 0x61, 0x70, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x4d, 0x61, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x2e, 0x6d, 0x61, 0x70, 0x5f, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x61, 0x70, 0x5f, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x41, 0x62, 0x68, 0x69, 0x6c, 0x61, 0x73, 0x68, 0x4a, 0x4e, 0x2f, 0x67, 0x6f, 0x63,
	0x73, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x61, 0x70, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_map_details_map_details_proto_rawDescOnce sync.Once
	file_services_map_details_map_details_proto_rawDescData = file_services_map_details_map_details_proto_rawDesc
)

func file_services_map_details_map_details_proto_rawDescGZIP() []byte {
	file_services_map_details_map_details_proto_rawDescOnce.Do(func() {
		file_services_map_details_map_details_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_map_details_map_details_proto_rawDescData)
	})
	return file_services_map_details_map_details_proto_rawDescData
}

var file_services_map_details_map_details_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_services_map_details_map_details_proto_goTypes = []interface{}{
	(*MapNameRequest)(nil),  // 0: map_details.MapNameRequest
	(*MapNameResponse)(nil), // 1: map_details.MapNameResponse
}
var file_services_map_details_map_details_proto_depIdxs = []int32{
	0, // 0: map_details.MapDetailsService.GetMapName:input_type -> map_details.MapNameRequest
	1, // 1: map_details.MapDetailsService.GetMapName:output_type -> map_details.MapNameResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_map_details_map_details_proto_init() }
func file_services_map_details_map_details_proto_init() {
	if File_services_map_details_map_details_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_map_details_map_details_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapNameRequest); i {
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
		file_services_map_details_map_details_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapNameResponse); i {
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
			RawDescriptor: file_services_map_details_map_details_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_map_details_map_details_proto_goTypes,
		DependencyIndexes: file_services_map_details_map_details_proto_depIdxs,
		MessageInfos:      file_services_map_details_map_details_proto_msgTypes,
	}.Build()
	File_services_map_details_map_details_proto = out.File
	file_services_map_details_map_details_proto_rawDesc = nil
	file_services_map_details_map_details_proto_goTypes = nil
	file_services_map_details_map_details_proto_depIdxs = nil
}
