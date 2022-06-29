// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: proto/search.proto

package searchpb

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

// GetSearchClient is the client API for GetSearch service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetSearchClient interface {
	// Get search
	GetSearch(ctx context.Context, in *Search, opts ...grpc.CallOption) (*SearchResponse, error)
	// Post search
	PostSearch(ctx context.Context, in *Search, opts ...grpc.CallOption) (*SearchResponse, error)
}

type getSearchClient struct {
	cc grpc.ClientConnInterface
}

func NewGetSearchClient(cc grpc.ClientConnInterface) GetSearchClient {
	return &getSearchClient{cc}
}

func (c *getSearchClient) GetSearch(ctx context.Context, in *Search, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/search.GetSearch/GetSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getSearchClient) PostSearch(ctx context.Context, in *Search, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/search.GetSearch/PostSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetSearchServer is the server API for GetSearch service.
// All implementations must embed UnimplementedGetSearchServer
// for forward compatibility
type GetSearchServer interface {
	// Get search
	GetSearch(context.Context, *Search) (*SearchResponse, error)
	// Post search
	PostSearch(context.Context, *Search) (*SearchResponse, error)
	mustEmbedUnimplementedGetSearchServer()
}

// UnimplementedGetSearchServer must be embedded to have forward compatible implementations.
type UnimplementedGetSearchServer struct {
}

func (UnimplementedGetSearchServer) GetSearch(context.Context, *Search) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSearch not implemented")
}
func (UnimplementedGetSearchServer) PostSearch(context.Context, *Search) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostSearch not implemented")
}
func (UnimplementedGetSearchServer) mustEmbedUnimplementedGetSearchServer() {}

// UnsafeGetSearchServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetSearchServer will
// result in compilation errors.
type UnsafeGetSearchServer interface {
	mustEmbedUnimplementedGetSearchServer()
}

func RegisterGetSearchServer(s grpc.ServiceRegistrar, srv GetSearchServer) {
	s.RegisterService(&GetSearch_ServiceDesc, srv)
}

func _GetSearch_GetSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Search)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetSearchServer).GetSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search.GetSearch/GetSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetSearchServer).GetSearch(ctx, req.(*Search))
	}
	return interceptor(ctx, in, info, handler)
}

func _GetSearch_PostSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Search)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetSearchServer).PostSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search.GetSearch/PostSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetSearchServer).PostSearch(ctx, req.(*Search))
	}
	return interceptor(ctx, in, info, handler)
}

// GetSearch_ServiceDesc is the grpc.ServiceDesc for GetSearch service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetSearch_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "search.GetSearch",
	HandlerType: (*GetSearchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSearch",
			Handler:    _GetSearch_GetSearch_Handler,
		},
		{
			MethodName: "PostSearch",
			Handler:    _GetSearch_PostSearch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/search.proto",
}