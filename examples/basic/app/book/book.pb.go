// Code generated by protoc-gen-go. DO NOT EDIT.
// source: book/book.proto

package book

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	author "github.com/ysugimoto/grpc-graphql-gateway/examples/basic/app/author"
	_ "github.com/ysugimoto/grpc-graphql-gateway/graphql"
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

type BookType int32

const (
	BookType_JAVASCRIPT  BookType = 0
	BookType_ECMASCRIPT  BookType = 1
	BookType_GIT         BookType = 2
	BookType_ASP_DOT_NET BookType = 3
)

var BookType_name = map[int32]string{
	0: "JAVASCRIPT",
	1: "ECMASCRIPT",
	2: "GIT",
	3: "ASP_DOT_NET",
}

var BookType_value = map[string]int32{
	"JAVASCRIPT":  0,
	"ECMASCRIPT":  1,
	"GIT":         2,
	"ASP_DOT_NET": 3,
}

func (x BookType) String() string {
	return proto.EnumName(BookType_name, int32(x))
}

func (BookType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ee9082fb44230b1b, []int{0}
}

type Book struct {
	Id                   int64             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string            `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Type                 BookType          `protobuf:"varint,3,opt,name=type,proto3,enum=book.BookType" json:"type,omitempty"`
	Author               *author.Author    `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	Description          *Book_Description `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee9082fb44230b1b, []int{0}
}

func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (m *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(m, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Book) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Book) GetType() BookType {
	if m != nil {
		return m.Type
	}
	return BookType_JAVASCRIPT
}

func (m *Book) GetAuthor() *author.Author {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *Book) GetDescription() *Book_Description {
	if m != nil {
		return m.Description
	}
	return nil
}

type Book_Description struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Book_Description) Reset()         { *m = Book_Description{} }
func (m *Book_Description) String() string { return proto.CompactTextString(m) }
func (*Book_Description) ProtoMessage()    {}
func (*Book_Description) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee9082fb44230b1b, []int{0, 0}
}

func (m *Book_Description) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book_Description.Unmarshal(m, b)
}
func (m *Book_Description) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book_Description.Marshal(b, m, deterministic)
}
func (m *Book_Description) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book_Description.Merge(m, src)
}
func (m *Book_Description) XXX_Size() int {
	return xxx_messageInfo_Book_Description.Size(m)
}
func (m *Book_Description) XXX_DiscardUnknown() {
	xxx_messageInfo_Book_Description.DiscardUnknown(m)
}

var xxx_messageInfo_Book_Description proto.InternalMessageInfo

func (m *Book_Description) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ListBooksRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBooksRequest) Reset()         { *m = ListBooksRequest{} }
func (m *ListBooksRequest) String() string { return proto.CompactTextString(m) }
func (*ListBooksRequest) ProtoMessage()    {}
func (*ListBooksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee9082fb44230b1b, []int{1}
}

func (m *ListBooksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBooksRequest.Unmarshal(m, b)
}
func (m *ListBooksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBooksRequest.Marshal(b, m, deterministic)
}
func (m *ListBooksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBooksRequest.Merge(m, src)
}
func (m *ListBooksRequest) XXX_Size() int {
	return xxx_messageInfo_ListBooksRequest.Size(m)
}
func (m *ListBooksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBooksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListBooksRequest proto.InternalMessageInfo

type ListBooksResponse struct {
	Books                []*Book  `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBooksResponse) Reset()         { *m = ListBooksResponse{} }
func (m *ListBooksResponse) String() string { return proto.CompactTextString(m) }
func (*ListBooksResponse) ProtoMessage()    {}
func (*ListBooksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee9082fb44230b1b, []int{2}
}

func (m *ListBooksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBooksResponse.Unmarshal(m, b)
}
func (m *ListBooksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBooksResponse.Marshal(b, m, deterministic)
}
func (m *ListBooksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBooksResponse.Merge(m, src)
}
func (m *ListBooksResponse) XXX_Size() int {
	return xxx_messageInfo_ListBooksResponse.Size(m)
}
func (m *ListBooksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBooksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListBooksResponse proto.InternalMessageInfo

func (m *ListBooksResponse) GetBooks() []*Book {
	if m != nil {
		return m.Books
	}
	return nil
}

type GetBookRequest struct {
	// this is example comment for id field
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBookRequest) Reset()         { *m = GetBookRequest{} }
func (m *GetBookRequest) String() string { return proto.CompactTextString(m) }
func (*GetBookRequest) ProtoMessage()    {}
func (*GetBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee9082fb44230b1b, []int{3}
}

func (m *GetBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBookRequest.Unmarshal(m, b)
}
func (m *GetBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBookRequest.Marshal(b, m, deterministic)
}
func (m *GetBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBookRequest.Merge(m, src)
}
func (m *GetBookRequest) XXX_Size() int {
	return xxx_messageInfo_GetBookRequest.Size(m)
}
func (m *GetBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBookRequest proto.InternalMessageInfo

func (m *GetBookRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CreateBookRequest struct {
	Title                string                       `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Type                 BookType                     `protobuf:"varint,2,opt,name=type,proto3,enum=book.BookType" json:"type,omitempty"`
	Author               *author.Author               `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Nested               *CreateBookRequest_NestedOne `protobuf:"bytes,4,opt,name=nested,proto3" json:"nested,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *CreateBookRequest) Reset()         { *m = CreateBookRequest{} }
func (m *CreateBookRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBookRequest) ProtoMessage()    {}
func (*CreateBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee9082fb44230b1b, []int{4}
}

func (m *CreateBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBookRequest.Unmarshal(m, b)
}
func (m *CreateBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBookRequest.Marshal(b, m, deterministic)
}
func (m *CreateBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBookRequest.Merge(m, src)
}
func (m *CreateBookRequest) XXX_Size() int {
	return xxx_messageInfo_CreateBookRequest.Size(m)
}
func (m *CreateBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBookRequest proto.InternalMessageInfo

func (m *CreateBookRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreateBookRequest) GetType() BookType {
	if m != nil {
		return m.Type
	}
	return BookType_JAVASCRIPT
}

func (m *CreateBookRequest) GetAuthor() *author.Author {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *CreateBookRequest) GetNested() *CreateBookRequest_NestedOne {
	if m != nil {
		return m.Nested
	}
	return nil
}

type CreateBookRequest_NestedOne struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBookRequest_NestedOne) Reset()         { *m = CreateBookRequest_NestedOne{} }
func (m *CreateBookRequest_NestedOne) String() string { return proto.CompactTextString(m) }
func (*CreateBookRequest_NestedOne) ProtoMessage()    {}
func (*CreateBookRequest_NestedOne) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee9082fb44230b1b, []int{4, 0}
}

func (m *CreateBookRequest_NestedOne) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBookRequest_NestedOne.Unmarshal(m, b)
}
func (m *CreateBookRequest_NestedOne) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBookRequest_NestedOne.Marshal(b, m, deterministic)
}
func (m *CreateBookRequest_NestedOne) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBookRequest_NestedOne.Merge(m, src)
}
func (m *CreateBookRequest_NestedOne) XXX_Size() int {
	return xxx_messageInfo_CreateBookRequest_NestedOne.Size(m)
}
func (m *CreateBookRequest_NestedOne) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBookRequest_NestedOne.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBookRequest_NestedOne proto.InternalMessageInfo

func (m *CreateBookRequest_NestedOne) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterEnum("book.BookType", BookType_name, BookType_value)
	proto.RegisterType((*Book)(nil), "book.Book")
	proto.RegisterType((*Book_Description)(nil), "book.Book.Description")
	proto.RegisterType((*ListBooksRequest)(nil), "book.ListBooksRequest")
	proto.RegisterType((*ListBooksResponse)(nil), "book.ListBooksResponse")
	proto.RegisterType((*GetBookRequest)(nil), "book.GetBookRequest")
	proto.RegisterType((*CreateBookRequest)(nil), "book.CreateBookRequest")
	proto.RegisterType((*CreateBookRequest_NestedOne)(nil), "book.CreateBookRequest.NestedOne")
}

func init() { proto.RegisterFile("book/book.proto", fileDescriptor_ee9082fb44230b1b) }

var fileDescriptor_ee9082fb44230b1b = []byte{
	// 577 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x4e, 0xdb, 0x40,
	0x10, 0xee, 0xda, 0x49, 0x20, 0x63, 0xc9, 0x98, 0x2d, 0x3f, 0x5b, 0xf7, 0xe2, 0xba, 0x97, 0x08,
	0x09, 0x1b, 0x52, 0x55, 0xa5, 0xbd, 0x85, 0x80, 0x10, 0x15, 0x05, 0x64, 0x22, 0x0e, 0xbd, 0x20,
	0xc7, 0x59, 0x25, 0x16, 0x4e, 0x76, 0xf1, 0xae, 0xd3, 0xe6, 0xde, 0x53, 0x8e, 0x7d, 0x84, 0xbe,
	0x43, 0x1e, 0xa6, 0xe7, 0xbe, 0x48, 0xe5, 0xb5, 0x49, 0x02, 0xb4, 0x12, 0x97, 0x4c, 0xe6, 0xe7,
	0x9b, 0x19, 0x7f, 0x9f, 0xc7, 0xb0, 0xd6, 0x65, 0xec, 0xd6, 0xcf, 0x7f, 0x3c, 0x9e, 0x32, 0xc9,
	0x70, 0x25, 0xff, 0x6f, 0xbf, 0x0c, 0x33, 0x39, 0x60, 0xa9, 0x5f, 0x98, 0x22, 0x65, 0x6f, 0xf6,
	0xd3, 0x90, 0x0f, 0xee, 0x12, 0xbf, 0xb4, 0x45, 0xd8, 0xfd, 0x83, 0xa0, 0x72, 0xc8, 0xd8, 0x2d,
	0x36, 0x41, 0x8b, 0x7b, 0x04, 0x39, 0xa8, 0xa1, 0x07, 0x5a, 0xdc, 0xc3, 0x1b, 0x50, 0x95, 0xb1,
	0x4c, 0x28, 0xd1, 0x1c, 0xd4, 0xa8, 0x07, 0x85, 0x83, 0x5d, 0xa8, 0xc8, 0x09, 0xa7, 0x44, 0x77,
	0x50, 0xc3, 0x6c, 0x9a, 0x9e, 0x9a, 0x9d, 0xe3, 0x3b, 0x13, 0x4e, 0x03, 0x95, 0xc3, 0x4d, 0xa8,
	0x15, 0x93, 0x49, 0xc5, 0x41, 0x0d, 0xa3, 0x69, 0x7b, 0x82, 0x0d, 0xa9, 0x17, 0x72, 0xee, 0x8d,
	0xf7, 0xbd, 0x72, 0xa9, 0x96, 0x32, 0x41, 0x59, 0x89, 0x0f, 0xc0, 0xe8, 0x51, 0x11, 0xa5, 0x31,
	0x97, 0x31, 0x1b, 0x91, 0xaa, 0x02, 0x6e, 0x2d, 0xda, 0x7b, 0x47, 0x8b, 0x6c, 0xb0, 0x5c, 0x6a,
	0xbf, 0x05, 0x63, 0x29, 0x97, 0xaf, 0x3d, 0x0e, 0x93, 0x8c, 0xaa, 0x27, 0xa9, 0x07, 0x85, 0xe3,
	0x62, 0xb0, 0xce, 0x62, 0x21, 0xf3, 0x4e, 0x22, 0xa0, 0x77, 0x19, 0x15, 0xd2, 0x7d, 0x0f, 0xeb,
	0x4b, 0x31, 0xc1, 0xd9, 0x48, 0x50, 0xec, 0x40, 0x35, 0x9f, 0x29, 0x08, 0x72, 0xf4, 0x86, 0xd1,
	0x84, 0xc5, 0x06, 0x41, 0x91, 0x70, 0x77, 0xc0, 0x3c, 0xa1, 0x0a, 0x55, 0x36, 0xc2, 0x64, 0xc1,
	0xdc, 0xe1, 0xea, 0xaf, 0x19, 0xa9, 0xd8, 0xda, 0xfe, 0x5e, 0xce, 0xa1, 0xfb, 0x1b, 0xc1, 0x7a,
	0x3b, 0xa5, 0xa1, 0xa4, 0xcb, 0xf5, 0x73, 0x66, 0xd1, 0xbf, 0x98, 0xd5, 0x9e, 0xc5, 0xac, 0xfe,
	0x6c, 0x66, 0x3f, 0x42, 0x6d, 0x44, 0x85, 0xa4, 0xbd, 0x52, 0x8d, 0x37, 0x45, 0xe7, 0x27, 0x6b,
	0x79, 0xe7, 0xaa, 0xea, 0x62, 0x44, 0x83, 0x12, 0x60, 0xbf, 0x86, 0xfa, 0x3c, 0xf8, 0xf8, 0xfd,
	0xd8, 0x39, 0x82, 0xd5, 0xfb, 0xed, 0xb0, 0x09, 0xf0, 0xb9, 0x75, 0xdd, 0xba, 0x6a, 0x07, 0xa7,
	0x97, 0x1d, 0xeb, 0x45, 0xee, 0x1f, 0xb7, 0xbf, 0xdc, 0xfb, 0x08, 0xaf, 0x80, 0x7e, 0x72, 0xda,
	0xb1, 0x34, 0xbc, 0x06, 0x46, 0xeb, 0xea, 0xf2, 0xe6, 0xe8, 0xa2, 0x73, 0x73, 0x7e, 0xdc, 0xb1,
	0xf4, 0xe6, 0x0f, 0x0d, 0x8c, 0xbc, 0xcd, 0x15, 0x4d, 0xc7, 0x71, 0x44, 0xf1, 0x35, 0xd4, 0xe7,
	0xa2, 0xe0, 0x52, 0xff, 0xc7, 0xca, 0xd9, 0xdb, 0x4f, 0xe2, 0x85, 0x7a, 0x2e, 0x99, 0xce, 0xc8,
	0x46, 0xa9, 0xa0, 0x6b, 0xac, 0x22, 0x0b, 0xd9, 0x85, 0x83, 0x3f, 0xc0, 0x4a, 0xa9, 0x1a, 0xde,
	0x28, 0xd0, 0x0f, 0x45, 0xb4, 0x97, 0x94, 0x76, 0x61, 0x3a, 0x23, 0x35, 0x50, 0xb7, 0x84, 0xcf,
	0x00, 0x16, 0x54, 0xe1, 0xed, 0xff, 0x90, 0xf7, 0x00, 0xfe, 0xea, 0xe7, 0x8c, 0x6c, 0x02, 0x44,
	0xaa, 0x48, 0x1d, 0xe4, 0x0a, 0x54, 0xe3, 0x11, 0xcf, 0xa4, 0xbd, 0x35, 0x9d, 0x11, 0x0c, 0x66,
	0xc2, 0xa2, 0x30, 0x19, 0x30, 0x21, 0x3f, 0x1d, 0xec, 0x1d, 0xec, 0x59, 0xe8, 0xb0, 0xfd, 0xb5,
	0xd5, 0x8f, 0xe5, 0x20, 0xeb, 0x7a, 0x11, 0x1b, 0xfa, 0x13, 0x91, 0xf5, 0xe3, 0x21, 0x93, 0xcc,
	0xef, 0xa7, 0x3c, 0xda, 0x2d, 0x0f, 0x76, 0xb7, 0x1f, 0x4a, 0xfa, 0x2d, 0x9c, 0xf8, 0xf4, 0x7b,
	0x38, 0xe4, 0x09, 0x15, 0x7e, 0x37, 0x14, 0x71, 0xe4, 0x87, 0x9c, 0xab, 0x4f, 0x40, 0xb7, 0xa6,
	0x2e, 0xfa, 0xdd, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x51, 0x94, 0xd4, 0x27, 0x16, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BookServiceClient is the client API for BookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BookServiceClient interface {
	ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (*ListBooksResponse, error)
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error)
	CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*Book, error)
}

type bookServiceClient struct {
	cc *grpc.ClientConn
}

func NewBookServiceClient(cc *grpc.ClientConn) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (*ListBooksResponse, error) {
	out := new(ListBooksResponse)
	err := c.cc.Invoke(ctx, "/book.BookService/ListBooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/book.BookService/GetBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/book.BookService/CreateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServiceServer is the server API for BookService service.
type BookServiceServer interface {
	ListBooks(context.Context, *ListBooksRequest) (*ListBooksResponse, error)
	GetBook(context.Context, *GetBookRequest) (*Book, error)
	CreateBook(context.Context, *CreateBookRequest) (*Book, error)
}

// UnimplementedBookServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBookServiceServer struct {
}

func (*UnimplementedBookServiceServer) ListBooks(ctx context.Context, req *ListBooksRequest) (*ListBooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBooks not implemented")
}
func (*UnimplementedBookServiceServer) GetBook(ctx context.Context, req *GetBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (*UnimplementedBookServiceServer) CreateBook(ctx context.Context, req *CreateBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}

func RegisterBookServiceServer(s *grpc.Server, srv BookServiceServer) {
	s.RegisterService(&_BookService_serviceDesc, srv)
}

func _BookService_ListBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).ListBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookService/ListBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).ListBooks(ctx, req.(*ListBooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookService/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetBook(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookService/CreateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).CreateBook(ctx, req.(*CreateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BookService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "book.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBooks",
			Handler:    _BookService_ListBooks_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _BookService_GetBook_Handler,
		},
		{
			MethodName: "CreateBook",
			Handler:    _BookService_CreateBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "book/book.proto",
}