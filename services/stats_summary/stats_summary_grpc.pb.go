// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package stats_summary

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

// StatsSummaryServiceClient is the client API for StatsSummaryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatsSummaryServiceClient interface {
	GetStatsSummary(ctx context.Context, in *GetStatsSummaryRequest, opts ...grpc.CallOption) (*GetStatsSummaryResponse, error)
}

type statsSummaryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStatsSummaryServiceClient(cc grpc.ClientConnInterface) StatsSummaryServiceClient {
	return &statsSummaryServiceClient{cc}
}

func (c *statsSummaryServiceClient) GetStatsSummary(ctx context.Context, in *GetStatsSummaryRequest, opts ...grpc.CallOption) (*GetStatsSummaryResponse, error) {
	out := new(GetStatsSummaryResponse)
	err := c.cc.Invoke(ctx, "/stats_summary.StatsSummaryService/GetStatsSummary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatsSummaryServiceServer is the server API for StatsSummaryService service.
// All implementations must embed UnimplementedStatsSummaryServiceServer
// for forward compatibility
type StatsSummaryServiceServer interface {
	GetStatsSummary(context.Context, *GetStatsSummaryRequest) (*GetStatsSummaryResponse, error)
	mustEmbedUnimplementedStatsSummaryServiceServer()
}

// UnimplementedStatsSummaryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStatsSummaryServiceServer struct {
}

func (UnimplementedStatsSummaryServiceServer) GetStatsSummary(context.Context, *GetStatsSummaryRequest) (*GetStatsSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatsSummary not implemented")
}
func (UnimplementedStatsSummaryServiceServer) mustEmbedUnimplementedStatsSummaryServiceServer() {}

// UnsafeStatsSummaryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatsSummaryServiceServer will
// result in compilation errors.
type UnsafeStatsSummaryServiceServer interface {
	mustEmbedUnimplementedStatsSummaryServiceServer()
}

func RegisterStatsSummaryServiceServer(s grpc.ServiceRegistrar, srv StatsSummaryServiceServer) {
	s.RegisterService(&StatsSummaryService_ServiceDesc, srv)
}

func _StatsSummaryService_GetStatsSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatsSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsSummaryServiceServer).GetStatsSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stats_summary.StatsSummaryService/GetStatsSummary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsSummaryServiceServer).GetStatsSummary(ctx, req.(*GetStatsSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StatsSummaryService_ServiceDesc is the grpc.ServiceDesc for StatsSummaryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StatsSummaryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stats_summary.StatsSummaryService",
	HandlerType: (*StatsSummaryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatsSummary",
			Handler:    _StatsSummaryService_GetStatsSummary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/stats_summary/stats_summary.proto",
}
