// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: book.proto

package bookpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	BookService_CreateBook_FullMethodName      = "/BookService/CreateBook"
	BookService_GetBooks_FullMethodName        = "/BookService/GetBooks"
	BookService_CreateManyBooks_FullMethodName = "/BookService/CreateManyBooks"
)

// BookServiceClient is the client API for BookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookServiceClient interface {
	CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*CreateBookResponse, error)
	GetBooks(ctx context.Context, in *GetBooksRequest, opts ...grpc.CallOption) (BookService_GetBooksClient, error)
	CreateManyBooks(ctx context.Context, opts ...grpc.CallOption) (BookService_CreateManyBooksClient, error)
}

type bookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookServiceClient(cc grpc.ClientConnInterface) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*CreateBookResponse, error) {
	out := new(CreateBookResponse)
	err := c.cc.Invoke(ctx, BookService_CreateBook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) GetBooks(ctx context.Context, in *GetBooksRequest, opts ...grpc.CallOption) (BookService_GetBooksClient, error) {
	stream, err := c.cc.NewStream(ctx, &BookService_ServiceDesc.Streams[0], BookService_GetBooks_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &bookServiceGetBooksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BookService_GetBooksClient interface {
	Recv() (*GetBooksResponse, error)
	grpc.ClientStream
}

type bookServiceGetBooksClient struct {
	grpc.ClientStream
}

func (x *bookServiceGetBooksClient) Recv() (*GetBooksResponse, error) {
	m := new(GetBooksResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bookServiceClient) CreateManyBooks(ctx context.Context, opts ...grpc.CallOption) (BookService_CreateManyBooksClient, error) {
	stream, err := c.cc.NewStream(ctx, &BookService_ServiceDesc.Streams[1], BookService_CreateManyBooks_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &bookServiceCreateManyBooksClient{stream}
	return x, nil
}

type BookService_CreateManyBooksClient interface {
	Send(*CreateManyBooksRequest) error
	CloseAndRecv() (*CreateManyBooksResponse, error)
	grpc.ClientStream
}

type bookServiceCreateManyBooksClient struct {
	grpc.ClientStream
}

func (x *bookServiceCreateManyBooksClient) Send(m *CreateManyBooksRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bookServiceCreateManyBooksClient) CloseAndRecv() (*CreateManyBooksResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CreateManyBooksResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BookServiceServer is the server API for BookService service.
// All implementations must embed UnimplementedBookServiceServer
// for forward compatibility
type BookServiceServer interface {
	CreateBook(context.Context, *CreateBookRequest) (*CreateBookResponse, error)
	GetBooks(*GetBooksRequest, BookService_GetBooksServer) error
	CreateManyBooks(BookService_CreateManyBooksServer) error
	mustEmbedUnimplementedBookServiceServer()
}

// UnimplementedBookServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookServiceServer struct {
}

func (UnimplementedBookServiceServer) CreateBook(context.Context, *CreateBookRequest) (*CreateBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}
func (UnimplementedBookServiceServer) GetBooks(*GetBooksRequest, BookService_GetBooksServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBooks not implemented")
}
func (UnimplementedBookServiceServer) CreateManyBooks(BookService_CreateManyBooksServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateManyBooks not implemented")
}
func (UnimplementedBookServiceServer) mustEmbedUnimplementedBookServiceServer() {}

// UnsafeBookServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookServiceServer will
// result in compilation errors.
type UnsafeBookServiceServer interface {
	mustEmbedUnimplementedBookServiceServer()
}

func RegisterBookServiceServer(s grpc.ServiceRegistrar, srv BookServiceServer) {
	s.RegisterService(&BookService_ServiceDesc, srv)
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
		FullMethod: BookService_CreateBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).CreateBook(ctx, req.(*CreateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_GetBooks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetBooksRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BookServiceServer).GetBooks(m, &bookServiceGetBooksServer{stream})
}

type BookService_GetBooksServer interface {
	Send(*GetBooksResponse) error
	grpc.ServerStream
}

type bookServiceGetBooksServer struct {
	grpc.ServerStream
}

func (x *bookServiceGetBooksServer) Send(m *GetBooksResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _BookService_CreateManyBooks_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BookServiceServer).CreateManyBooks(&bookServiceCreateManyBooksServer{stream})
}

type BookService_CreateManyBooksServer interface {
	SendAndClose(*CreateManyBooksResponse) error
	Recv() (*CreateManyBooksRequest, error)
	grpc.ServerStream
}

type bookServiceCreateManyBooksServer struct {
	grpc.ServerStream
}

func (x *bookServiceCreateManyBooksServer) SendAndClose(m *CreateManyBooksResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bookServiceCreateManyBooksServer) Recv() (*CreateManyBooksRequest, error) {
	m := new(CreateManyBooksRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BookService_ServiceDesc is the grpc.ServiceDesc for BookService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBook",
			Handler:    _BookService_CreateBook_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetBooks",
			Handler:       _BookService_GetBooks_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CreateManyBooks",
			Handler:       _BookService_CreateManyBooks_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "book.proto",
}