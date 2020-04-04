// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fileps.proto

package chunker

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type FileMeta struct {
	FilePath             string   `protobuf:"bytes,1,opt,name=filePath,proto3" json:"filePath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileMeta) Reset()         { *m = FileMeta{} }
func (m *FileMeta) String() string { return proto.CompactTextString(m) }
func (*FileMeta) ProtoMessage()    {}
func (*FileMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3dc9f0412321c81, []int{0}
}

func (m *FileMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileMeta.Unmarshal(m, b)
}
func (m *FileMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileMeta.Marshal(b, m, deterministic)
}
func (m *FileMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileMeta.Merge(m, src)
}
func (m *FileMeta) XXX_Size() int {
	return xxx_messageInfo_FileMeta.Size(m)
}
func (m *FileMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_FileMeta.DiscardUnknown(m)
}

var xxx_messageInfo_FileMeta proto.InternalMessageInfo

func (m *FileMeta) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

type Chunk struct {
	Chunk                []byte   `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chunk) Reset()         { *m = Chunk{} }
func (m *Chunk) String() string { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()    {}
func (*Chunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3dc9f0412321c81, []int{1}
}

func (m *Chunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chunk.Unmarshal(m, b)
}
func (m *Chunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chunk.Marshal(b, m, deterministic)
}
func (m *Chunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chunk.Merge(m, src)
}
func (m *Chunk) XXX_Size() int {
	return xxx_messageInfo_Chunk.Size(m)
}
func (m *Chunk) XXX_DiscardUnknown() {
	xxx_messageInfo_Chunk.DiscardUnknown(m)
}

var xxx_messageInfo_Chunk proto.InternalMessageInfo

func (m *Chunk) GetChunk() []byte {
	if m != nil {
		return m.Chunk
	}
	return nil
}

func init() {
	proto.RegisterType((*FileMeta)(nil), "chunker.FileMeta")
	proto.RegisterType((*Chunk)(nil), "chunker.Chunk")
}

func init() {
	proto.RegisterFile("fileps.proto", fileDescriptor_e3dc9f0412321c81)
}

var fileDescriptor_e3dc9f0412321c81 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xcb, 0xcc, 0x49,
	0x2d, 0x28, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0xce, 0x28, 0xcd, 0xcb, 0x4e,
	0x2d, 0x52, 0x52, 0xe3, 0xe2, 0x70, 0xcb, 0xcc, 0x49, 0xf5, 0x4d, 0x2d, 0x49, 0x14, 0x92, 0xe2,
	0xe2, 0x00, 0x29, 0x0a, 0x48, 0x2c, 0xc9, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf3,
	0x95, 0x64, 0xb9, 0x58, 0x9d, 0x41, 0x5a, 0x84, 0x44, 0xb8, 0x58, 0xc1, 0x7a, 0xc1, 0x2a, 0x78,
	0x82, 0x20, 0x1c, 0xa3, 0x52, 0x2e, 0x6e, 0x90, 0x31, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9,
	0x42, 0xc6, 0x5c, 0x5c, 0x2e, 0xf9, 0xe5, 0x79, 0x39, 0xf9, 0x89, 0x29, 0xa9, 0x45, 0x42, 0x82,
	0x7a, 0x50, 0xdb, 0xf4, 0x60, 0x56, 0x49, 0xf1, 0xc1, 0x85, 0xc0, 0xa6, 0x2a, 0x31, 0x18, 0x30,
	0x0a, 0x19, 0x72, 0x71, 0x84, 0x16, 0x40, 0xb5, 0xa0, 0xc9, 0x4b, 0x61, 0x1a, 0xa1, 0xc4, 0xa0,
	0xc1, 0x98, 0xc4, 0x06, 0xf6, 0x8d, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xdf, 0xbd, 0x0e, 0xec,
	0xdd, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileServiceClient interface {
	Downloader(ctx context.Context, in *FileMeta, opts ...grpc.CallOption) (FileService_DownloaderClient, error)
	Uploader(ctx context.Context, opts ...grpc.CallOption) (FileService_UploaderClient, error)
}

type fileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServiceClient(cc grpc.ClientConnInterface) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) Downloader(ctx context.Context, in *FileMeta, opts ...grpc.CallOption) (FileService_DownloaderClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FileService_serviceDesc.Streams[0], "/chunker.FileService/Downloader", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileServiceDownloaderClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileService_DownloaderClient interface {
	Recv() (*Chunk, error)
	grpc.ClientStream
}

type fileServiceDownloaderClient struct {
	grpc.ClientStream
}

func (x *fileServiceDownloaderClient) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileServiceClient) Uploader(ctx context.Context, opts ...grpc.CallOption) (FileService_UploaderClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FileService_serviceDesc.Streams[1], "/chunker.FileService/Uploader", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileServiceUploaderClient{stream}
	return x, nil
}

type FileService_UploaderClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*FileMeta, error)
	grpc.ClientStream
}

type fileServiceUploaderClient struct {
	grpc.ClientStream
}

func (x *fileServiceUploaderClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileServiceUploaderClient) CloseAndRecv() (*FileMeta, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(FileMeta)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileServiceServer is the server API for FileService service.
type FileServiceServer interface {
	Downloader(*FileMeta, FileService_DownloaderServer) error
	Uploader(FileService_UploaderServer) error
}

// UnimplementedFileServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFileServiceServer struct {
}

func (*UnimplementedFileServiceServer) Downloader(req *FileMeta, srv FileService_DownloaderServer) error {
	return status.Errorf(codes.Unimplemented, "method Downloader not implemented")
}
func (*UnimplementedFileServiceServer) Uploader(srv FileService_UploaderServer) error {
	return status.Errorf(codes.Unimplemented, "method Uploader not implemented")
}

func RegisterFileServiceServer(s *grpc.Server, srv FileServiceServer) {
	s.RegisterService(&_FileService_serviceDesc, srv)
}

func _FileService_Downloader_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FileMeta)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileServiceServer).Downloader(m, &fileServiceDownloaderServer{stream})
}

type FileService_DownloaderServer interface {
	Send(*Chunk) error
	grpc.ServerStream
}

type fileServiceDownloaderServer struct {
	grpc.ServerStream
}

func (x *fileServiceDownloaderServer) Send(m *Chunk) error {
	return x.ServerStream.SendMsg(m)
}

func _FileService_Uploader_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileServiceServer).Uploader(&fileServiceUploaderServer{stream})
}

type FileService_UploaderServer interface {
	SendAndClose(*FileMeta) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type fileServiceUploaderServer struct {
	grpc.ServerStream
}

func (x *fileServiceUploaderServer) SendAndClose(m *FileMeta) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileServiceUploaderServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _FileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chunker.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Downloader",
			Handler:       _FileService_Downloader_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Uploader",
			Handler:       _FileService_Uploader_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "fileps.proto",
}