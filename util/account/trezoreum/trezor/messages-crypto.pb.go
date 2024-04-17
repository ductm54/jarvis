// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: messages-crypto.proto

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

// *
// Request: Ask device to encrypt or decrypt value of given key
// @start
// @next CipheredKeyValue
// @next Failure
type CipherKeyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddressN     []uint32 `protobuf:"varint,1,rep,name=address_n,json=addressN" json:"address_n,omitempty"`               // BIP-32 path to derive the key from master node
	Key          *string  `protobuf:"bytes,2,req,name=key" json:"key,omitempty"`                                          // key component of key:value
	Value        []byte   `protobuf:"bytes,3,req,name=value" json:"value,omitempty"`                                      // value component of key:value
	Encrypt      *bool    `protobuf:"varint,4,opt,name=encrypt" json:"encrypt,omitempty"`                                 // are we encrypting (True) or decrypting (False)?
	AskOnEncrypt *bool    `protobuf:"varint,5,opt,name=ask_on_encrypt,json=askOnEncrypt" json:"ask_on_encrypt,omitempty"` // should we ask on encrypt operation?
	AskOnDecrypt *bool    `protobuf:"varint,6,opt,name=ask_on_decrypt,json=askOnDecrypt" json:"ask_on_decrypt,omitempty"` // should we ask on decrypt operation?
	Iv           []byte   `protobuf:"bytes,7,opt,name=iv" json:"iv,omitempty"`                                            // initialization vector (will be computed if not set)
}

func (x *CipherKeyValue) Reset() {
	*x = CipherKeyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_crypto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CipherKeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CipherKeyValue) ProtoMessage() {}

func (x *CipherKeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_messages_crypto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CipherKeyValue.ProtoReflect.Descriptor instead.
func (*CipherKeyValue) Descriptor() ([]byte, []int) {
	return file_messages_crypto_proto_rawDescGZIP(), []int{0}
}

func (x *CipherKeyValue) GetAddressN() []uint32 {
	if x != nil {
		return x.AddressN
	}
	return nil
}

func (x *CipherKeyValue) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *CipherKeyValue) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *CipherKeyValue) GetEncrypt() bool {
	if x != nil && x.Encrypt != nil {
		return *x.Encrypt
	}
	return false
}

func (x *CipherKeyValue) GetAskOnEncrypt() bool {
	if x != nil && x.AskOnEncrypt != nil {
		return *x.AskOnEncrypt
	}
	return false
}

func (x *CipherKeyValue) GetAskOnDecrypt() bool {
	if x != nil && x.AskOnDecrypt != nil {
		return *x.AskOnDecrypt
	}
	return false
}

func (x *CipherKeyValue) GetIv() []byte {
	if x != nil {
		return x.Iv
	}
	return nil
}

// *
// Response: Return ciphered/deciphered value
// @end
type CipheredKeyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,req,name=value" json:"value,omitempty"` // ciphered/deciphered value
}

func (x *CipheredKeyValue) Reset() {
	*x = CipheredKeyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_crypto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CipheredKeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CipheredKeyValue) ProtoMessage() {}

func (x *CipheredKeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_messages_crypto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CipheredKeyValue.ProtoReflect.Descriptor instead.
func (*CipheredKeyValue) Descriptor() ([]byte, []int) {
	return file_messages_crypto_proto_rawDescGZIP(), []int{1}
}

func (x *CipheredKeyValue) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

// *
// Structure representing identity data
// @embed
type IdentityType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Proto *string `protobuf:"bytes,1,opt,name=proto" json:"proto,omitempty"`        // proto part of URI
	User  *string `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`          // user part of URI
	Host  *string `protobuf:"bytes,3,opt,name=host" json:"host,omitempty"`          // host part of URI
	Port  *string `protobuf:"bytes,4,opt,name=port" json:"port,omitempty"`          // port part of URI
	Path  *string `protobuf:"bytes,5,opt,name=path" json:"path,omitempty"`          // path part of URI
	Index *uint32 `protobuf:"varint,6,opt,name=index,def=0" json:"index,omitempty"` // identity index
}

// Default values for IdentityType fields.
const (
	Default_IdentityType_Index = uint32(0)
)

func (x *IdentityType) Reset() {
	*x = IdentityType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_crypto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityType) ProtoMessage() {}

func (x *IdentityType) ProtoReflect() protoreflect.Message {
	mi := &file_messages_crypto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityType.ProtoReflect.Descriptor instead.
func (*IdentityType) Descriptor() ([]byte, []int) {
	return file_messages_crypto_proto_rawDescGZIP(), []int{2}
}

func (x *IdentityType) GetProto() string {
	if x != nil && x.Proto != nil {
		return *x.Proto
	}
	return ""
}

func (x *IdentityType) GetUser() string {
	if x != nil && x.User != nil {
		return *x.User
	}
	return ""
}

func (x *IdentityType) GetHost() string {
	if x != nil && x.Host != nil {
		return *x.Host
	}
	return ""
}

func (x *IdentityType) GetPort() string {
	if x != nil && x.Port != nil {
		return *x.Port
	}
	return ""
}

func (x *IdentityType) GetPath() string {
	if x != nil && x.Path != nil {
		return *x.Path
	}
	return ""
}

func (x *IdentityType) GetIndex() uint32 {
	if x != nil && x.Index != nil {
		return *x.Index
	}
	return Default_IdentityType_Index
}

// *
// Request: Ask device to sign identity
// @start
// @next SignedIdentity
// @next Failure
type SignIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity        *IdentityType `protobuf:"bytes,1,req,name=identity" json:"identity,omitempty"`                                           // identity
	ChallengeHidden []byte        `protobuf:"bytes,2,opt,name=challenge_hidden,json=challengeHidden,def=" json:"challenge_hidden,omitempty"` // non-visible challenge
	ChallengeVisual *string       `protobuf:"bytes,3,opt,name=challenge_visual,json=challengeVisual,def=" json:"challenge_visual,omitempty"` // challenge shown on display (e.g. date+time)
	EcdsaCurveName  *string       `protobuf:"bytes,4,opt,name=ecdsa_curve_name,json=ecdsaCurveName" json:"ecdsa_curve_name,omitempty"`       // ECDSA curve name to use
}

// Default values for SignIdentity fields.
const (
	Default_SignIdentity_ChallengeVisual = string("")
)

// Default values for SignIdentity fields.
var (
	Default_SignIdentity_ChallengeHidden = []byte("")
)

func (x *SignIdentity) Reset() {
	*x = SignIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_crypto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignIdentity) ProtoMessage() {}

func (x *SignIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_messages_crypto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignIdentity.ProtoReflect.Descriptor instead.
func (*SignIdentity) Descriptor() ([]byte, []int) {
	return file_messages_crypto_proto_rawDescGZIP(), []int{3}
}

func (x *SignIdentity) GetIdentity() *IdentityType {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *SignIdentity) GetChallengeHidden() []byte {
	if x != nil && x.ChallengeHidden != nil {
		return x.ChallengeHidden
	}
	return append([]byte(nil), Default_SignIdentity_ChallengeHidden...)
}

func (x *SignIdentity) GetChallengeVisual() string {
	if x != nil && x.ChallengeVisual != nil {
		return *x.ChallengeVisual
	}
	return Default_SignIdentity_ChallengeVisual
}

func (x *SignIdentity) GetEcdsaCurveName() string {
	if x != nil && x.EcdsaCurveName != nil {
		return *x.EcdsaCurveName
	}
	return ""
}

// *
// Response: Device provides signed identity
// @end
type SignedIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address   *string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`                      // identity address
	PublicKey []byte  `protobuf:"bytes,2,req,name=public_key,json=publicKey" json:"public_key,omitempty"` // identity public key
	Signature []byte  `protobuf:"bytes,3,req,name=signature" json:"signature,omitempty"`                  // signature of the identity data
}

func (x *SignedIdentity) Reset() {
	*x = SignedIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_crypto_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedIdentity) ProtoMessage() {}

func (x *SignedIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_messages_crypto_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedIdentity.ProtoReflect.Descriptor instead.
func (*SignedIdentity) Descriptor() ([]byte, []int) {
	return file_messages_crypto_proto_rawDescGZIP(), []int{4}
}

func (x *SignedIdentity) GetAddress() string {
	if x != nil && x.Address != nil {
		return *x.Address
	}
	return ""
}

func (x *SignedIdentity) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *SignedIdentity) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

// *
// Request: Ask device to generate ECDH session key
// @start
// @next ECDHSessionKey
// @next Failure
type GetECDHSessionKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity       *IdentityType `protobuf:"bytes,1,req,name=identity" json:"identity,omitempty"`                                     // identity
	PeerPublicKey  []byte        `protobuf:"bytes,2,req,name=peer_public_key,json=peerPublicKey" json:"peer_public_key,omitempty"`    // peer's public key
	EcdsaCurveName *string       `protobuf:"bytes,3,opt,name=ecdsa_curve_name,json=ecdsaCurveName" json:"ecdsa_curve_name,omitempty"` // ECDSA curve name to use
}

func (x *GetECDHSessionKey) Reset() {
	*x = GetECDHSessionKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_crypto_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetECDHSessionKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetECDHSessionKey) ProtoMessage() {}

func (x *GetECDHSessionKey) ProtoReflect() protoreflect.Message {
	mi := &file_messages_crypto_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetECDHSessionKey.ProtoReflect.Descriptor instead.
func (*GetECDHSessionKey) Descriptor() ([]byte, []int) {
	return file_messages_crypto_proto_rawDescGZIP(), []int{5}
}

func (x *GetECDHSessionKey) GetIdentity() *IdentityType {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *GetECDHSessionKey) GetPeerPublicKey() []byte {
	if x != nil {
		return x.PeerPublicKey
	}
	return nil
}

func (x *GetECDHSessionKey) GetEcdsaCurveName() string {
	if x != nil && x.EcdsaCurveName != nil {
		return *x.EcdsaCurveName
	}
	return ""
}

// *
// Response: Device provides ECDH session key
// @end
type ECDHSessionKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionKey []byte `protobuf:"bytes,1,req,name=session_key,json=sessionKey" json:"session_key,omitempty"` // ECDH session key
	PublicKey  []byte `protobuf:"bytes,2,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`    // identity public key
}

func (x *ECDHSessionKey) Reset() {
	*x = ECDHSessionKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_crypto_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ECDHSessionKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ECDHSessionKey) ProtoMessage() {}

func (x *ECDHSessionKey) ProtoReflect() protoreflect.Message {
	mi := &file_messages_crypto_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ECDHSessionKey.ProtoReflect.Descriptor instead.
func (*ECDHSessionKey) Descriptor() ([]byte, []int) {
	return file_messages_crypto_proto_rawDescGZIP(), []int{6}
}

func (x *ECDHSessionKey) GetSessionKey() []byte {
	if x != nil {
		return x.SessionKey
	}
	return nil
}

func (x *ECDHSessionKey) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

// *
// Request: Ask device to commit to CoSi signing
// @start
// @next CosiCommitment
// @next Failure
type CosiCommit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddressN []uint32 `protobuf:"varint,1,rep,name=address_n,json=addressN" json:"address_n,omitempty"` // BIP-32 path to derive the key from master node
	// Deprecated: Marked as deprecated in messages-crypto.proto.
	Data []byte `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"` // Data to be signed. Deprecated in 1.10.2, the field is not needed, since CoSi commitments are no longer deterministic.
}

func (x *CosiCommit) Reset() {
	*x = CosiCommit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_crypto_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CosiCommit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CosiCommit) ProtoMessage() {}

func (x *CosiCommit) ProtoReflect() protoreflect.Message {
	mi := &file_messages_crypto_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CosiCommit.ProtoReflect.Descriptor instead.
func (*CosiCommit) Descriptor() ([]byte, []int) {
	return file_messages_crypto_proto_rawDescGZIP(), []int{7}
}

func (x *CosiCommit) GetAddressN() []uint32 {
	if x != nil {
		return x.AddressN
	}
	return nil
}

// Deprecated: Marked as deprecated in messages-crypto.proto.
func (x *CosiCommit) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// *
// Response: Contains a CoSi commitment
// @end
type CosiCommitment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commitment []byte `protobuf:"bytes,1,req,name=commitment" json:"commitment,omitempty"` // Commitment
	Pubkey     []byte `protobuf:"bytes,2,req,name=pubkey" json:"pubkey,omitempty"`         // Public key
}

func (x *CosiCommitment) Reset() {
	*x = CosiCommitment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_crypto_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CosiCommitment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CosiCommitment) ProtoMessage() {}

func (x *CosiCommitment) ProtoReflect() protoreflect.Message {
	mi := &file_messages_crypto_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CosiCommitment.ProtoReflect.Descriptor instead.
func (*CosiCommitment) Descriptor() ([]byte, []int) {
	return file_messages_crypto_proto_rawDescGZIP(), []int{8}
}

func (x *CosiCommitment) GetCommitment() []byte {
	if x != nil {
		return x.Commitment
	}
	return nil
}

func (x *CosiCommitment) GetPubkey() []byte {
	if x != nil {
		return x.Pubkey
	}
	return nil
}

// *
// Request: Ask device to sign using CoSi
// @start
// @next CosiSignature
// @next Failure
type CosiSign struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddressN         []uint32 `protobuf:"varint,1,rep,name=address_n,json=addressN" json:"address_n,omitempty"`                        // BIP-32 path to derive the key from master node
	Data             []byte   `protobuf:"bytes,2,req,name=data" json:"data,omitempty"`                                                 // Data to be signed
	GlobalCommitment []byte   `protobuf:"bytes,3,req,name=global_commitment,json=globalCommitment" json:"global_commitment,omitempty"` // Aggregated commitment
	GlobalPubkey     []byte   `protobuf:"bytes,4,req,name=global_pubkey,json=globalPubkey" json:"global_pubkey,omitempty"`             // Aggregated public key
}

func (x *CosiSign) Reset() {
	*x = CosiSign{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_crypto_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CosiSign) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CosiSign) ProtoMessage() {}

func (x *CosiSign) ProtoReflect() protoreflect.Message {
	mi := &file_messages_crypto_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CosiSign.ProtoReflect.Descriptor instead.
func (*CosiSign) Descriptor() ([]byte, []int) {
	return file_messages_crypto_proto_rawDescGZIP(), []int{9}
}

func (x *CosiSign) GetAddressN() []uint32 {
	if x != nil {
		return x.AddressN
	}
	return nil
}

func (x *CosiSign) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CosiSign) GetGlobalCommitment() []byte {
	if x != nil {
		return x.GlobalCommitment
	}
	return nil
}

func (x *CosiSign) GetGlobalPubkey() []byte {
	if x != nil {
		return x.GlobalPubkey
	}
	return nil
}

// *
// Response: Contains a CoSi signature
// @end
type CosiSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature []byte `protobuf:"bytes,1,req,name=signature" json:"signature,omitempty"` // Signature
}

func (x *CosiSignature) Reset() {
	*x = CosiSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_crypto_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CosiSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CosiSignature) ProtoMessage() {}

func (x *CosiSignature) ProtoReflect() protoreflect.Message {
	mi := &file_messages_crypto_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CosiSignature.ProtoReflect.Descriptor instead.
func (*CosiSignature) Descriptor() ([]byte, []int) {
	return file_messages_crypto_proto_rawDescGZIP(), []int{10}
}

func (x *CosiSignature) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_messages_crypto_proto protoreflect.FileDescriptor

var file_messages_crypto_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2d, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x68, 0x77, 0x2e, 0x74, 0x72, 0x65, 0x7a,
	0x6f, 0x72, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x1a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xcb, 0x01, 0x0a, 0x0e, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x4b, 0x65, 0x79,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x5f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x4e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20,
	0x02, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x61, 0x73, 0x6b, 0x5f, 0x6f, 0x6e, 0x5f, 0x65,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x61, 0x73,
	0x6b, 0x4f, 0x6e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x61, 0x73,
	0x6b, 0x5f, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0c, 0x61, 0x73, 0x6b, 0x4f, 0x6e, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x76, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x76,
	0x22, 0x28, 0x0a, 0x10, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x8d, 0x01, 0x0a, 0x0c, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x17, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x3a, 0x01, 0x30, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0xd7, 0x01, 0x0a, 0x0c, 0x53,
	0x69, 0x67, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x43, 0x0a, 0x08, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x68, 0x77, 0x2e, 0x74, 0x72, 0x65, 0x7a, 0x6f, 0x72, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x2b, 0x0a, 0x10, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x68, 0x69,
	0x64, 0x64, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x3a, 0x00, 0x52, 0x0f, 0x63, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x12, 0x2b, 0x0a,
	0x10, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x76, 0x69, 0x73, 0x75, 0x61,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x3a, 0x00, 0x52, 0x0f, 0x63, 0x68, 0x61, 0x6c, 0x6c,
	0x65, 0x6e, 0x67, 0x65, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x63,
	0x64, 0x73, 0x61, 0x5f, 0x63, 0x75, 0x72, 0x76, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x63, 0x64, 0x73, 0x61, 0x43, 0x75, 0x72, 0x76, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x67, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x02, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x02,
	0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0xaa, 0x01,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x45, 0x43, 0x44, 0x48, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x4b, 0x65, 0x79, 0x12, 0x43, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x68, 0x77, 0x2e, 0x74, 0x72, 0x65, 0x7a, 0x6f,
	0x72, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x65, 0x65, 0x72,
	0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x02, 0x28,
	0x0c, 0x52, 0x0d, 0x70, 0x65, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x12, 0x28, 0x0a, 0x10, 0x65, 0x63, 0x64, 0x73, 0x61, 0x5f, 0x63, 0x75, 0x72, 0x76, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x63, 0x64, 0x73,
	0x61, 0x43, 0x75, 0x72, 0x76, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x50, 0x0a, 0x0e, 0x45, 0x43,
	0x44, 0x48, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x0c, 0x52, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x41, 0x0a, 0x0a,
	0x43, 0x6f, 0x73, 0x69, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x08, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4e, 0x12, 0x16, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x02, 0x18, 0x01, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x48, 0x0a, 0x0e, 0x43, 0x6f, 0x73, 0x69, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x02, 0x28,
	0x0c, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x22, 0x8d, 0x01, 0x0a, 0x08, 0x43, 0x6f,
	0x73, 0x69, 0x53, 0x69, 0x67, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x5f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x4e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x02, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2b, 0x0a, 0x11, 0x67, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x02,
	0x28, 0x0c, 0x52, 0x10, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x70,
	0x75, 0x62, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x0c, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x22, 0x2d, 0x0a, 0x0d, 0x43, 0x6f, 0x73,
	0x69, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x5b, 0x80, 0xa6, 0x1d, 0x01, 0x0a, 0x23,
	0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x74, 0x6f, 0x73, 0x68, 0x69, 0x6c, 0x61, 0x62, 0x73, 0x2e,
	0x74, 0x72, 0x65, 0x7a, 0x6f, 0x72, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x42, 0x13, 0x54, 0x72, 0x65, 0x7a, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5a, 0x1b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x74, 0x72, 0x65, 0x7a, 0x6f, 0x72, 0x2f, 0x74,
	0x72, 0x65, 0x7a, 0x6f, 0x72,
}

var (
	file_messages_crypto_proto_rawDescOnce sync.Once
	file_messages_crypto_proto_rawDescData = file_messages_crypto_proto_rawDesc
)

func file_messages_crypto_proto_rawDescGZIP() []byte {
	file_messages_crypto_proto_rawDescOnce.Do(func() {
		file_messages_crypto_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_crypto_proto_rawDescData)
	})
	return file_messages_crypto_proto_rawDescData
}

var file_messages_crypto_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_messages_crypto_proto_goTypes = []interface{}{
	(*CipherKeyValue)(nil),    // 0: hw.trezor.messages.crypto.CipherKeyValue
	(*CipheredKeyValue)(nil),  // 1: hw.trezor.messages.crypto.CipheredKeyValue
	(*IdentityType)(nil),      // 2: hw.trezor.messages.crypto.IdentityType
	(*SignIdentity)(nil),      // 3: hw.trezor.messages.crypto.SignIdentity
	(*SignedIdentity)(nil),    // 4: hw.trezor.messages.crypto.SignedIdentity
	(*GetECDHSessionKey)(nil), // 5: hw.trezor.messages.crypto.GetECDHSessionKey
	(*ECDHSessionKey)(nil),    // 6: hw.trezor.messages.crypto.ECDHSessionKey
	(*CosiCommit)(nil),        // 7: hw.trezor.messages.crypto.CosiCommit
	(*CosiCommitment)(nil),    // 8: hw.trezor.messages.crypto.CosiCommitment
	(*CosiSign)(nil),          // 9: hw.trezor.messages.crypto.CosiSign
	(*CosiSignature)(nil),     // 10: hw.trezor.messages.crypto.CosiSignature
}
var file_messages_crypto_proto_depIdxs = []int32{
	2, // 0: hw.trezor.messages.crypto.SignIdentity.identity:type_name -> hw.trezor.messages.crypto.IdentityType
	2, // 1: hw.trezor.messages.crypto.GetECDHSessionKey.identity:type_name -> hw.trezor.messages.crypto.IdentityType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_messages_crypto_proto_init() }
func file_messages_crypto_proto_init() {
	if File_messages_crypto_proto != nil {
		return
	}
	file_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_messages_crypto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CipherKeyValue); i {
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
		file_messages_crypto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CipheredKeyValue); i {
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
		file_messages_crypto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityType); i {
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
		file_messages_crypto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignIdentity); i {
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
		file_messages_crypto_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedIdentity); i {
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
		file_messages_crypto_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetECDHSessionKey); i {
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
		file_messages_crypto_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ECDHSessionKey); i {
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
		file_messages_crypto_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CosiCommit); i {
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
		file_messages_crypto_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CosiCommitment); i {
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
		file_messages_crypto_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CosiSign); i {
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
		file_messages_crypto_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CosiSignature); i {
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
			RawDescriptor: file_messages_crypto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_crypto_proto_goTypes,
		DependencyIndexes: file_messages_crypto_proto_depIdxs,
		MessageInfos:      file_messages_crypto_proto_msgTypes,
	}.Build()
	File_messages_crypto_proto = out.File
	file_messages_crypto_proto_rawDesc = nil
	file_messages_crypto_proto_goTypes = nil
	file_messages_crypto_proto_depIdxs = nil
}
