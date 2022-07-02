// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alexa.proto

package Anki_Vector_external_interface

import (
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AlexaAuthState int32

const (
	// Invalid/error/versioning issue
	AlexaAuthState_ALEXA_AUTH_INVALID AlexaAuthState = 0
	// Not opted in, or opt-in attempted but failed
	AlexaAuthState_ALEXA_AUTH_UNINITIALIZED AlexaAuthState = 1
	// Opted in, and attempting to authorize
	AlexaAuthState_ALEXA_AUTH_REQUESTING_AUTH AlexaAuthState = 2
	// Opted in, and waiting on the user to enter a code
	AlexaAuthState_ALEXA_AUTH_WAITING_FOR_CODE AlexaAuthState = 3
	// Opted in, and authorized / in use
	AlexaAuthState_ALEXA_AUTH_AUTHORIZED AlexaAuthState = 4
)

var AlexaAuthState_name = map[int32]string{
	0: "ALEXA_AUTH_INVALID",
	1: "ALEXA_AUTH_UNINITIALIZED",
	2: "ALEXA_AUTH_REQUESTING_AUTH",
	3: "ALEXA_AUTH_WAITING_FOR_CODE",
	4: "ALEXA_AUTH_AUTHORIZED",
}
var AlexaAuthState_value = map[string]int32{
	"ALEXA_AUTH_INVALID":          0,
	"ALEXA_AUTH_UNINITIALIZED":    1,
	"ALEXA_AUTH_REQUESTING_AUTH":  2,
	"ALEXA_AUTH_WAITING_FOR_CODE": 3,
	"ALEXA_AUTH_AUTHORIZED":       4,
}

func (x AlexaAuthState) String() string {
	return proto.EnumName(AlexaAuthState_name, int32(x))
}
func (AlexaAuthState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_alexa_0c8502c9981f41e3, []int{0}
}

type AlexaAuthStateRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlexaAuthStateRequest) Reset()         { *m = AlexaAuthStateRequest{} }
func (m *AlexaAuthStateRequest) String() string { return proto.CompactTextString(m) }
func (*AlexaAuthStateRequest) ProtoMessage()    {}
func (*AlexaAuthStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_alexa_0c8502c9981f41e3, []int{0}
}
func (m *AlexaAuthStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlexaAuthStateRequest.Unmarshal(m, b)
}
func (m *AlexaAuthStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlexaAuthStateRequest.Marshal(b, m, deterministic)
}
func (dst *AlexaAuthStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlexaAuthStateRequest.Merge(dst, src)
}
func (m *AlexaAuthStateRequest) XXX_Size() int {
	return xxx_messageInfo_AlexaAuthStateRequest.Size(m)
}
func (m *AlexaAuthStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AlexaAuthStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AlexaAuthStateRequest proto.InternalMessageInfo

type AlexaAuthStateResponse struct {
	Status               *ResponseStatus `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	AuthState            AlexaAuthState  `protobuf:"varint,2,opt,name=auth_state,json=authState,enum=Anki.Vector.external_interface.AlexaAuthState" json:"auth_state,omitempty"`
	Extra                string          `protobuf:"bytes,3,opt,name=extra" json:"extra,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AlexaAuthStateResponse) Reset()         { *m = AlexaAuthStateResponse{} }
func (m *AlexaAuthStateResponse) String() string { return proto.CompactTextString(m) }
func (*AlexaAuthStateResponse) ProtoMessage()    {}
func (*AlexaAuthStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_alexa_0c8502c9981f41e3, []int{1}
}
func (m *AlexaAuthStateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlexaAuthStateResponse.Unmarshal(m, b)
}
func (m *AlexaAuthStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlexaAuthStateResponse.Marshal(b, m, deterministic)
}
func (dst *AlexaAuthStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlexaAuthStateResponse.Merge(dst, src)
}
func (m *AlexaAuthStateResponse) XXX_Size() int {
	return xxx_messageInfo_AlexaAuthStateResponse.Size(m)
}
func (m *AlexaAuthStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AlexaAuthStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AlexaAuthStateResponse proto.InternalMessageInfo

func (m *AlexaAuthStateResponse) GetStatus() *ResponseStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *AlexaAuthStateResponse) GetAuthState() AlexaAuthState {
	if m != nil {
		return m.AuthState
	}
	return AlexaAuthState_ALEXA_AUTH_INVALID
}

func (m *AlexaAuthStateResponse) GetExtra() string {
	if m != nil {
		return m.Extra
	}
	return ""
}

type AlexaOptInRequest struct {
	OptIn                bool     `protobuf:"varint,1,opt,name=opt_in,json=optIn" json:"opt_in,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlexaOptInRequest) Reset()         { *m = AlexaOptInRequest{} }
func (m *AlexaOptInRequest) String() string { return proto.CompactTextString(m) }
func (*AlexaOptInRequest) ProtoMessage()    {}
func (*AlexaOptInRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_alexa_0c8502c9981f41e3, []int{2}
}
func (m *AlexaOptInRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlexaOptInRequest.Unmarshal(m, b)
}
func (m *AlexaOptInRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlexaOptInRequest.Marshal(b, m, deterministic)
}
func (dst *AlexaOptInRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlexaOptInRequest.Merge(dst, src)
}
func (m *AlexaOptInRequest) XXX_Size() int {
	return xxx_messageInfo_AlexaOptInRequest.Size(m)
}
func (m *AlexaOptInRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AlexaOptInRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AlexaOptInRequest proto.InternalMessageInfo

func (m *AlexaOptInRequest) GetOptIn() bool {
	if m != nil {
		return m.OptIn
	}
	return false
}

type AlexaOptInResponse struct {
	Status               *ResponseStatus `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AlexaOptInResponse) Reset()         { *m = AlexaOptInResponse{} }
func (m *AlexaOptInResponse) String() string { return proto.CompactTextString(m) }
func (*AlexaOptInResponse) ProtoMessage()    {}
func (*AlexaOptInResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_alexa_0c8502c9981f41e3, []int{3}
}
func (m *AlexaOptInResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlexaOptInResponse.Unmarshal(m, b)
}
func (m *AlexaOptInResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlexaOptInResponse.Marshal(b, m, deterministic)
}
func (dst *AlexaOptInResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlexaOptInResponse.Merge(dst, src)
}
func (m *AlexaOptInResponse) XXX_Size() int {
	return xxx_messageInfo_AlexaOptInResponse.Size(m)
}
func (m *AlexaOptInResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AlexaOptInResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AlexaOptInResponse proto.InternalMessageInfo

func (m *AlexaOptInResponse) GetStatus() *ResponseStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type AlexaAuthEvent struct {
	AuthState            AlexaAuthState `protobuf:"varint,1,opt,name=auth_state,json=authState,enum=Anki.Vector.external_interface.AlexaAuthState" json:"auth_state,omitempty"`
	Extra                string         `protobuf:"bytes,2,opt,name=extra" json:"extra,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AlexaAuthEvent) Reset()         { *m = AlexaAuthEvent{} }
func (m *AlexaAuthEvent) String() string { return proto.CompactTextString(m) }
func (*AlexaAuthEvent) ProtoMessage()    {}
func (*AlexaAuthEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_alexa_0c8502c9981f41e3, []int{4}
}
func (m *AlexaAuthEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlexaAuthEvent.Unmarshal(m, b)
}
func (m *AlexaAuthEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlexaAuthEvent.Marshal(b, m, deterministic)
}
func (dst *AlexaAuthEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlexaAuthEvent.Merge(dst, src)
}
func (m *AlexaAuthEvent) XXX_Size() int {
	return xxx_messageInfo_AlexaAuthEvent.Size(m)
}
func (m *AlexaAuthEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_AlexaAuthEvent.DiscardUnknown(m)
}

var xxx_messageInfo_AlexaAuthEvent proto.InternalMessageInfo

func (m *AlexaAuthEvent) GetAuthState() AlexaAuthState {
	if m != nil {
		return m.AuthState
	}
	return AlexaAuthState_ALEXA_AUTH_INVALID
}

func (m *AlexaAuthEvent) GetExtra() string {
	if m != nil {
		return m.Extra
	}
	return ""
}

func init() {
	proto.RegisterType((*AlexaAuthStateRequest)(nil), "Anki.Vector.external_interface.AlexaAuthStateRequest")
	proto.RegisterType((*AlexaAuthStateResponse)(nil), "Anki.Vector.external_interface.AlexaAuthStateResponse")
	proto.RegisterType((*AlexaOptInRequest)(nil), "Anki.Vector.external_interface.AlexaOptInRequest")
	proto.RegisterType((*AlexaOptInResponse)(nil), "Anki.Vector.external_interface.AlexaOptInResponse")
	proto.RegisterType((*AlexaAuthEvent)(nil), "Anki.Vector.external_interface.AlexaAuthEvent")
	proto.RegisterEnum("Anki.Vector.external_interface.AlexaAuthState", AlexaAuthState_name, AlexaAuthState_value)
}

func init() { proto.RegisterFile("alexa.proto", fileDescriptor_alexa_0c8502c9981f41e3) }

var fileDescriptor_alexa_0c8502c9981f41e3 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x5d, 0x4f, 0xc2, 0x30,
	0x14, 0xb5, 0x20, 0x44, 0x2e, 0x09, 0x99, 0x8d, 0xc3, 0x89, 0x06, 0xc9, 0x9e, 0x08, 0x0f, 0x7b,
	0xd0, 0x5f, 0xd0, 0xc8, 0xd0, 0x26, 0xb8, 0xc5, 0xf1, 0xa1, 0x31, 0x26, 0x4b, 0x25, 0x35, 0x10,
	0xc9, 0x36, 0xb7, 0x3b, 0xc3, 0xef, 0xf1, 0xbf, 0xf8, 0xbf, 0x0c, 0x1d, 0x98, 0x86, 0x07, 0x7d,
	0xe1, 0xa5, 0xe9, 0xed, 0xb9, 0xe7, 0xdc, 0x7b, 0x4e, 0x0a, 0x75, 0xb1, 0x94, 0x2b, 0xe1, 0x24,
	0x69, 0x8c, 0x31, 0x6d, 0xb3, 0xe8, 0x7d, 0xe1, 0x4c, 0xe5, 0x0c, 0xe3, 0xd4, 0x91, 0x2b, 0x94,
	0x69, 0x24, 0x96, 0xe1, 0x22, 0x42, 0x99, 0xbe, 0x89, 0x99, 0x6c, 0x99, 0xa9, 0xcc, 0x92, 0x38,
	0xca, 0x64, 0x98, 0xa1, 0xc0, 0x3c, 0x2b, 0x68, 0xf6, 0x29, 0x98, 0x6c, 0xad, 0xc2, 0x72, 0x9c,
	0x8f, 0x50, 0xa0, 0x0c, 0xe4, 0x47, 0x2e, 0x33, 0xb4, 0xbf, 0x09, 0x34, 0x77, 0x91, 0x42, 0x80,
	0x0e, 0xa0, 0x5a, 0x68, 0x58, 0xa4, 0x43, 0xba, 0xf5, 0x2b, 0xc7, 0xf9, 0x7b, 0xb6, 0xb3, 0x65,
	0x8e, 0x14, 0x2b, 0xd8, 0xb0, 0xe9, 0x3d, 0x80, 0xc8, 0x71, 0xae, 0x16, 0x92, 0x56, 0xa9, 0x43,
	0xba, 0x8d, 0xff, 0xb5, 0x76, 0x76, 0xaa, 0x89, 0xed, 0x95, 0x9e, 0x40, 0x45, 0xae, 0x30, 0x15,
	0x56, 0xb9, 0x43, 0xba, 0xb5, 0xa0, 0x28, 0xec, 0x1e, 0x1c, 0x2b, 0x8a, 0x9f, 0x20, 0x8f, 0x36,
	0xe6, 0xa8, 0x09, 0xd5, 0x38, 0xc1, 0x70, 0x11, 0x29, 0x07, 0x47, 0x41, 0x25, 0x5e, 0xa3, 0xf6,
	0x0b, 0x50, 0xbd, 0x77, 0xbf, 0x76, 0xed, 0x1c, 0x1a, 0xbf, 0xcb, 0xbb, 0x9f, 0x32, 0xc2, 0x9d,
	0x00, 0xc8, 0xde, 0x02, 0x28, 0x69, 0x01, 0xf4, 0xbe, 0x88, 0x36, 0xb7, 0x68, 0x6c, 0x02, 0x65,
	0x43, 0xf7, 0x89, 0x85, 0x6c, 0x32, 0xbe, 0x0b, 0xb9, 0x37, 0x65, 0x43, 0xde, 0x37, 0x0e, 0xe8,
	0x05, 0x58, 0xda, 0xfb, 0xc4, 0xe3, 0x1e, 0x1f, 0x73, 0x36, 0xe4, 0xcf, 0x6e, 0xdf, 0x20, 0xb4,
	0x0d, 0x2d, 0x0d, 0x0d, 0xdc, 0x87, 0x89, 0x3b, 0x1a, 0x73, 0xef, 0x56, 0xd5, 0x46, 0x89, 0x5e,
	0xc2, 0xb9, 0x86, 0x3f, 0x32, 0xae, 0xc0, 0x81, 0x1f, 0x84, 0x37, 0x7e, 0xdf, 0x35, 0xca, 0xf4,
	0x0c, 0x4c, 0xad, 0x61, 0x7d, 0xf8, 0x81, 0xd2, 0x3e, 0x7c, 0xad, 0xaa, 0xdf, 0x78, 0xfd, 0x13,
	0x00, 0x00, 0xff, 0xff, 0xc8, 0x5a, 0x8c, 0x2a, 0xd3, 0x02, 0x00, 0x00,
}
