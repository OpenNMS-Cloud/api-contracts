// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: ingester/ingester.proto

package v1

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

// IngesterClient is the client API for Ingester service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IngesterClient interface {
	StoreFlowDocument(ctx context.Context, in *StoreFlowDocumentRequest, opts ...grpc.CallOption) (*StoreFlowDocumentResponse, error)
	StoreFlowDocuments(ctx context.Context, in *StoreFlowDocumentsRequest, opts ...grpc.CallOption) (*StoreFlowDocumentsResponse, error)
}

type ingesterClient struct {
	cc grpc.ClientConnInterface
}

func NewIngesterClient(cc grpc.ClientConnInterface) IngesterClient {
	return &ingesterClient{cc}
}

func (c *ingesterClient) StoreFlowDocument(ctx context.Context, in *StoreFlowDocumentRequest, opts ...grpc.CallOption) (*StoreFlowDocumentResponse, error) {
	out := new(StoreFlowDocumentResponse)
	err := c.cc.Invoke(ctx, "/org.opennms.apicontracts.flows.ingester.v1.Ingester/StoreFlowDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ingesterClient) StoreFlowDocuments(ctx context.Context, in *StoreFlowDocumentsRequest, opts ...grpc.CallOption) (*StoreFlowDocumentsResponse, error) {
	out := new(StoreFlowDocumentsResponse)
	err := c.cc.Invoke(ctx, "/org.opennms.apicontracts.flows.ingester.v1.Ingester/StoreFlowDocuments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IngesterServer is the server API for Ingester service.
// All implementations must embed UnimplementedIngesterServer
// for forward compatibility
type IngesterServer interface {
	StoreFlowDocument(context.Context, *StoreFlowDocumentRequest) (*StoreFlowDocumentResponse, error)
	StoreFlowDocuments(context.Context, *StoreFlowDocumentsRequest) (*StoreFlowDocumentsResponse, error)
	mustEmbedUnimplementedIngesterServer()
}

// UnimplementedIngesterServer must be embedded to have forward compatible implementations.
type UnimplementedIngesterServer struct {
}

func (UnimplementedIngesterServer) StoreFlowDocument(context.Context, *StoreFlowDocumentRequest) (*StoreFlowDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreFlowDocument not implemented")
}
func (UnimplementedIngesterServer) StoreFlowDocuments(context.Context, *StoreFlowDocumentsRequest) (*StoreFlowDocumentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreFlowDocuments not implemented")
}
func (UnimplementedIngesterServer) mustEmbedUnimplementedIngesterServer() {}

// UnsafeIngesterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IngesterServer will
// result in compilation errors.
type UnsafeIngesterServer interface {
	mustEmbedUnimplementedIngesterServer()
}

func RegisterIngesterServer(s grpc.ServiceRegistrar, srv IngesterServer) {
	s.RegisterService(&Ingester_ServiceDesc, srv)
}

func _Ingester_StoreFlowDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreFlowDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngesterServer).StoreFlowDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.opennms.apicontracts.flows.ingester.v1.Ingester/StoreFlowDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngesterServer).StoreFlowDocument(ctx, req.(*StoreFlowDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ingester_StoreFlowDocuments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreFlowDocumentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngesterServer).StoreFlowDocuments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.opennms.apicontracts.flows.ingester.v1.Ingester/StoreFlowDocuments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngesterServer).StoreFlowDocuments(ctx, req.(*StoreFlowDocumentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Ingester_ServiceDesc is the grpc.ServiceDesc for Ingester service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ingester_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "org.opennms.apicontracts.flows.ingester.v1.Ingester",
	HandlerType: (*IngesterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreFlowDocument",
			Handler:    _Ingester_StoreFlowDocument_Handler,
		},
		{
			MethodName: "StoreFlowDocuments",
			Handler:    _Ingester_StoreFlowDocuments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ingester/ingester.proto",
}
