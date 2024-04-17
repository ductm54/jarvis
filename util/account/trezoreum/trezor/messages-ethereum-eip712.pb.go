// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: messages-ethereum-eip712.proto

package trezor

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

type EthereumTypedDataStructAck_EthereumDataType int32

const (
	EthereumTypedDataStructAck_UINT    EthereumTypedDataStructAck_EthereumDataType = 1
	EthereumTypedDataStructAck_INT     EthereumTypedDataStructAck_EthereumDataType = 2
	EthereumTypedDataStructAck_BYTES   EthereumTypedDataStructAck_EthereumDataType = 3
	EthereumTypedDataStructAck_STRING  EthereumTypedDataStructAck_EthereumDataType = 4
	EthereumTypedDataStructAck_BOOL    EthereumTypedDataStructAck_EthereumDataType = 5
	EthereumTypedDataStructAck_ADDRESS EthereumTypedDataStructAck_EthereumDataType = 6
	EthereumTypedDataStructAck_ARRAY   EthereumTypedDataStructAck_EthereumDataType = 7
	EthereumTypedDataStructAck_STRUCT  EthereumTypedDataStructAck_EthereumDataType = 8
)

// Enum value maps for EthereumTypedDataStructAck_EthereumDataType.
var (
	EthereumTypedDataStructAck_EthereumDataType_name = map[int32]string{
		1: "UINT",
		2: "INT",
		3: "BYTES",
		4: "STRING",
		5: "BOOL",
		6: "ADDRESS",
		7: "ARRAY",
		8: "STRUCT",
	}
	EthereumTypedDataStructAck_EthereumDataType_value = map[string]int32{
		"UINT":    1,
		"INT":     2,
		"BYTES":   3,
		"STRING":  4,
		"BOOL":    5,
		"ADDRESS": 6,
		"ARRAY":   7,
		"STRUCT":  8,
	}
)

func (x EthereumTypedDataStructAck_EthereumDataType) Enum() *EthereumTypedDataStructAck_EthereumDataType {
	p := new(EthereumTypedDataStructAck_EthereumDataType)
	*p = x
	return p
}

func (x EthereumTypedDataStructAck_EthereumDataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EthereumTypedDataStructAck_EthereumDataType) Descriptor() protoreflect.EnumDescriptor {
	return file_messages_ethereum_eip712_proto_enumTypes[0].Descriptor()
}

func (EthereumTypedDataStructAck_EthereumDataType) Type() protoreflect.EnumType {
	return &file_messages_ethereum_eip712_proto_enumTypes[0]
}

func (x EthereumTypedDataStructAck_EthereumDataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EthereumTypedDataStructAck_EthereumDataType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EthereumTypedDataStructAck_EthereumDataType(num)
	return nil
}

// Deprecated: Use EthereumTypedDataStructAck_EthereumDataType.Descriptor instead.
func (EthereumTypedDataStructAck_EthereumDataType) EnumDescriptor() ([]byte, []int) {
	return file_messages_ethereum_eip712_proto_rawDescGZIP(), []int{2, 0}
}

// *
// Request: Ask device to sign typed data
// @start
// @next EthereumTypedDataStructRequest
// @next EthereumTypedDataValueRequest
// @next EthereumTypedDataSignature
// @next Failure
type EthereumSignTypedData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddressN         []uint32             `protobuf:"varint,1,rep,name=address_n,json=addressN" json:"address_n,omitempty"`                                 // BIP-32 path to derive the key from master node
	PrimaryType      *string              `protobuf:"bytes,2,req,name=primary_type,json=primaryType" json:"primary_type,omitempty"`                         // name of the root message struct
	MetamaskV4Compat *bool                `protobuf:"varint,3,opt,name=metamask_v4_compat,json=metamaskV4Compat,def=1" json:"metamask_v4_compat,omitempty"` // use MetaMask v4 (see https://github.com/MetaMask/eth-sig-util/issues/106)
	Definitions      *EthereumDefinitions `protobuf:"bytes,4,opt,name=definitions" json:"definitions,omitempty"`                                            // network and/or token definitions
}

// Default values for EthereumSignTypedData fields.
const (
	Default_EthereumSignTypedData_MetamaskV4Compat = bool(true)
)

func (x *EthereumSignTypedData) Reset() {
	*x = EthereumSignTypedData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_ethereum_eip712_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthereumSignTypedData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthereumSignTypedData) ProtoMessage() {}

func (x *EthereumSignTypedData) ProtoReflect() protoreflect.Message {
	mi := &file_messages_ethereum_eip712_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthereumSignTypedData.ProtoReflect.Descriptor instead.
func (*EthereumSignTypedData) Descriptor() ([]byte, []int) {
	return file_messages_ethereum_eip712_proto_rawDescGZIP(), []int{0}
}

func (x *EthereumSignTypedData) GetAddressN() []uint32 {
	if x != nil {
		return x.AddressN
	}
	return nil
}

func (x *EthereumSignTypedData) GetPrimaryType() string {
	if x != nil && x.PrimaryType != nil {
		return *x.PrimaryType
	}
	return ""
}

func (x *EthereumSignTypedData) GetMetamaskV4Compat() bool {
	if x != nil && x.MetamaskV4Compat != nil {
		return *x.MetamaskV4Compat
	}
	return Default_EthereumSignTypedData_MetamaskV4Compat
}

func (x *EthereumSignTypedData) GetDefinitions() *EthereumDefinitions {
	if x != nil {
		return x.Definitions
	}
	return nil
}

// *
// Response: Device asks for type information about a struct.
// @next EthereumTypedDataStructAck
type EthereumTypedDataStructRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"` // name of the requested struct
}

func (x *EthereumTypedDataStructRequest) Reset() {
	*x = EthereumTypedDataStructRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_ethereum_eip712_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthereumTypedDataStructRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthereumTypedDataStructRequest) ProtoMessage() {}

func (x *EthereumTypedDataStructRequest) ProtoReflect() protoreflect.Message {
	mi := &file_messages_ethereum_eip712_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthereumTypedDataStructRequest.ProtoReflect.Descriptor instead.
func (*EthereumTypedDataStructRequest) Descriptor() ([]byte, []int) {
	return file_messages_ethereum_eip712_proto_rawDescGZIP(), []int{1}
}

func (x *EthereumTypedDataStructRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

// *
// Request: Type information about a struct.
// @next EthereumTypedDataStructRequest
type EthereumTypedDataStructAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Members []*EthereumTypedDataStructAck_EthereumStructMember `protobuf:"bytes,1,rep,name=members" json:"members,omitempty"`
}

func (x *EthereumTypedDataStructAck) Reset() {
	*x = EthereumTypedDataStructAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_ethereum_eip712_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthereumTypedDataStructAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthereumTypedDataStructAck) ProtoMessage() {}

func (x *EthereumTypedDataStructAck) ProtoReflect() protoreflect.Message {
	mi := &file_messages_ethereum_eip712_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthereumTypedDataStructAck.ProtoReflect.Descriptor instead.
func (*EthereumTypedDataStructAck) Descriptor() ([]byte, []int) {
	return file_messages_ethereum_eip712_proto_rawDescGZIP(), []int{2}
}

func (x *EthereumTypedDataStructAck) GetMembers() []*EthereumTypedDataStructAck_EthereumStructMember {
	if x != nil {
		return x.Members
	}
	return nil
}

// *
// Response: Device asks for data at the specific member path.
// @next EthereumTypedDataValueAck
type EthereumTypedDataValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberPath []uint32 `protobuf:"varint,1,rep,name=member_path,json=memberPath" json:"member_path,omitempty"` // member path requested by device
}

func (x *EthereumTypedDataValueRequest) Reset() {
	*x = EthereumTypedDataValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_ethereum_eip712_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthereumTypedDataValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthereumTypedDataValueRequest) ProtoMessage() {}

func (x *EthereumTypedDataValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_messages_ethereum_eip712_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthereumTypedDataValueRequest.ProtoReflect.Descriptor instead.
func (*EthereumTypedDataValueRequest) Descriptor() ([]byte, []int) {
	return file_messages_ethereum_eip712_proto_rawDescGZIP(), []int{3}
}

func (x *EthereumTypedDataValueRequest) GetMemberPath() []uint32 {
	if x != nil {
		return x.MemberPath
	}
	return nil
}

// *
// Request: Single value of a specific atomic field.
// @next EthereumTypedDataValueRequest
type EthereumTypedDataValueAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,req,name=value" json:"value,omitempty"`
}

func (x *EthereumTypedDataValueAck) Reset() {
	*x = EthereumTypedDataValueAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_ethereum_eip712_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthereumTypedDataValueAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthereumTypedDataValueAck) ProtoMessage() {}

func (x *EthereumTypedDataValueAck) ProtoReflect() protoreflect.Message {
	mi := &file_messages_ethereum_eip712_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthereumTypedDataValueAck.ProtoReflect.Descriptor instead.
func (*EthereumTypedDataValueAck) Descriptor() ([]byte, []int) {
	return file_messages_ethereum_eip712_proto_rawDescGZIP(), []int{4}
}

func (x *EthereumTypedDataValueAck) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type EthereumTypedDataStructAck_EthereumStructMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type *EthereumTypedDataStructAck_EthereumFieldType `protobuf:"bytes,1,req,name=type" json:"type,omitempty"`
	Name *string                                       `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
}

func (x *EthereumTypedDataStructAck_EthereumStructMember) Reset() {
	*x = EthereumTypedDataStructAck_EthereumStructMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_ethereum_eip712_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthereumTypedDataStructAck_EthereumStructMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthereumTypedDataStructAck_EthereumStructMember) ProtoMessage() {}

func (x *EthereumTypedDataStructAck_EthereumStructMember) ProtoReflect() protoreflect.Message {
	mi := &file_messages_ethereum_eip712_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthereumTypedDataStructAck_EthereumStructMember.ProtoReflect.Descriptor instead.
func (*EthereumTypedDataStructAck_EthereumStructMember) Descriptor() ([]byte, []int) {
	return file_messages_ethereum_eip712_proto_rawDescGZIP(), []int{2, 0}
}

func (x *EthereumTypedDataStructAck_EthereumStructMember) GetType() *EthereumTypedDataStructAck_EthereumFieldType {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *EthereumTypedDataStructAck_EthereumStructMember) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type EthereumTypedDataStructAck_EthereumFieldType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataType *EthereumTypedDataStructAck_EthereumDataType `protobuf:"varint,1,req,name=data_type,json=dataType,enum=hw.trezor.messages.ethereum_eip712.EthereumTypedDataStructAck_EthereumDataType" json:"data_type,omitempty"`
	Size     *uint32                                      `protobuf:"varint,2,opt,name=size" json:"size,omitempty"` // for integer types: size in bytes (uint8 has size 1, uint256 has size 32)
	// for bytes types: size in bytes, or unset for dynamic
	// for arrays: size in elements, or unset for dynamic
	// for structs: number of members
	// for string, bool and address: unset
	EntryType  *EthereumTypedDataStructAck_EthereumFieldType `protobuf:"bytes,3,opt,name=entry_type,json=entryType" json:"entry_type,omitempty"`    // for array types, type of single entry
	StructName *string                                       `protobuf:"bytes,4,opt,name=struct_name,json=structName" json:"struct_name,omitempty"` // for structs: its name
}

func (x *EthereumTypedDataStructAck_EthereumFieldType) Reset() {
	*x = EthereumTypedDataStructAck_EthereumFieldType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_ethereum_eip712_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthereumTypedDataStructAck_EthereumFieldType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthereumTypedDataStructAck_EthereumFieldType) ProtoMessage() {}

func (x *EthereumTypedDataStructAck_EthereumFieldType) ProtoReflect() protoreflect.Message {
	mi := &file_messages_ethereum_eip712_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthereumTypedDataStructAck_EthereumFieldType.ProtoReflect.Descriptor instead.
func (*EthereumTypedDataStructAck_EthereumFieldType) Descriptor() ([]byte, []int) {
	return file_messages_ethereum_eip712_proto_rawDescGZIP(), []int{2, 1}
}

func (x *EthereumTypedDataStructAck_EthereumFieldType) GetDataType() EthereumTypedDataStructAck_EthereumDataType {
	if x != nil && x.DataType != nil {
		return *x.DataType
	}
	return EthereumTypedDataStructAck_UINT
}

func (x *EthereumTypedDataStructAck_EthereumFieldType) GetSize() uint32 {
	if x != nil && x.Size != nil {
		return *x.Size
	}
	return 0
}

func (x *EthereumTypedDataStructAck_EthereumFieldType) GetEntryType() *EthereumTypedDataStructAck_EthereumFieldType {
	if x != nil {
		return x.EntryType
	}
	return nil
}

func (x *EthereumTypedDataStructAck_EthereumFieldType) GetStructName() string {
	if x != nil && x.StructName != nil {
		return *x.StructName
	}
	return ""
}

var File_messages_ethereum_eip712_proto protoreflect.FileDescriptor

var file_messages_ethereum_eip712_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2d, 0x65, 0x74, 0x68, 0x65, 0x72,
	0x65, 0x75, 0x6d, 0x2d, 0x65, 0x69, 0x70, 0x37, 0x31, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x22, 0x68, 0x77, 0x2e, 0x74, 0x72, 0x65, 0x7a, 0x6f, 0x72, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5f, 0x65, 0x69,
	0x70, 0x37, 0x31, 0x32, 0x1a, 0x23, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2d, 0x65,
	0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2d, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb, 0x01, 0x0a, 0x15, 0x45, 0x74,
	0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x53, 0x69, 0x67, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6e,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4e,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x12, 0x6d, 0x65, 0x74, 0x61, 0x6d, 0x61, 0x73, 0x6b, 0x5f,
	0x76, 0x34, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x3a,
	0x04, 0x74, 0x72, 0x75, 0x65, 0x52, 0x10, 0x6d, 0x65, 0x74, 0x61, 0x6d, 0x61, 0x73, 0x6b, 0x56,
	0x34, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x12, 0x5e, 0x0a, 0x0b, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x68,
	0x77, 0x2e, 0x74, 0x72, 0x65, 0x7a, 0x6f, 0x72, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x44,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x34, 0x0a, 0x1e, 0x45, 0x74, 0x68, 0x65, 0x72,
	0x65, 0x75, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xb4, 0x05,
	0x0a, 0x1a, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x64, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x41, 0x63, 0x6b, 0x12, 0x6d, 0x0a, 0x07,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x53, 0x2e,
	0x68, 0x77, 0x2e, 0x74, 0x72, 0x65, 0x7a, 0x6f, 0x72, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5f, 0x65, 0x69, 0x70, 0x37,
	0x31, 0x32, 0x2e, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x64,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x41, 0x63, 0x6b, 0x2e, 0x45, 0x74,
	0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x1a, 0x90, 0x01, 0x0a, 0x14,
	0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x64, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x0b, 0x32, 0x50, 0x2e, 0x68, 0x77, 0x2e, 0x74, 0x72, 0x65, 0x7a, 0x6f, 0x72, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d,
	0x5f, 0x65, 0x69, 0x70, 0x37, 0x31, 0x32, 0x2e, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d,
	0x54, 0x79, 0x70, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x41,
	0x63, 0x6b, 0x2e, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0xa7,
	0x02, 0x0a, 0x11, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x6c, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x4f, 0x2e, 0x68, 0x77, 0x2e, 0x74, 0x72, 0x65,
	0x7a, 0x6f, 0x72, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x65, 0x74, 0x68,
	0x65, 0x72, 0x65, 0x75, 0x6d, 0x5f, 0x65, 0x69, 0x70, 0x37, 0x31, 0x32, 0x2e, 0x45, 0x74, 0x68,
	0x65, 0x72, 0x65, 0x75, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x41, 0x63, 0x6b, 0x2e, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d,
	0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x6f, 0x0a, 0x0a, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x50, 0x2e, 0x68, 0x77, 0x2e,
	0x74, 0x72, 0x65, 0x7a, 0x6f, 0x72, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5f, 0x65, 0x69, 0x70, 0x37, 0x31, 0x32, 0x2e,
	0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x41, 0x63, 0x6b, 0x2e, 0x45, 0x74, 0x68, 0x65, 0x72,
	0x65, 0x75, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x6a, 0x0a, 0x10, 0x45, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04,
	0x55, 0x49, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x4e, 0x54, 0x10, 0x02, 0x12,
	0x09, 0x0a, 0x05, 0x42, 0x59, 0x54, 0x45, 0x53, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54,
	0x52, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x4f, 0x4f, 0x4c, 0x10, 0x05,
	0x12, 0x0b, 0x0a, 0x07, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x10, 0x06, 0x12, 0x09, 0x0a,
	0x05, 0x41, 0x52, 0x52, 0x41, 0x59, 0x10, 0x07, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x55,
	0x43, 0x54, 0x10, 0x08, 0x22, 0x40, 0x0a, 0x1d, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d,
	0x54, 0x79, 0x70, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x22, 0x31, 0x0a, 0x19, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65,
	0x75, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x41, 0x63, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x5f, 0x0a, 0x23, 0x63, 0x6f, 0x6d,
	0x2e, 0x73, 0x61, 0x74, 0x6f, 0x73, 0x68, 0x69, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x74, 0x72, 0x65,
	0x7a, 0x6f, 0x72, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x42, 0x1b, 0x54, 0x72, 0x65, 0x7a, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45,
	0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x45, 0x49, 0x50, 0x37, 0x31, 0x32, 0x5a, 0x1b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x74, 0x72, 0x65,
	0x7a, 0x6f, 0x72, 0x2f, 0x74, 0x72, 0x65, 0x7a, 0x6f, 0x72,
}

var (
	file_messages_ethereum_eip712_proto_rawDescOnce sync.Once
	file_messages_ethereum_eip712_proto_rawDescData = file_messages_ethereum_eip712_proto_rawDesc
)

func file_messages_ethereum_eip712_proto_rawDescGZIP() []byte {
	file_messages_ethereum_eip712_proto_rawDescOnce.Do(func() {
		file_messages_ethereum_eip712_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_ethereum_eip712_proto_rawDescData)
	})
	return file_messages_ethereum_eip712_proto_rawDescData
}

var file_messages_ethereum_eip712_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_messages_ethereum_eip712_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_messages_ethereum_eip712_proto_goTypes = []interface{}{
	(EthereumTypedDataStructAck_EthereumDataType)(0),        // 0: hw.trezor.messages.ethereum_eip712.EthereumTypedDataStructAck.EthereumDataType
	(*EthereumSignTypedData)(nil),                           // 1: hw.trezor.messages.ethereum_eip712.EthereumSignTypedData
	(*EthereumTypedDataStructRequest)(nil),                  // 2: hw.trezor.messages.ethereum_eip712.EthereumTypedDataStructRequest
	(*EthereumTypedDataStructAck)(nil),                      // 3: hw.trezor.messages.ethereum_eip712.EthereumTypedDataStructAck
	(*EthereumTypedDataValueRequest)(nil),                   // 4: hw.trezor.messages.ethereum_eip712.EthereumTypedDataValueRequest
	(*EthereumTypedDataValueAck)(nil),                       // 5: hw.trezor.messages.ethereum_eip712.EthereumTypedDataValueAck
	(*EthereumTypedDataStructAck_EthereumStructMember)(nil), // 6: hw.trezor.messages.ethereum_eip712.EthereumTypedDataStructAck.EthereumStructMember
	(*EthereumTypedDataStructAck_EthereumFieldType)(nil),    // 7: hw.trezor.messages.ethereum_eip712.EthereumTypedDataStructAck.EthereumFieldType
	(*EthereumDefinitions)(nil),                             // 8: hw.trezor.messages.ethereum_definitions.EthereumDefinitions
}
var file_messages_ethereum_eip712_proto_depIdxs = []int32{
	8, // 0: hw.trezor.messages.ethereum_eip712.EthereumSignTypedData.definitions:type_name -> hw.trezor.messages.ethereum_definitions.EthereumDefinitions
	6, // 1: hw.trezor.messages.ethereum_eip712.EthereumTypedDataStructAck.members:type_name -> hw.trezor.messages.ethereum_eip712.EthereumTypedDataStructAck.EthereumStructMember
	7, // 2: hw.trezor.messages.ethereum_eip712.EthereumTypedDataStructAck.EthereumStructMember.type:type_name -> hw.trezor.messages.ethereum_eip712.EthereumTypedDataStructAck.EthereumFieldType
	0, // 3: hw.trezor.messages.ethereum_eip712.EthereumTypedDataStructAck.EthereumFieldType.data_type:type_name -> hw.trezor.messages.ethereum_eip712.EthereumTypedDataStructAck.EthereumDataType
	7, // 4: hw.trezor.messages.ethereum_eip712.EthereumTypedDataStructAck.EthereumFieldType.entry_type:type_name -> hw.trezor.messages.ethereum_eip712.EthereumTypedDataStructAck.EthereumFieldType
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_messages_ethereum_eip712_proto_init() }
func file_messages_ethereum_eip712_proto_init() {
	if File_messages_ethereum_eip712_proto != nil {
		return
	}
	file_messages_ethereum_definitions_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_messages_ethereum_eip712_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EthereumSignTypedData); i {
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
		file_messages_ethereum_eip712_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EthereumTypedDataStructRequest); i {
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
		file_messages_ethereum_eip712_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EthereumTypedDataStructAck); i {
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
		file_messages_ethereum_eip712_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EthereumTypedDataValueRequest); i {
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
		file_messages_ethereum_eip712_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EthereumTypedDataValueAck); i {
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
		file_messages_ethereum_eip712_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EthereumTypedDataStructAck_EthereumStructMember); i {
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
		file_messages_ethereum_eip712_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EthereumTypedDataStructAck_EthereumFieldType); i {
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
			RawDescriptor: file_messages_ethereum_eip712_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_ethereum_eip712_proto_goTypes,
		DependencyIndexes: file_messages_ethereum_eip712_proto_depIdxs,
		EnumInfos:         file_messages_ethereum_eip712_proto_enumTypes,
		MessageInfos:      file_messages_ethereum_eip712_proto_msgTypes,
	}.Build()
	File_messages_ethereum_eip712_proto = out.File
	file_messages_ethereum_eip712_proto_rawDesc = nil
	file_messages_ethereum_eip712_proto_goTypes = nil
	file_messages_ethereum_eip712_proto_depIdxs = nil
}
