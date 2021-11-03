// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: zmagmacore/magmasc/pb/proto/access_point.proto

package pb

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

// AccessPoint represents access point node stored in blockchain.
type AccessPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" yaml:"id"`                                              // @gotags: json:"id" yaml:"id"
	ProviderExtId string `protobuf:"bytes,2,opt,name=provider_ext_id,json=providerExtId,proto3" json:"provider_ext_id" yaml:"provider_ext_id"` // @gotags: json:"provider_ext_id" yaml:"provider_ext_id"
	Terms         *Terms `protobuf:"bytes,3,opt,name=terms,proto3" json:"terms,omitempty" yaml:"terms"`                                        // @gotags: yaml:"terms"
}

func (x *AccessPoint) Reset() {
	*x = AccessPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zmagmacore_magmasc_pb_proto_access_point_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessPoint) ProtoMessage() {}

func (x *AccessPoint) ProtoReflect() protoreflect.Message {
	mi := &file_zmagmacore_magmasc_pb_proto_access_point_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessPoint.ProtoReflect.Descriptor instead.
func (*AccessPoint) Descriptor() ([]byte, []int) {
	return file_zmagmacore_magmasc_pb_proto_access_point_proto_rawDescGZIP(), []int{0}
}

func (x *AccessPoint) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AccessPoint) GetProviderExtId() string {
	if x != nil {
		return x.ProviderExtId
	}
	return ""
}

func (x *AccessPoint) GetTerms() *Terms {
	if x != nil {
		return x.Terms
	}
	return nil
}

var File_zmagmacore_magmasc_pb_proto_access_point_proto protoreflect.FileDescriptor

var file_zmagmacore_magmasc_pb_proto_access_point_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x7a, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x61, 0x67,
	0x6d, 0x61, 0x73, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x7a, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x62, 0x2e, 0x6d, 0x61, 0x67, 0x6d,
	0x61, 0x73, 0x63, 0x1a, 0x27, 0x7a, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x6d, 0x61, 0x67, 0x6d, 0x61, 0x73, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x0b,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x45, 0x78,
	0x74, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x05, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x7a, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x62, 0x2e, 0x6d,
	0x61, 0x67, 0x6d, 0x61, 0x73, 0x63, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x52, 0x05, 0x74, 0x65,
	0x72, 0x6d, 0x73, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x30, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x7a,
	0x6d, 0x61, 0x67, 0x6d, 0x61, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x73,
	0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_zmagmacore_magmasc_pb_proto_access_point_proto_rawDescOnce sync.Once
	file_zmagmacore_magmasc_pb_proto_access_point_proto_rawDescData = file_zmagmacore_magmasc_pb_proto_access_point_proto_rawDesc
)

func file_zmagmacore_magmasc_pb_proto_access_point_proto_rawDescGZIP() []byte {
	file_zmagmacore_magmasc_pb_proto_access_point_proto_rawDescOnce.Do(func() {
		file_zmagmacore_magmasc_pb_proto_access_point_proto_rawDescData = protoimpl.X.CompressGZIP(file_zmagmacore_magmasc_pb_proto_access_point_proto_rawDescData)
	})
	return file_zmagmacore_magmasc_pb_proto_access_point_proto_rawDescData
}

var file_zmagmacore_magmasc_pb_proto_access_point_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_zmagmacore_magmasc_pb_proto_access_point_proto_goTypes = []interface{}{
	(*AccessPoint)(nil), // 0: zchain.pb.magmasc.AccessPoint
	(*Terms)(nil),       // 1: zchain.pb.magmasc.Terms
}
var file_zmagmacore_magmasc_pb_proto_access_point_proto_depIdxs = []int32{
	1, // 0: zchain.pb.magmasc.AccessPoint.terms:type_name -> zchain.pb.magmasc.Terms
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_zmagmacore_magmasc_pb_proto_access_point_proto_init() }
func file_zmagmacore_magmasc_pb_proto_access_point_proto_init() {
	if File_zmagmacore_magmasc_pb_proto_access_point_proto != nil {
		return
	}
	file_zmagmacore_magmasc_pb_proto_terms_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_zmagmacore_magmasc_pb_proto_access_point_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessPoint); i {
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
			RawDescriptor: file_zmagmacore_magmasc_pb_proto_access_point_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_zmagmacore_magmasc_pb_proto_access_point_proto_goTypes,
		DependencyIndexes: file_zmagmacore_magmasc_pb_proto_access_point_proto_depIdxs,
		MessageInfos:      file_zmagmacore_magmasc_pb_proto_access_point_proto_msgTypes,
	}.Build()
	File_zmagmacore_magmasc_pb_proto_access_point_proto = out.File
	file_zmagmacore_magmasc_pb_proto_access_point_proto_rawDesc = nil
	file_zmagmacore_magmasc_pb_proto_access_point_proto_goTypes = nil
	file_zmagmacore_magmasc_pb_proto_access_point_proto_depIdxs = nil
}
