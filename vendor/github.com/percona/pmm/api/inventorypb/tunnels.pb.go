// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: inventorypb/tunnels.proto

package inventorypb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// TunnelType represents tunnel type.
type TunnelType int32

const (
	// Invalid tunnel type.
	TunnelType_TUNNEL_TYPE_INVALID TunnelType = 0
	// pmm-agent will act as a TCP client and connect to the given address.
	TunnelType_TCP_CONNECT TunnelType = 1
	// pmm-agent will act as a TCP server and listen on the given address.
	TunnelType_TCP_LISTEN TunnelType = 2
	// pmm-agent will send a file to the PMM Server.
	TunnelType_FILE_SEND TunnelType = 3
)

// Enum value maps for TunnelType.
var (
	TunnelType_name = map[int32]string{
		0: "TUNNEL_TYPE_INVALID",
		1: "TCP_CONNECT",
		2: "TCP_LISTEN",
		3: "FILE_SEND",
	}
	TunnelType_value = map[string]int32{
		"TUNNEL_TYPE_INVALID": 0,
		"TCP_CONNECT":         1,
		"TCP_LISTEN":          2,
		"FILE_SEND":           3,
	}
)

func (x TunnelType) Enum() *TunnelType {
	p := new(TunnelType)
	*p = x
	return p
}

func (x TunnelType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TunnelType) Descriptor() protoreflect.EnumDescriptor {
	return file_inventorypb_tunnels_proto_enumTypes[0].Descriptor()
}

func (TunnelType) Type() protoreflect.EnumType {
	return &file_inventorypb_tunnels_proto_enumTypes[0]
}

func (x TunnelType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TunnelType.Descriptor instead.
func (TunnelType) EnumDescriptor() ([]byte, []int) {
	return file_inventorypb_tunnels_proto_rawDescGZIP(), []int{0}
}

var File_inventorypb_tunnels_proto protoreflect.FileDescriptor

var file_inventorypb_tunnels_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x70, 0x62, 0x2f, 0x74, 0x75,
	0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2a, 0x55, 0x0a, 0x0a, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x54, 0x55, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a,
	0x0b, 0x54, 0x43, 0x50, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x10, 0x01, 0x12, 0x0e,
	0x0a, 0x0a, 0x54, 0x43, 0x50, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x45, 0x4e, 0x10, 0x02, 0x12, 0x0d,
	0x0a, 0x09, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x53, 0x45, 0x4e, 0x44, 0x10, 0x03, 0x42, 0x1d, 0x5a,
	0x1b, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x70, 0x62,
	0x3b, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inventorypb_tunnels_proto_rawDescOnce sync.Once
	file_inventorypb_tunnels_proto_rawDescData = file_inventorypb_tunnels_proto_rawDesc
)

func file_inventorypb_tunnels_proto_rawDescGZIP() []byte {
	file_inventorypb_tunnels_proto_rawDescOnce.Do(func() {
		file_inventorypb_tunnels_proto_rawDescData = protoimpl.X.CompressGZIP(file_inventorypb_tunnels_proto_rawDescData)
	})
	return file_inventorypb_tunnels_proto_rawDescData
}

var file_inventorypb_tunnels_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_inventorypb_tunnels_proto_goTypes = []interface{}{
	(TunnelType)(0), // 0: inventory.TunnelType
}
var file_inventorypb_tunnels_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_inventorypb_tunnels_proto_init() }
func file_inventorypb_tunnels_proto_init() {
	if File_inventorypb_tunnels_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_inventorypb_tunnels_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_inventorypb_tunnels_proto_goTypes,
		DependencyIndexes: file_inventorypb_tunnels_proto_depIdxs,
		EnumInfos:         file_inventorypb_tunnels_proto_enumTypes,
	}.Build()
	File_inventorypb_tunnels_proto = out.File
	file_inventorypb_tunnels_proto_rawDesc = nil
	file_inventorypb_tunnels_proto_goTypes = nil
	file_inventorypb_tunnels_proto_depIdxs = nil
}
