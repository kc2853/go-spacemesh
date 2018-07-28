// Code generated by protoc-gen-go. DO NOT EDIT.
// source: consensus/pb/messages.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ConsensusMessage struct {
	Msg                  *MessageData `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Validators           []*Validator `protobuf:"bytes,2,rep,name=validators,proto3" json:"validators,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ConsensusMessage) Reset()         { *m = ConsensusMessage{} }
func (m *ConsensusMessage) String() string { return proto.CompactTextString(m) }
func (*ConsensusMessage) ProtoMessage()    {}
func (*ConsensusMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_6586337b5fd43775, []int{0}
}
func (m *ConsensusMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsensusMessage.Unmarshal(m, b)
}
func (m *ConsensusMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsensusMessage.Marshal(b, m, deterministic)
}
func (dst *ConsensusMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusMessage.Merge(dst, src)
}
func (m *ConsensusMessage) XXX_Size() int {
	return xxx_messageInfo_ConsensusMessage.Size(m)
}
func (m *ConsensusMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusMessage proto.InternalMessageInfo

func (m *ConsensusMessage) GetMsg() *MessageData {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *ConsensusMessage) GetValidators() []*Validator {
	if m != nil {
		return m.Validators
	}
	return nil
}

type MessageData struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	AuthPubKey           []byte   `protobuf:"bytes,2,opt,name=authPubKey,proto3" json:"authPubKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageData) Reset()         { *m = MessageData{} }
func (m *MessageData) String() string { return proto.CompactTextString(m) }
func (*MessageData) ProtoMessage()    {}
func (*MessageData) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_6586337b5fd43775, []int{1}
}
func (m *MessageData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageData.Unmarshal(m, b)
}
func (m *MessageData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageData.Marshal(b, m, deterministic)
}
func (dst *MessageData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageData.Merge(dst, src)
}
func (m *MessageData) XXX_Size() int {
	return xxx_messageInfo_MessageData.Size(m)
}
func (m *MessageData) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageData.DiscardUnknown(m)
}

var xxx_messageInfo_MessageData proto.InternalMessageInfo

func (m *MessageData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *MessageData) GetAuthPubKey() []byte {
	if m != nil {
		return m.AuthPubKey
	}
	return nil
}

type Validator struct {
	AuthPubKey           []byte   `protobuf:"bytes,1,opt,name=authPubKey,proto3" json:"authPubKey,omitempty"`
	AuthorSign           []byte   `protobuf:"bytes,2,opt,name=authorSign,proto3" json:"authorSign,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Validator) Reset()         { *m = Validator{} }
func (m *Validator) String() string { return proto.CompactTextString(m) }
func (*Validator) ProtoMessage()    {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_6586337b5fd43775, []int{2}
}
func (m *Validator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Validator.Unmarshal(m, b)
}
func (m *Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
}
func (dst *Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(dst, src)
}
func (m *Validator) XXX_Size() int {
	return xxx_messageInfo_Validator.Size(m)
}
func (m *Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

func (m *Validator) GetAuthPubKey() []byte {
	if m != nil {
		return m.AuthPubKey
	}
	return nil
}

func (m *Validator) GetAuthorSign() []byte {
	if m != nil {
		return m.AuthorSign
	}
	return nil
}

func init() {
	proto.RegisterType((*ConsensusMessage)(nil), "pb.consensusMessage")
	proto.RegisterType((*MessageData)(nil), "pb.messageData")
	proto.RegisterType((*Validator)(nil), "pb.validator")
}

func init() {
	proto.RegisterFile("consensus/pb/messages.proto", fileDescriptor_messages_6586337b5fd43775)
}

var fileDescriptor_messages_6586337b5fd43775 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0xce, 0xcf, 0x2b,
	0x4e, 0xcd, 0x2b, 0x2e, 0x2d, 0xd6, 0x2f, 0x48, 0xd2, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x4a, 0xe1, 0x12, 0x80,
	0x2b, 0xf1, 0x85, 0x48, 0x0b, 0x29, 0x72, 0x31, 0xe7, 0x16, 0xa7, 0x4b, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x1b, 0xf1, 0xeb, 0x15, 0x24, 0xe9, 0x41, 0x35, 0xba, 0x24, 0x96, 0x24, 0x06, 0x81, 0xe4,
	0x84, 0x74, 0xb9, 0xb8, 0xca, 0x12, 0x73, 0x32, 0x53, 0x12, 0x4b, 0xf2, 0x8b, 0x8a, 0x25, 0x98,
	0x14, 0x98, 0x35, 0xb8, 0x8d, 0x78, 0x41, 0x2a, 0xe1, 0xa2, 0x41, 0x48, 0x0a, 0x94, 0x1c, 0xb9,
	0xb8, 0x91, 0x8c, 0x10, 0x12, 0xe2, 0x62, 0x49, 0x49, 0x2c, 0x49, 0x04, 0xdb, 0xc0, 0x13, 0x04,
	0x66, 0x0b, 0xc9, 0x71, 0x71, 0x25, 0x96, 0x96, 0x64, 0x04, 0x94, 0x26, 0x79, 0xa7, 0x56, 0x4a,
	0x30, 0x81, 0x65, 0x90, 0x44, 0x94, 0xbc, 0xb9, 0x38, 0xe1, 0x06, 0xa2, 0x29, 0x66, 0x44, 0x57,
	0x0c, 0x93, 0xcf, 0x2f, 0x0a, 0xce, 0x4c, 0xcf, 0x43, 0x36, 0x0c, 0x22, 0xe2, 0xc4, 0x12, 0xc5,
	0x54, 0x90, 0x94, 0xc4, 0x06, 0x0e, 0x06, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x27, 0x89,
	0xbd, 0x88, 0x25, 0x01, 0x00, 0x00,
}
