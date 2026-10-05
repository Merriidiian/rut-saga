package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion9
const (
	ReservationService_CreateReservation_FullMethodName  = "/reservation.v1.ReservationService/CreateReservation"
	ReservationService_GetReservation_FullMethodName     = "/reservation.v1.ReservationService/GetReservation"
	ReservationService_ConfirmReservation_FullMethodName = "/reservation.v1.ReservationService/ConfirmReservation"
	ReservationService_CancelReservation_FullMethodName  = "/reservation.v1.ReservationService/CancelReservation"
)

type ReservationServiceClient interface {
	CreateReservation(ctx context.Context, in *CreateReservationRequest, opts ...grpc.CallOption) (*Reservation, error)
	GetReservation(ctx context.Context, in *ReservationRequest, opts ...grpc.CallOption) (*Reservation, error)
	ConfirmReservation(ctx context.Context, in *ReservationRequest, opts ...grpc.CallOption) (*Reservation, error)
	CancelReservation(ctx context.Context, in *ReservationRequest, opts ...grpc.CallOption) (*Reservation, error)
}
type reservationServiceClient struct{ cc grpc.ClientConnInterface }

func NewReservationServiceClient(cc grpc.ClientConnInterface) ReservationServiceClient {
	return &reservationServiceClient{cc}
}
func (c *reservationServiceClient) CreateReservation(ctx context.Context, in *CreateReservationRequest, opts ...grpc.CallOption) (*Reservation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Reservation)
	err := c.cc.Invoke(ctx, ReservationService_CreateReservation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *reservationServiceClient) GetReservation(ctx context.Context, in *ReservationRequest, opts ...grpc.CallOption) (*Reservation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Reservation)
	err := c.cc.Invoke(ctx, ReservationService_GetReservation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *reservationServiceClient) ConfirmReservation(ctx context.Context, in *ReservationRequest, opts ...grpc.CallOption) (*Reservation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Reservation)
	err := c.cc.Invoke(ctx, ReservationService_ConfirmReservation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *reservationServiceClient) CancelReservation(ctx context.Context, in *ReservationRequest, opts ...grpc.CallOption) (*Reservation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Reservation)
	err := c.cc.Invoke(ctx, ReservationService_CancelReservation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type ReservationServiceServer interface {
	CreateReservation(context.Context, *CreateReservationRequest) (*Reservation, error)
	GetReservation(context.Context, *ReservationRequest) (*Reservation, error)
	ConfirmReservation(context.Context, *ReservationRequest) (*Reservation, error)
	CancelReservation(context.Context, *ReservationRequest) (*Reservation, error)
	mustEmbedUnimplementedReservationServiceServer()
}
type UnimplementedReservationServiceServer struct{}

func (UnimplementedReservationServiceServer) CreateReservation(context.Context, *CreateReservationRequest) (*Reservation, error) {
	return nil, status.Error(codes.Unimplemented, "method CreateReservation not implemented")
}
func (UnimplementedReservationServiceServer) GetReservation(context.Context, *ReservationRequest) (*Reservation, error) {
	return nil, status.Error(codes.Unimplemented, "method GetReservation not implemented")
}
func (UnimplementedReservationServiceServer) ConfirmReservation(context.Context, *ReservationRequest) (*Reservation, error) {
	return nil, status.Error(codes.Unimplemented, "method ConfirmReservation not implemented")
}
func (UnimplementedReservationServiceServer) CancelReservation(context.Context, *ReservationRequest) (*Reservation, error) {
	return nil, status.Error(codes.Unimplemented, "method CancelReservation not implemented")
}
func (UnimplementedReservationServiceServer) mustEmbedUnimplementedReservationServiceServer() {
}
func (UnimplementedReservationServiceServer) testEmbeddedByValue() {
}

type UnsafeReservationServiceServer interface{ mustEmbedUnimplementedReservationServiceServer() }

func RegisterReservationServiceServer(s grpc.ServiceRegistrar, srv ReservationServiceServer) {
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ReservationService_ServiceDesc, srv)
}
func _ReservationService_CreateReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).CreateReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: ReservationService_CreateReservation_FullMethodName}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).CreateReservation(ctx, req.(*CreateReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _ReservationService_GetReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: ReservationService_GetReservation_FullMethodName}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetReservation(ctx, req.(*ReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _ReservationService_ConfirmReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).ConfirmReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: ReservationService_ConfirmReservation_FullMethodName}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).ConfirmReservation(ctx, req.(*ReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _ReservationService_CancelReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).CancelReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: ReservationService_CancelReservation_FullMethodName}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).CancelReservation(ctx, req.(*ReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var ReservationService_ServiceDesc = grpc.ServiceDesc{ServiceName: "reservation.v1.ReservationService", HandlerType: (*ReservationServiceServer)(nil), Methods: []grpc.MethodDesc{{MethodName: "CreateReservation", Handler: _ReservationService_CreateReservation_Handler}, {MethodName: "GetReservation", Handler: _ReservationService_GetReservation_Handler}, {MethodName: "ConfirmReservation", Handler: _ReservationService_ConfirmReservation_Handler}, {MethodName: "CancelReservation", Handler: _ReservationService_CancelReservation_Handler}}, Streams: []grpc.StreamDesc{}, Metadata: "reservation.proto"}
