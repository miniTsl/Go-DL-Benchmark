// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: pkg/protos/physicalDeviceDescription.proto

package protos

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

type DeviceOSType int32

const (
	DeviceOSType_none    DeviceOSType = 0
	DeviceOSType_android DeviceOSType = 1
	DeviceOSType_ios     DeviceOSType = 2
	DeviceOSType_mac     DeviceOSType = 3
	DeviceOSType_linux   DeviceOSType = 4
	DeviceOSType_windows DeviceOSType = 5
)

// Enum value maps for DeviceOSType.
var (
	DeviceOSType_name = map[int32]string{
		0: "none",
		1: "android",
		2: "ios",
		3: "mac",
		4: "linux",
		5: "windows",
	}
	DeviceOSType_value = map[string]int32{
		"none":    0,
		"android": 1,
		"ios":     2,
		"mac":     3,
		"linux":   4,
		"windows": 5,
	}
)

func (x DeviceOSType) Enum() *DeviceOSType {
	p := new(DeviceOSType)
	*p = x
	return p
}

func (x DeviceOSType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeviceOSType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_protos_physicalDeviceDescription_proto_enumTypes[0].Descriptor()
}

func (DeviceOSType) Type() protoreflect.EnumType {
	return &file_pkg_protos_physicalDeviceDescription_proto_enumTypes[0]
}

func (x DeviceOSType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeviceOSType.Descriptor instead.
func (DeviceOSType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_protos_physicalDeviceDescription_proto_rawDescGZIP(), []int{0}
}

type ComputableChipType int32

const (
	ComputableChipType_cpu ComputableChipType = 0
	ComputableChipType_gpu ComputableChipType = 1
	ComputableChipType_npu ComputableChipType = 2
	ComputableChipType_tpu ComputableChipType = 3
)

// Enum value maps for ComputableChipType.
var (
	ComputableChipType_name = map[int32]string{
		0: "cpu",
		1: "gpu",
		2: "npu",
		3: "tpu",
	}
	ComputableChipType_value = map[string]int32{
		"cpu": 0,
		"gpu": 1,
		"npu": 2,
		"tpu": 3,
	}
)

func (x ComputableChipType) Enum() *ComputableChipType {
	p := new(ComputableChipType)
	*p = x
	return p
}

func (x ComputableChipType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ComputableChipType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_protos_physicalDeviceDescription_proto_enumTypes[1].Descriptor()
}

func (ComputableChipType) Type() protoreflect.EnumType {
	return &file_pkg_protos_physicalDeviceDescription_proto_enumTypes[1]
}

func (x ComputableChipType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ComputableChipType.Descriptor instead.
func (ComputableChipType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_protos_physicalDeviceDescription_proto_rawDescGZIP(), []int{1}
}

type ArchitectureType int32

const (
	ArchitectureType_x86_64 ArchitectureType = 0
	ArchitectureType_arm64  ArchitectureType = 1
)

// Enum value maps for ArchitectureType.
var (
	ArchitectureType_name = map[int32]string{
		0: "x86_64",
		1: "arm64",
	}
	ArchitectureType_value = map[string]int32{
		"x86_64": 0,
		"arm64":  1,
	}
)

func (x ArchitectureType) Enum() *ArchitectureType {
	p := new(ArchitectureType)
	*p = x
	return p
}

func (x ArchitectureType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ArchitectureType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_protos_physicalDeviceDescription_proto_enumTypes[2].Descriptor()
}

func (ArchitectureType) Type() protoreflect.EnumType {
	return &file_pkg_protos_physicalDeviceDescription_proto_enumTypes[2]
}

func (x ArchitectureType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ArchitectureType.Descriptor instead.
func (ArchitectureType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_protos_physicalDeviceDescription_proto_rawDescGZIP(), []int{2}
}

type PhysicalDeviceDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SerialNumber    string               `protobuf:"bytes,1,opt,name=SerialNumber,proto3" json:"SerialNumber,omitempty"`
	DeviceName      string               `protobuf:"bytes,2,opt,name=DeviceName,proto3" json:"DeviceName,omitempty"`
	OSType          DeviceOSType         `protobuf:"varint,3,opt,name=OSType,proto3,enum=protos.DeviceOSType" json:"OSType,omitempty"`
	ArchType        ArchitectureType     `protobuf:"varint,4,opt,name=ArchType,proto3,enum=protos.ArchitectureType" json:"ArchType,omitempty"`
	ComputableChips []ComputableChipType `protobuf:"varint,5,rep,packed,name=ComputableChips,proto3,enum=protos.ComputableChipType" json:"ComputableChips,omitempty"`
	// service address
	//
	// Types that are assignable to DeviceAddress:
	//
	//	*PhysicalDeviceDescription_PcAddr
	//	*PhysicalDeviceDescription_AndroidAddr
	//	*PhysicalDeviceDescription_UsbAddr
	DeviceAddress isPhysicalDeviceDescription_DeviceAddress `protobuf_oneof:"DeviceAddress"`
}

func (x *PhysicalDeviceDescription) Reset() {
	*x = PhysicalDeviceDescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protos_physicalDeviceDescription_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhysicalDeviceDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhysicalDeviceDescription) ProtoMessage() {}

func (x *PhysicalDeviceDescription) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protos_physicalDeviceDescription_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhysicalDeviceDescription.ProtoReflect.Descriptor instead.
func (*PhysicalDeviceDescription) Descriptor() ([]byte, []int) {
	return file_pkg_protos_physicalDeviceDescription_proto_rawDescGZIP(), []int{0}
}

func (x *PhysicalDeviceDescription) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *PhysicalDeviceDescription) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *PhysicalDeviceDescription) GetOSType() DeviceOSType {
	if x != nil {
		return x.OSType
	}
	return DeviceOSType_none
}

func (x *PhysicalDeviceDescription) GetArchType() ArchitectureType {
	if x != nil {
		return x.ArchType
	}
	return ArchitectureType_x86_64
}

func (x *PhysicalDeviceDescription) GetComputableChips() []ComputableChipType {
	if x != nil {
		return x.ComputableChips
	}
	return nil
}

func (m *PhysicalDeviceDescription) GetDeviceAddress() isPhysicalDeviceDescription_DeviceAddress {
	if m != nil {
		return m.DeviceAddress
	}
	return nil
}

func (x *PhysicalDeviceDescription) GetPcAddr() *PCDeviceAddress {
	if x, ok := x.GetDeviceAddress().(*PhysicalDeviceDescription_PcAddr); ok {
		return x.PcAddr
	}
	return nil
}

func (x *PhysicalDeviceDescription) GetAndroidAddr() *AndroidDeviceAddress {
	if x, ok := x.GetDeviceAddress().(*PhysicalDeviceDescription_AndroidAddr); ok {
		return x.AndroidAddr
	}
	return nil
}

func (x *PhysicalDeviceDescription) GetUsbAddr() *UsbDeviceAddress {
	if x, ok := x.GetDeviceAddress().(*PhysicalDeviceDescription_UsbAddr); ok {
		return x.UsbAddr
	}
	return nil
}

type isPhysicalDeviceDescription_DeviceAddress interface {
	isPhysicalDeviceDescription_DeviceAddress()
}

type PhysicalDeviceDescription_PcAddr struct {
	PcAddr *PCDeviceAddress `protobuf:"bytes,21,opt,name=PcAddr,proto3,oneof"`
}

type PhysicalDeviceDescription_AndroidAddr struct {
	AndroidAddr *AndroidDeviceAddress `protobuf:"bytes,22,opt,name=AndroidAddr,proto3,oneof"`
}

type PhysicalDeviceDescription_UsbAddr struct {
	UsbAddr *UsbDeviceAddress `protobuf:"bytes,23,opt,name=UsbAddr,proto3,oneof"`
}

func (*PhysicalDeviceDescription_PcAddr) isPhysicalDeviceDescription_DeviceAddress() {}

func (*PhysicalDeviceDescription_AndroidAddr) isPhysicalDeviceDescription_DeviceAddress() {}

func (*PhysicalDeviceDescription_UsbAddr) isPhysicalDeviceDescription_DeviceAddress() {}

type PCDeviceAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceIp   string `protobuf:"bytes,1,opt,name=DeviceIp,proto3" json:"DeviceIp,omitempty"`
	DevicePort int32  `protobuf:"varint,2,opt,name=DevicePort,proto3" json:"DevicePort,omitempty"`
}

func (x *PCDeviceAddress) Reset() {
	*x = PCDeviceAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protos_physicalDeviceDescription_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PCDeviceAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PCDeviceAddress) ProtoMessage() {}

func (x *PCDeviceAddress) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protos_physicalDeviceDescription_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PCDeviceAddress.ProtoReflect.Descriptor instead.
func (*PCDeviceAddress) Descriptor() ([]byte, []int) {
	return file_pkg_protos_physicalDeviceDescription_proto_rawDescGZIP(), []int{1}
}

func (x *PCDeviceAddress) GetDeviceIp() string {
	if x != nil {
		return x.DeviceIp
	}
	return ""
}

func (x *PCDeviceAddress) GetDevicePort() int32 {
	if x != nil {
		return x.DevicePort
	}
	return 0
}

type AndroidDeviceAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentAddress *PCDeviceAddress `protobuf:"bytes,1,opt,name=parentAddress,proto3" json:"parentAddress,omitempty"`
	SerialNumber  string           `protobuf:"bytes,2,opt,name=SerialNumber,proto3" json:"SerialNumber,omitempty"`
	IpAddress     string           `protobuf:"bytes,3,opt,name=IpAddress,proto3" json:"IpAddress,omitempty"`
}

func (x *AndroidDeviceAddress) Reset() {
	*x = AndroidDeviceAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protos_physicalDeviceDescription_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AndroidDeviceAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AndroidDeviceAddress) ProtoMessage() {}

func (x *AndroidDeviceAddress) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protos_physicalDeviceDescription_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AndroidDeviceAddress.ProtoReflect.Descriptor instead.
func (*AndroidDeviceAddress) Descriptor() ([]byte, []int) {
	return file_pkg_protos_physicalDeviceDescription_proto_rawDescGZIP(), []int{2}
}

func (x *AndroidDeviceAddress) GetParentAddress() *PCDeviceAddress {
	if x != nil {
		return x.ParentAddress
	}
	return nil
}

func (x *AndroidDeviceAddress) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *AndroidDeviceAddress) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

type UsbDeviceAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentAddress *PCDeviceAddress `protobuf:"bytes,1,opt,name=parentAddress,proto3" json:"parentAddress,omitempty"`
	IpAddress     string           `protobuf:"bytes,2,opt,name=IpAddress,proto3" json:"IpAddress,omitempty"`
}

func (x *UsbDeviceAddress) Reset() {
	*x = UsbDeviceAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protos_physicalDeviceDescription_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsbDeviceAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsbDeviceAddress) ProtoMessage() {}

func (x *UsbDeviceAddress) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protos_physicalDeviceDescription_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsbDeviceAddress.ProtoReflect.Descriptor instead.
func (*UsbDeviceAddress) Descriptor() ([]byte, []int) {
	return file_pkg_protos_physicalDeviceDescription_proto_rawDescGZIP(), []int{3}
}

func (x *UsbDeviceAddress) GetParentAddress() *PCDeviceAddress {
	if x != nil {
		return x.ParentAddress
	}
	return nil
}

func (x *UsbDeviceAddress) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

var File_pkg_protos_physicalDeviceDescription_proto protoreflect.FileDescriptor

var file_pkg_protos_physicalDeviceDescription_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x68, 0x79,
	0x73, 0x69, 0x63, 0x61, 0x6c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x22, 0xc5, 0x03, 0x0a, 0x19, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61,
	0x6c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x4f, 0x53, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x53, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x4f, 0x53,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x41, 0x72, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x41, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x08, 0x41, 0x72, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x43, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x68, 0x69, 0x70, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x68, 0x69, 0x70, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x68, 0x69, 0x70, 0x73,
	0x12, 0x31, 0x0a, 0x06, 0x50, 0x63, 0x41, 0x64, 0x64, 0x72, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x43, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x48, 0x00, 0x52, 0x06, 0x50, 0x63, 0x41,
	0x64, 0x64, 0x72, 0x12, 0x40, 0x0a, 0x0b, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x41, 0x64,
	0x64, 0x72, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x48, 0x00, 0x52, 0x0b, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69,
	0x64, 0x41, 0x64, 0x64, 0x72, 0x12, 0x34, 0x0a, 0x07, 0x55, 0x73, 0x62, 0x41, 0x64, 0x64, 0x72,
	0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x55, 0x73, 0x62, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x48, 0x00, 0x52, 0x07, 0x55, 0x73, 0x62, 0x41, 0x64, 0x64, 0x72, 0x42, 0x0f, 0x0a, 0x0d, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x4d, 0x0a, 0x0f,
	0x50, 0x43, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x97, 0x01, 0x0a, 0x14,
	0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x3d, 0x0a, 0x0d, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x43, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x0d, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x53, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x70, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x70, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x6f, 0x0a, 0x10, 0x55, 0x73, 0x62, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3d, 0x0a, 0x0d, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x43, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0d, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x70, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x70, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2a, 0x4f, 0x0a, 0x0c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x4f, 0x53, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x6e, 0x6f, 0x6e, 0x65, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x10, 0x01, 0x12, 0x07, 0x0a,
	0x03, 0x69, 0x6f, 0x73, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x6d, 0x61, 0x63, 0x10, 0x03, 0x12,
	0x09, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x77, 0x69,
	0x6e, 0x64, 0x6f, 0x77, 0x73, 0x10, 0x05, 0x2a, 0x38, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x68, 0x69, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a,
	0x03, 0x63, 0x70, 0x75, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x67, 0x70, 0x75, 0x10, 0x01, 0x12,
	0x07, 0x0a, 0x03, 0x6e, 0x70, 0x75, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x74, 0x70, 0x75, 0x10,
	0x03, 0x2a, 0x29, 0x0a, 0x10, 0x41, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x78, 0x38, 0x36, 0x5f, 0x36, 0x34, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x61, 0x72, 0x6d, 0x36, 0x34, 0x10, 0x01, 0x42, 0x0b, 0x5a, 0x09,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_protos_physicalDeviceDescription_proto_rawDescOnce sync.Once
	file_pkg_protos_physicalDeviceDescription_proto_rawDescData = file_pkg_protos_physicalDeviceDescription_proto_rawDesc
)

func file_pkg_protos_physicalDeviceDescription_proto_rawDescGZIP() []byte {
	file_pkg_protos_physicalDeviceDescription_proto_rawDescOnce.Do(func() {
		file_pkg_protos_physicalDeviceDescription_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_protos_physicalDeviceDescription_proto_rawDescData)
	})
	return file_pkg_protos_physicalDeviceDescription_proto_rawDescData
}

var file_pkg_protos_physicalDeviceDescription_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_pkg_protos_physicalDeviceDescription_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_protos_physicalDeviceDescription_proto_goTypes = []interface{}{
	(DeviceOSType)(0),                 // 0: protos.DeviceOSType
	(ComputableChipType)(0),           // 1: protos.ComputableChipType
	(ArchitectureType)(0),             // 2: protos.ArchitectureType
	(*PhysicalDeviceDescription)(nil), // 3: protos.PhysicalDeviceDescription
	(*PCDeviceAddress)(nil),           // 4: protos.PCDeviceAddress
	(*AndroidDeviceAddress)(nil),      // 5: protos.AndroidDeviceAddress
	(*UsbDeviceAddress)(nil),          // 6: protos.UsbDeviceAddress
}
var file_pkg_protos_physicalDeviceDescription_proto_depIdxs = []int32{
	0, // 0: protos.PhysicalDeviceDescription.OSType:type_name -> protos.DeviceOSType
	2, // 1: protos.PhysicalDeviceDescription.ArchType:type_name -> protos.ArchitectureType
	1, // 2: protos.PhysicalDeviceDescription.ComputableChips:type_name -> protos.ComputableChipType
	4, // 3: protos.PhysicalDeviceDescription.PcAddr:type_name -> protos.PCDeviceAddress
	5, // 4: protos.PhysicalDeviceDescription.AndroidAddr:type_name -> protos.AndroidDeviceAddress
	6, // 5: protos.PhysicalDeviceDescription.UsbAddr:type_name -> protos.UsbDeviceAddress
	4, // 6: protos.AndroidDeviceAddress.parentAddress:type_name -> protos.PCDeviceAddress
	4, // 7: protos.UsbDeviceAddress.parentAddress:type_name -> protos.PCDeviceAddress
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_pkg_protos_physicalDeviceDescription_proto_init() }
func file_pkg_protos_physicalDeviceDescription_proto_init() {
	if File_pkg_protos_physicalDeviceDescription_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_protos_physicalDeviceDescription_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhysicalDeviceDescription); i {
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
		file_pkg_protos_physicalDeviceDescription_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PCDeviceAddress); i {
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
		file_pkg_protos_physicalDeviceDescription_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AndroidDeviceAddress); i {
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
		file_pkg_protos_physicalDeviceDescription_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsbDeviceAddress); i {
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
	file_pkg_protos_physicalDeviceDescription_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PhysicalDeviceDescription_PcAddr)(nil),
		(*PhysicalDeviceDescription_AndroidAddr)(nil),
		(*PhysicalDeviceDescription_UsbAddr)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_protos_physicalDeviceDescription_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_protos_physicalDeviceDescription_proto_goTypes,
		DependencyIndexes: file_pkg_protos_physicalDeviceDescription_proto_depIdxs,
		EnumInfos:         file_pkg_protos_physicalDeviceDescription_proto_enumTypes,
		MessageInfos:      file_pkg_protos_physicalDeviceDescription_proto_msgTypes,
	}.Build()
	File_pkg_protos_physicalDeviceDescription_proto = out.File
	file_pkg_protos_physicalDeviceDescription_proto_rawDesc = nil
	file_pkg_protos_physicalDeviceDescription_proto_goTypes = nil
	file_pkg_protos_physicalDeviceDescription_proto_depIdxs = nil
}
