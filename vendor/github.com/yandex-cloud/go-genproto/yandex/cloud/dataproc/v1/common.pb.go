// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: yandex/cloud/dataproc/v1/common.proto

package dataproc

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

type Health int32

const (
	// Object is in unknown state (we have no data).
	Health_HEALTH_UNKNOWN Health = 0
	// Object is alive and well (for example, all hosts of the cluster are alive).
	Health_ALIVE Health = 1
	// Object is inoperable (it cannot perform any of its essential functions).
	Health_DEAD Health = 2
	// Object is partially alive (it can perform some of its essential functions).
	Health_DEGRADED Health = 3
)

// Enum value maps for Health.
var (
	Health_name = map[int32]string{
		0: "HEALTH_UNKNOWN",
		1: "ALIVE",
		2: "DEAD",
		3: "DEGRADED",
	}
	Health_value = map[string]int32{
		"HEALTH_UNKNOWN": 0,
		"ALIVE":          1,
		"DEAD":           2,
		"DEGRADED":       3,
	}
)

func (x Health) Enum() *Health {
	p := new(Health)
	*p = x
	return p
}

func (x Health) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Health) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_dataproc_v1_common_proto_enumTypes[0].Descriptor()
}

func (Health) Type() protoreflect.EnumType {
	return &file_yandex_cloud_dataproc_v1_common_proto_enumTypes[0]
}

func (x Health) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Health.Descriptor instead.
func (Health) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_dataproc_v1_common_proto_rawDescGZIP(), []int{0}
}

type Resources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the resource preset for computational resources available to a host (CPU, memory etc.).
	// All available presets are listed in the [documentation](/docs/data-proc/concepts/instance-types).
	ResourcePresetId string `protobuf:"bytes,1,opt,name=resource_preset_id,json=resourcePresetId,proto3" json:"resource_preset_id,omitempty"`
	// Type of the storage environment for the host.
	// Possible values:
	// * network-hdd - network HDD drive,
	// * network-ssd - network SSD drive.
	DiskTypeId string `protobuf:"bytes,2,opt,name=disk_type_id,json=diskTypeId,proto3" json:"disk_type_id,omitempty"`
	// Volume of the storage available to a host, in bytes.
	DiskSize int64 `protobuf:"varint,3,opt,name=disk_size,json=diskSize,proto3" json:"disk_size,omitempty"`
}

func (x *Resources) Reset() {
	*x = Resources{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_dataproc_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resources) ProtoMessage() {}

func (x *Resources) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_dataproc_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resources.ProtoReflect.Descriptor instead.
func (*Resources) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_dataproc_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *Resources) GetResourcePresetId() string {
	if x != nil {
		return x.ResourcePresetId
	}
	return ""
}

func (x *Resources) GetDiskTypeId() string {
	if x != nil {
		return x.DiskTypeId
	}
	return ""
}

func (x *Resources) GetDiskSize() int64 {
	if x != nil {
		return x.DiskSize
	}
	return 0
}

var File_yandex_cloud_dataproc_v1_common_proto protoreflect.FileDescriptor

var file_yandex_cloud_dataproc_v1_common_proto_rawDesc = []byte{
	0x0a, 0x25, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x63, 0x2e, 0x76,
	0x31, 0x22, 0x78, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x2c,
	0x0a, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c,
	0x64, 0x69, 0x73, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x64, 0x69, 0x73, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x2a, 0x3f, 0x0a, 0x06, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x0e, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x5f,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x4c, 0x49,
	0x56, 0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x41, 0x44, 0x10, 0x02, 0x12, 0x0c,
	0x0a, 0x08, 0x44, 0x45, 0x47, 0x52, 0x41, 0x44, 0x45, 0x44, 0x10, 0x03, 0x42, 0x65, 0x0a, 0x1c,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x63, 0x2e, 0x76, 0x31, 0x5a, 0x45, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x70,
	0x72, 0x6f, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_dataproc_v1_common_proto_rawDescOnce sync.Once
	file_yandex_cloud_dataproc_v1_common_proto_rawDescData = file_yandex_cloud_dataproc_v1_common_proto_rawDesc
)

func file_yandex_cloud_dataproc_v1_common_proto_rawDescGZIP() []byte {
	file_yandex_cloud_dataproc_v1_common_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_dataproc_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_dataproc_v1_common_proto_rawDescData)
	})
	return file_yandex_cloud_dataproc_v1_common_proto_rawDescData
}

var file_yandex_cloud_dataproc_v1_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_dataproc_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_dataproc_v1_common_proto_goTypes = []interface{}{
	(Health)(0),       // 0: yandex.cloud.dataproc.v1.Health
	(*Resources)(nil), // 1: yandex.cloud.dataproc.v1.Resources
}
var file_yandex_cloud_dataproc_v1_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yandex_cloud_dataproc_v1_common_proto_init() }
func file_yandex_cloud_dataproc_v1_common_proto_init() {
	if File_yandex_cloud_dataproc_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_dataproc_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resources); i {
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
			RawDescriptor: file_yandex_cloud_dataproc_v1_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_dataproc_v1_common_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_dataproc_v1_common_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_dataproc_v1_common_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_dataproc_v1_common_proto_msgTypes,
	}.Build()
	File_yandex_cloud_dataproc_v1_common_proto = out.File
	file_yandex_cloud_dataproc_v1_common_proto_rawDesc = nil
	file_yandex_cloud_dataproc_v1_common_proto_goTypes = nil
	file_yandex_cloud_dataproc_v1_common_proto_depIdxs = nil
}