// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: services/stats_summary/damage_summary/damage_summary.proto

package damage_summary

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

type DamageSummaryItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Given int32  `protobuf:"varint,2,opt,name=given,proto3" json:"given,omitempty"`
	Taken int32  `protobuf:"varint,3,opt,name=taken,proto3" json:"taken,omitempty"`
}

func (x *DamageSummaryItem) Reset() {
	*x = DamageSummaryItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_stats_summary_damage_summary_damage_summary_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DamageSummaryItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DamageSummaryItem) ProtoMessage() {}

func (x *DamageSummaryItem) ProtoReflect() protoreflect.Message {
	mi := &file_services_stats_summary_damage_summary_damage_summary_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DamageSummaryItem.ProtoReflect.Descriptor instead.
func (*DamageSummaryItem) Descriptor() ([]byte, []int) {
	return file_services_stats_summary_damage_summary_damage_summary_proto_rawDescGZIP(), []int{0}
}

func (x *DamageSummaryItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DamageSummaryItem) GetGiven() int32 {
	if x != nil {
		return x.Given
	}
	return 0
}

func (x *DamageSummaryItem) GetTaken() int32 {
	if x != nil {
		return x.Taken
	}
	return 0
}

type DamageSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category string               `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	Items    []*DamageSummaryItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *DamageSummary) Reset() {
	*x = DamageSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_stats_summary_damage_summary_damage_summary_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DamageSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DamageSummary) ProtoMessage() {}

func (x *DamageSummary) ProtoReflect() protoreflect.Message {
	mi := &file_services_stats_summary_damage_summary_damage_summary_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DamageSummary.ProtoReflect.Descriptor instead.
func (*DamageSummary) Descriptor() ([]byte, []int) {
	return file_services_stats_summary_damage_summary_damage_summary_proto_rawDescGZIP(), []int{1}
}

func (x *DamageSummary) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *DamageSummary) GetItems() []*DamageSummaryItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type AllDamageSummaries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*DamageSummary `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *AllDamageSummaries) Reset() {
	*x = AllDamageSummaries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_stats_summary_damage_summary_damage_summary_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllDamageSummaries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllDamageSummaries) ProtoMessage() {}

func (x *AllDamageSummaries) ProtoReflect() protoreflect.Message {
	mi := &file_services_stats_summary_damage_summary_damage_summary_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllDamageSummaries.ProtoReflect.Descriptor instead.
func (*AllDamageSummaries) Descriptor() ([]byte, []int) {
	return file_services_stats_summary_damage_summary_damage_summary_proto_rawDescGZIP(), []int{2}
}

func (x *AllDamageSummaries) GetData() []*DamageSummary {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_services_stats_summary_damage_summary_damage_summary_proto protoreflect.FileDescriptor

var file_services_stats_summary_damage_summary_damage_summary_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2f, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2f, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x64, 0x61,
	0x6d, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x53, 0x0a, 0x11,
	0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x61, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x61, 0x6b, 0x65,
	0x6e, 0x22, 0x64, 0x0a, 0x0d, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x37,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x44,
	0x61, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x47, 0x0a, 0x12, 0x41, 0x6c, 0x6c, 0x44, 0x61,
	0x6d, 0x61, 0x67, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x65, 0x73, 0x12, 0x31, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x64, 0x61,
	0x6d, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x44, 0x61, 0x6d,
	0x61, 0x67, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x42, 0x4e, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41,
	0x62, 0x68, 0x69, 0x6c, 0x61, 0x73, 0x68, 0x4a, 0x4e, 0x2f, 0x67, 0x6f, 0x63, 0x73, 0x2d, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x2f, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_stats_summary_damage_summary_damage_summary_proto_rawDescOnce sync.Once
	file_services_stats_summary_damage_summary_damage_summary_proto_rawDescData = file_services_stats_summary_damage_summary_damage_summary_proto_rawDesc
)

func file_services_stats_summary_damage_summary_damage_summary_proto_rawDescGZIP() []byte {
	file_services_stats_summary_damage_summary_damage_summary_proto_rawDescOnce.Do(func() {
		file_services_stats_summary_damage_summary_damage_summary_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_stats_summary_damage_summary_damage_summary_proto_rawDescData)
	})
	return file_services_stats_summary_damage_summary_damage_summary_proto_rawDescData
}

var file_services_stats_summary_damage_summary_damage_summary_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_services_stats_summary_damage_summary_damage_summary_proto_goTypes = []interface{}{
	(*DamageSummaryItem)(nil),  // 0: damage_summary.DamageSummaryItem
	(*DamageSummary)(nil),      // 1: damage_summary.DamageSummary
	(*AllDamageSummaries)(nil), // 2: damage_summary.AllDamageSummaries
}
var file_services_stats_summary_damage_summary_damage_summary_proto_depIdxs = []int32{
	0, // 0: damage_summary.DamageSummary.items:type_name -> damage_summary.DamageSummaryItem
	1, // 1: damage_summary.AllDamageSummaries.data:type_name -> damage_summary.DamageSummary
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_services_stats_summary_damage_summary_damage_summary_proto_init() }
func file_services_stats_summary_damage_summary_damage_summary_proto_init() {
	if File_services_stats_summary_damage_summary_damage_summary_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_stats_summary_damage_summary_damage_summary_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DamageSummaryItem); i {
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
		file_services_stats_summary_damage_summary_damage_summary_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DamageSummary); i {
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
		file_services_stats_summary_damage_summary_damage_summary_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllDamageSummaries); i {
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
			RawDescriptor: file_services_stats_summary_damage_summary_damage_summary_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_stats_summary_damage_summary_damage_summary_proto_goTypes,
		DependencyIndexes: file_services_stats_summary_damage_summary_damage_summary_proto_depIdxs,
		MessageInfos:      file_services_stats_summary_damage_summary_damage_summary_proto_msgTypes,
	}.Build()
	File_services_stats_summary_damage_summary_damage_summary_proto = out.File
	file_services_stats_summary_damage_summary_damage_summary_proto_rawDesc = nil
	file_services_stats_summary_damage_summary_damage_summary_proto_goTypes = nil
	file_services_stats_summary_damage_summary_damage_summary_proto_depIdxs = nil
}
