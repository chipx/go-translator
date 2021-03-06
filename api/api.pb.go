// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SimpleRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleRequest) Reset()         { *m = SimpleRequest{} }
func (m *SimpleRequest) String() string { return proto.CompactTextString(m) }
func (*SimpleRequest) ProtoMessage()    {}
func (*SimpleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *SimpleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleRequest.Unmarshal(m, b)
}
func (m *SimpleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleRequest.Marshal(b, m, deterministic)
}
func (m *SimpleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleRequest.Merge(m, src)
}
func (m *SimpleRequest) XXX_Size() int {
	return xxx_messageInfo_SimpleRequest.Size(m)
}
func (m *SimpleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleRequest proto.InternalMessageInfo

type SimpleResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleResponse) Reset()         { *m = SimpleResponse{} }
func (m *SimpleResponse) String() string { return proto.CompactTextString(m) }
func (*SimpleResponse) ProtoMessage()    {}
func (*SimpleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *SimpleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleResponse.Unmarshal(m, b)
}
func (m *SimpleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleResponse.Marshal(b, m, deterministic)
}
func (m *SimpleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleResponse.Merge(m, src)
}
func (m *SimpleResponse) XXX_Size() int {
	return xxx_messageInfo_SimpleResponse.Size(m)
}
func (m *SimpleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleResponse proto.InternalMessageInfo

func (m *SimpleResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *SimpleResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type TranslateRequest struct {
	Lang                 string   `protobuf:"bytes,1,opt,name=lang,proto3" json:"lang,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TranslateRequest) Reset()         { *m = TranslateRequest{} }
func (m *TranslateRequest) String() string { return proto.CompactTextString(m) }
func (*TranslateRequest) ProtoMessage()    {}
func (*TranslateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *TranslateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranslateRequest.Unmarshal(m, b)
}
func (m *TranslateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranslateRequest.Marshal(b, m, deterministic)
}
func (m *TranslateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranslateRequest.Merge(m, src)
}
func (m *TranslateRequest) XXX_Size() int {
	return xxx_messageInfo_TranslateRequest.Size(m)
}
func (m *TranslateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TranslateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TranslateRequest proto.InternalMessageInfo

func (m *TranslateRequest) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *TranslateRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type SetTranslateRequest struct {
	Lang                 string   `protobuf:"bytes,1,opt,name=lang,proto3" json:"lang,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetTranslateRequest) Reset()         { *m = SetTranslateRequest{} }
func (m *SetTranslateRequest) String() string { return proto.CompactTextString(m) }
func (*SetTranslateRequest) ProtoMessage()    {}
func (*SetTranslateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *SetTranslateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetTranslateRequest.Unmarshal(m, b)
}
func (m *SetTranslateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetTranslateRequest.Marshal(b, m, deterministic)
}
func (m *SetTranslateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetTranslateRequest.Merge(m, src)
}
func (m *SetTranslateRequest) XXX_Size() int {
	return xxx_messageInfo_SetTranslateRequest.Size(m)
}
func (m *SetTranslateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetTranslateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetTranslateRequest proto.InternalMessageInfo

func (m *SetTranslateRequest) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *SetTranslateRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SetTranslateRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GetAllRequest struct {
	KeyPrefix            string               `protobuf:"bytes,1,opt,name=keyPrefix,proto3" json:"keyPrefix,omitempty"`
	Since                *timestamp.Timestamp `protobuf:"bytes,2,opt,name=since,proto3" json:"since,omitempty"`
	Langs                []string             `protobuf:"bytes,3,rep,name=langs,proto3" json:"langs,omitempty"`
	Trunslated           int32                `protobuf:"varint,4,opt,name=trunslated,proto3" json:"trunslated,omitempty"`
	Limit                int32                `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	Page                 int32                `protobuf:"varint,6,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetAllRequest) Reset()         { *m = GetAllRequest{} }
func (m *GetAllRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllRequest) ProtoMessage()    {}
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *GetAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllRequest.Unmarshal(m, b)
}
func (m *GetAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllRequest.Marshal(b, m, deterministic)
}
func (m *GetAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllRequest.Merge(m, src)
}
func (m *GetAllRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllRequest.Size(m)
}
func (m *GetAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllRequest proto.InternalMessageInfo

func (m *GetAllRequest) GetKeyPrefix() string {
	if m != nil {
		return m.KeyPrefix
	}
	return ""
}

func (m *GetAllRequest) GetSince() *timestamp.Timestamp {
	if m != nil {
		return m.Since
	}
	return nil
}

func (m *GetAllRequest) GetLangs() []string {
	if m != nil {
		return m.Langs
	}
	return nil
}

func (m *GetAllRequest) GetTrunslated() int32 {
	if m != nil {
		return m.Trunslated
	}
	return 0
}

func (m *GetAllRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetAllRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

type GetAllResponse struct {
	List                 []*Vocabulary `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetAllResponse) Reset()         { *m = GetAllResponse{} }
func (m *GetAllResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllResponse) ProtoMessage()    {}
func (*GetAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{5}
}

func (m *GetAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllResponse.Unmarshal(m, b)
}
func (m *GetAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllResponse.Marshal(b, m, deterministic)
}
func (m *GetAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllResponse.Merge(m, src)
}
func (m *GetAllResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllResponse.Size(m)
}
func (m *GetAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllResponse proto.InternalMessageInfo

func (m *GetAllResponse) GetList() []*Vocabulary {
	if m != nil {
		return m.List
	}
	return nil
}

type Translate struct {
	Lang                 string   `protobuf:"bytes,1,opt,name=lang,proto3" json:"lang,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Translate) Reset()         { *m = Translate{} }
func (m *Translate) String() string { return proto.CompactTextString(m) }
func (*Translate) ProtoMessage()    {}
func (*Translate) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{6}
}

func (m *Translate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Translate.Unmarshal(m, b)
}
func (m *Translate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Translate.Marshal(b, m, deterministic)
}
func (m *Translate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Translate.Merge(m, src)
}
func (m *Translate) XXX_Size() int {
	return xxx_messageInfo_Translate.Size(m)
}
func (m *Translate) XXX_DiscardUnknown() {
	xxx_messageInfo_Translate.DiscardUnknown(m)
}

var xxx_messageInfo_Translate proto.InternalMessageInfo

func (m *Translate) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *Translate) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Translate) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Vocabulary struct {
	Lang                 string       `protobuf:"bytes,1,opt,name=lang,proto3" json:"lang,omitempty"`
	Data                 []*Translate `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Vocabulary) Reset()         { *m = Vocabulary{} }
func (m *Vocabulary) String() string { return proto.CompactTextString(m) }
func (*Vocabulary) ProtoMessage()    {}
func (*Vocabulary) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{7}
}

func (m *Vocabulary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vocabulary.Unmarshal(m, b)
}
func (m *Vocabulary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vocabulary.Marshal(b, m, deterministic)
}
func (m *Vocabulary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vocabulary.Merge(m, src)
}
func (m *Vocabulary) XXX_Size() int {
	return xxx_messageInfo_Vocabulary.Size(m)
}
func (m *Vocabulary) XXX_DiscardUnknown() {
	xxx_messageInfo_Vocabulary.DiscardUnknown(m)
}

var xxx_messageInfo_Vocabulary proto.InternalMessageInfo

func (m *Vocabulary) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *Vocabulary) GetData() []*Translate {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*SimpleRequest)(nil), "SimpleRequest")
	proto.RegisterType((*SimpleResponse)(nil), "SimpleResponse")
	proto.RegisterType((*TranslateRequest)(nil), "TranslateRequest")
	proto.RegisterType((*SetTranslateRequest)(nil), "SetTranslateRequest")
	proto.RegisterType((*GetAllRequest)(nil), "GetAllRequest")
	proto.RegisterType((*GetAllResponse)(nil), "GetAllResponse")
	proto.RegisterType((*Translate)(nil), "Translate")
	proto.RegisterType((*Vocabulary)(nil), "Vocabulary")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0xe3, 0x24, 0xe0, 0x89, 0x9a, 0x94, 0xa5, 0x07, 0xcb, 0x42, 0xad, 0xe5, 0x93, 0x2b,
	0xa1, 0x2d, 0x84, 0x4b, 0x2f, 0x48, 0x54, 0x42, 0xca, 0x01, 0x2a, 0x21, 0xa7, 0xe5, 0xbe, 0x89,
	0x27, 0xd1, 0x2a, 0xb6, 0xd7, 0x78, 0xc7, 0x12, 0xf9, 0x09, 0xfc, 0x2a, 0xfe, 0x1a, 0xf2, 0x6e,
	0x96, 0x34, 0xa1, 0x42, 0x88, 0xde, 0x76, 0xde, 0x7c, 0xec, 0x7b, 0x6f, 0x06, 0x02, 0x51, 0x4b,
	0x5e, 0x37, 0x8a, 0x54, 0x74, 0xb1, 0x56, 0x6a, 0x5d, 0xe0, 0x95, 0x89, 0x16, 0xed, 0xea, 0x8a,
	0x64, 0x89, 0x9a, 0x44, 0x59, 0xdb, 0x82, 0x64, 0x02, 0x27, 0x73, 0x59, 0xd6, 0x05, 0x66, 0xf8,
	0xad, 0x45, 0x4d, 0xc9, 0x47, 0x18, 0x3b, 0x40, 0xd7, 0xaa, 0xd2, 0xc8, 0x42, 0x78, 0xa6, 0xdb,
	0xe5, 0x12, 0xb5, 0x0e, 0xbd, 0xd8, 0x4b, 0x9f, 0x67, 0x2e, 0xec, 0x32, 0x25, 0x6a, 0x2d, 0xd6,
	0x18, 0xf6, 0x62, 0x2f, 0x0d, 0x32, 0x17, 0x26, 0xd7, 0x70, 0x7a, 0xd7, 0x88, 0x4a, 0x17, 0x82,
	0xdc, 0x64, 0xc6, 0xa0, 0x5f, 0x88, 0x6a, 0x6d, 0x86, 0x04, 0x99, 0x79, 0xb3, 0x53, 0xf0, 0x37,
	0xb8, 0xdd, 0x75, 0x77, 0xcf, 0xe4, 0x1e, 0x5e, 0xce, 0x91, 0xfe, 0xaf, 0xf9, 0x21, 0x21, 0xff,
	0x90, 0xd0, 0x4f, 0x0f, 0x4e, 0x66, 0x48, 0x37, 0x45, 0xe1, 0x26, 0xbe, 0x82, 0x60, 0x83, 0xdb,
	0x2f, 0x0d, 0xae, 0xe4, 0xf7, 0xdd, 0xd8, 0x3d, 0xc0, 0xde, 0xc0, 0x40, 0xcb, 0x6a, 0x69, 0x85,
	0x8d, 0xa6, 0x11, 0xb7, 0x46, 0x72, 0x67, 0x24, 0xbf, 0x73, 0x46, 0x66, 0xb6, 0x90, 0x9d, 0xc1,
	0xa0, 0x63, 0xa5, 0x43, 0x3f, 0xf6, 0xd3, 0x20, 0xb3, 0x01, 0x3b, 0x07, 0xa0, 0xa6, 0xb5, 0x5a,
	0xf2, 0xb0, 0x1f, 0x7b, 0xe9, 0x20, 0x7b, 0x80, 0x98, 0x2e, 0x59, 0x4a, 0x0a, 0x07, 0x26, 0x65,
	0x83, 0x4e, 0x6d, 0xdd, 0x89, 0x18, 0x1a, 0xd0, 0xbc, 0x93, 0xb7, 0x30, 0x76, 0x02, 0x76, 0x8b,
	0xb9, 0x80, 0x7e, 0x21, 0x35, 0x85, 0x5e, 0xec, 0xa7, 0xa3, 0xe9, 0x88, 0x7f, 0x55, 0x4b, 0xb1,
	0x68, 0x0b, 0xd1, 0x6c, 0x33, 0x93, 0x48, 0x3e, 0x41, 0xf0, 0xdb, 0xc8, 0x27, 0x3b, 0xf8, 0x01,
	0x60, 0xff, 0xc1, 0xa3, 0xd3, 0xce, 0xa1, 0x9f, 0x0b, 0x12, 0x61, 0xcf, 0xf0, 0x01, 0xbe, 0x5f,
	0xa2, 0xc1, 0xa7, 0x3f, 0x7a, 0x00, 0x0e, 0x53, 0x0d, 0xbb, 0x04, 0x7f, 0x86, 0xc4, 0x5e, 0xf0,
	0xe3, 0x65, 0x47, 0x13, 0x7e, 0x74, 0x82, 0x97, 0x30, 0xb4, 0xda, 0xd9, 0x98, 0x1f, 0x6c, 0x31,
	0x9a, 0xf0, 0x23, 0x53, 0xde, 0xc3, 0x64, 0x86, 0xf4, 0x59, 0x68, 0xba, 0x55, 0xb9, 0x5c, 0x49,
	0xcc, 0xd9, 0x98, 0x1f, 0x9c, 0x78, 0xf4, 0x97, 0x65, 0xb2, 0x6b, 0x60, 0xb7, 0xa2, 0xd9, 0xdc,
	0xe8, 0xfb, 0x8a, 0x1c, 0xad, 0xfc, 0x9f, 0x38, 0xbe, 0x06, 0x7f, 0x8e, 0xc4, 0xce, 0xf8, 0x23,
	0xe7, 0xfb, 0x47, 0xf5, 0x62, 0x68, 0xfe, 0x7e, 0xf7, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xed, 0x8f,
	0x74, 0xf1, 0xac, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TranslatorClient is the client API for Translator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TranslatorClient interface {
	Get(ctx context.Context, in *TranslateRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	GetLastModified(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*timestamp.Timestamp, error)
	MarkAsUntranslated(ctx context.Context, in *TranslateRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
	Set(ctx context.Context, in *SetTranslateRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
}

type translatorClient struct {
	cc *grpc.ClientConn
}

func NewTranslatorClient(cc *grpc.ClientConn) TranslatorClient {
	return &translatorClient{cc}
}

func (c *translatorClient) Get(ctx context.Context, in *TranslateRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := c.cc.Invoke(ctx, "/Translator/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translatorClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/Translator/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translatorClient) GetLastModified(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*timestamp.Timestamp, error) {
	out := new(timestamp.Timestamp)
	err := c.cc.Invoke(ctx, "/Translator/GetLastModified", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translatorClient) MarkAsUntranslated(ctx context.Context, in *TranslateRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := c.cc.Invoke(ctx, "/Translator/MarkAsUntranslated", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translatorClient) Set(ctx context.Context, in *SetTranslateRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := c.cc.Invoke(ctx, "/Translator/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TranslatorServer is the server API for Translator service.
type TranslatorServer interface {
	Get(context.Context, *TranslateRequest) (*SimpleResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	GetLastModified(context.Context, *SimpleRequest) (*timestamp.Timestamp, error)
	MarkAsUntranslated(context.Context, *TranslateRequest) (*SimpleResponse, error)
	Set(context.Context, *SetTranslateRequest) (*SimpleResponse, error)
}

// UnimplementedTranslatorServer can be embedded to have forward compatible implementations.
type UnimplementedTranslatorServer struct {
}

func (*UnimplementedTranslatorServer) Get(ctx context.Context, req *TranslateRequest) (*SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedTranslatorServer) GetAll(ctx context.Context, req *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedTranslatorServer) GetLastModified(ctx context.Context, req *SimpleRequest) (*timestamp.Timestamp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastModified not implemented")
}
func (*UnimplementedTranslatorServer) MarkAsUntranslated(ctx context.Context, req *TranslateRequest) (*SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkAsUntranslated not implemented")
}
func (*UnimplementedTranslatorServer) Set(ctx context.Context, req *SetTranslateRequest) (*SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}

func RegisterTranslatorServer(s *grpc.Server, srv TranslatorServer) {
	s.RegisterService(&_Translator_serviceDesc, srv)
}

func _Translator_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranslateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslatorServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Translator/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslatorServer).Get(ctx, req.(*TranslateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Translator_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslatorServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Translator/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslatorServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Translator_GetLastModified_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslatorServer).GetLastModified(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Translator/GetLastModified",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslatorServer).GetLastModified(ctx, req.(*SimpleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Translator_MarkAsUntranslated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranslateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslatorServer).MarkAsUntranslated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Translator/MarkAsUntranslated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslatorServer).MarkAsUntranslated(ctx, req.(*TranslateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Translator_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTranslateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslatorServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Translator/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslatorServer).Set(ctx, req.(*SetTranslateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Translator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Translator",
	HandlerType: (*TranslatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Translator_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _Translator_GetAll_Handler,
		},
		{
			MethodName: "GetLastModified",
			Handler:    _Translator_GetLastModified_Handler,
		},
		{
			MethodName: "MarkAsUntranslated",
			Handler:    _Translator_MarkAsUntranslated_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _Translator_Set_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
