// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.29.3
// source: protobuf/avs.proto

package avsproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AggregatorClient is the client API for Aggregator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AggregatorClient interface {
	// Exchange for an Auth Key to authenticate in subsequent request
	GetKey(ctx context.Context, in *GetKeyReq, opts ...grpc.CallOption) (*KeyResp, error)
	// Smart Acccount Operation
	GetNonce(ctx context.Context, in *NonceRequest, opts ...grpc.CallOption) (*NonceResp, error)
	GetWallet(ctx context.Context, in *GetWalletReq, opts ...grpc.CallOption) (*GetWalletResp, error)
	ListWallets(ctx context.Context, in *ListWalletReq, opts ...grpc.CallOption) (*ListWalletResp, error)
	// Task Management Operation
	CreateTask(ctx context.Context, in *CreateTaskReq, opts ...grpc.CallOption) (*CreateTaskResp, error)
	ListTasks(ctx context.Context, in *ListTasksReq, opts ...grpc.CallOption) (*ListTasksResp, error)
	GetTask(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Task, error)
	ListExecutions(ctx context.Context, in *ListExecutionsReq, opts ...grpc.CallOption) (*ListExecutionsResp, error)
	GetExecution(ctx context.Context, in *ExecutionReq, opts ...grpc.CallOption) (*Execution, error)
	GetExecutionStatus(ctx context.Context, in *ExecutionReq, opts ...grpc.CallOption) (*ExecutionStatusResp, error)
	CancelTask(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error)
	DeleteTask(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error)
	TriggerTask(ctx context.Context, in *UserTriggerTaskReq, opts ...grpc.CallOption) (*UserTriggerTaskResp, error)
	// CreateSecret allow you to define a secret to be used in your tasks. The secret can be used with a special syntax of ${{secrets.name }}.
	// You can decide whether to grant secret to a single workflow or many workflow, or all of your workflow
	// By default, your secret is available across all of your tasks.
	CreateSecret(ctx context.Context, in *CreateOrUpdateSecretReq, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error)
	DeleteSecret(ctx context.Context, in *DeleteSecretReq, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error)
	// Return all secrets belong to this user. Currently we don't support any fine tune or filter yet.
	// Only secret names and config data are returned. The secret value aren't returned.
	ListSecrets(ctx context.Context, in *ListSecretsReq, opts ...grpc.CallOption) (*ListSecretsResp, error)
	// For simplicity, currently only the user who create the secrets can update its value, or update its permission.
	// The current implementation is also limited, update is an override action, not an appending action. So when updating, you need to pass the whole payload
	UpdateSecret(ctx context.Context, in *CreateOrUpdateSecretReq, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error)
}

type aggregatorClient struct {
	cc grpc.ClientConnInterface
}

func NewAggregatorClient(cc grpc.ClientConnInterface) AggregatorClient {
	return &aggregatorClient{cc}
}

func (c *aggregatorClient) GetKey(ctx context.Context, in *GetKeyReq, opts ...grpc.CallOption) (*KeyResp, error) {
	out := new(KeyResp)
	err := c.cc.Invoke(ctx, "/aggregator.Aggregator/GetKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorClient) GetNonce(ctx context.Context, in *NonceRequest, opts ...grpc.CallOption) (*NonceResp, error) {
	out := new(NonceResp)
	err := c.cc.Invoke(ctx, "/aggregator.Aggregator/GetNonce", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorClient) GetWallet(ctx context.Context, in *GetWalletReq, opts ...grpc.CallOption) (*GetWalletResp, error) {
	out := new(GetWalletResp)
	err := c.cc.Invoke(ctx, "/aggregator.Aggregator/GetWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorClient) ListWallets(ctx context.Context, in *ListWalletReq, opts ...grpc.CallOption) (*ListWalletResp, error) {
	out := new(ListWalletResp)
	err := c.cc.Invoke(ctx, "/aggregator.Aggregator/ListWallets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorClient) CreateTask(ctx context.Context, in *CreateTaskReq, opts ...grpc.CallOption) (*CreateTaskResp, error) {
	out := new(CreateTaskResp)
	err := c.cc.Invoke(ctx, "/aggregator.Aggregator/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorClient) ListTasks(ctx context.Context, in *ListTasksReq, opts ...grpc.CallOption) (*ListTasksResp, error) {
	out := new(ListTasksResp)
	err := c.cc.Invoke(ctx, "/aggregator.Aggregator/ListTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorClient) GetTask(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/aggregator.Aggregator/GetTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorClient) ListExecutions(ctx context.Context, in *ListExecutionsReq, opts ...grpc.CallOption) (*ListExecutionsResp, error) {
	out := new(ListExecutionsResp)
	err := c.cc.Invoke(ctx, "/aggregator.Aggregator/ListExecutions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorClient) GetExecution(ctx context.Context, in *ExecutionReq, opts ...grpc.CallOption) (*Execution, error) {
	out := new(Execution)
	err := c.cc.Invoke(ctx, "/aggregator.Aggregator/GetExecution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorClient) GetExecutionStatus(ctx context.Context, in *ExecutionReq, opts ...grpc.CallOption) (*ExecutionStatusResp, error) {
	out := new(ExecutionStatusResp)
	err := c.cc.Invoke(ctx, "/aggregator.Aggregator/GetExecutionStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorClient) CancelTask(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error) {
	out := new(wrapperspb.BoolValue)
	err := c.cc.Invoke(ctx, "/aggregator.Aggregator/CancelTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorClient) DeleteTask(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error) {
	out := new(wrapperspb.BoolValue)
	err := c.cc.Invoke(ctx, "/aggregator.Aggregator/DeleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorClient) TriggerTask(ctx context.Context, in *UserTriggerTaskReq, opts ...grpc.CallOption) (*UserTriggerTaskResp, error) {
	out := new(UserTriggerTaskResp)
	err := c.cc.Invoke(ctx, "/aggregator.Aggregator/TriggerTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorClient) CreateSecret(ctx context.Context, in *CreateOrUpdateSecretReq, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error) {
	out := new(wrapperspb.BoolValue)
	err := c.cc.Invoke(ctx, "/aggregator.Aggregator/CreateSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorClient) DeleteSecret(ctx context.Context, in *DeleteSecretReq, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error) {
	out := new(wrapperspb.BoolValue)
	err := c.cc.Invoke(ctx, "/aggregator.Aggregator/DeleteSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorClient) ListSecrets(ctx context.Context, in *ListSecretsReq, opts ...grpc.CallOption) (*ListSecretsResp, error) {
	out := new(ListSecretsResp)
	err := c.cc.Invoke(ctx, "/aggregator.Aggregator/ListSecrets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorClient) UpdateSecret(ctx context.Context, in *CreateOrUpdateSecretReq, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error) {
	out := new(wrapperspb.BoolValue)
	err := c.cc.Invoke(ctx, "/aggregator.Aggregator/UpdateSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AggregatorServer is the server API for Aggregator service.
// All implementations must embed UnimplementedAggregatorServer
// for forward compatibility
type AggregatorServer interface {
	// Exchange for an Auth Key to authenticate in subsequent request
	GetKey(context.Context, *GetKeyReq) (*KeyResp, error)
	// Smart Acccount Operation
	GetNonce(context.Context, *NonceRequest) (*NonceResp, error)
	GetWallet(context.Context, *GetWalletReq) (*GetWalletResp, error)
	ListWallets(context.Context, *ListWalletReq) (*ListWalletResp, error)
	// Task Management Operation
	CreateTask(context.Context, *CreateTaskReq) (*CreateTaskResp, error)
	ListTasks(context.Context, *ListTasksReq) (*ListTasksResp, error)
	GetTask(context.Context, *IdReq) (*Task, error)
	ListExecutions(context.Context, *ListExecutionsReq) (*ListExecutionsResp, error)
	GetExecution(context.Context, *ExecutionReq) (*Execution, error)
	GetExecutionStatus(context.Context, *ExecutionReq) (*ExecutionStatusResp, error)
	CancelTask(context.Context, *IdReq) (*wrapperspb.BoolValue, error)
	DeleteTask(context.Context, *IdReq) (*wrapperspb.BoolValue, error)
	TriggerTask(context.Context, *UserTriggerTaskReq) (*UserTriggerTaskResp, error)
	// CreateSecret allow you to define a secret to be used in your tasks. The secret can be used with a special syntax of ${{secrets.name }}.
	// You can decide whether to grant secret to a single workflow or many workflow, or all of your workflow
	// By default, your secret is available across all of your tasks.
	CreateSecret(context.Context, *CreateOrUpdateSecretReq) (*wrapperspb.BoolValue, error)
	DeleteSecret(context.Context, *DeleteSecretReq) (*wrapperspb.BoolValue, error)
	// Return all secrets belong to this user. Currently we don't support any fine tune or filter yet.
	// Only secret names and config data are returned. The secret value aren't returned.
	ListSecrets(context.Context, *ListSecretsReq) (*ListSecretsResp, error)
	// For simplicity, currently only the user who create the secrets can update its value, or update its permission.
	// The current implementation is also limited, update is an override action, not an appending action. So when updating, you need to pass the whole payload
	UpdateSecret(context.Context, *CreateOrUpdateSecretReq) (*wrapperspb.BoolValue, error)
	mustEmbedUnimplementedAggregatorServer()
}

// UnimplementedAggregatorServer must be embedded to have forward compatible implementations.
type UnimplementedAggregatorServer struct {
}

func (UnimplementedAggregatorServer) GetKey(context.Context, *GetKeyReq) (*KeyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKey not implemented")
}
func (UnimplementedAggregatorServer) GetNonce(context.Context, *NonceRequest) (*NonceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNonce not implemented")
}
func (UnimplementedAggregatorServer) GetWallet(context.Context, *GetWalletReq) (*GetWalletResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWallet not implemented")
}
func (UnimplementedAggregatorServer) ListWallets(context.Context, *ListWalletReq) (*ListWalletResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWallets not implemented")
}
func (UnimplementedAggregatorServer) CreateTask(context.Context, *CreateTaskReq) (*CreateTaskResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedAggregatorServer) ListTasks(context.Context, *ListTasksReq) (*ListTasksResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTasks not implemented")
}
func (UnimplementedAggregatorServer) GetTask(context.Context, *IdReq) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (UnimplementedAggregatorServer) ListExecutions(context.Context, *ListExecutionsReq) (*ListExecutionsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListExecutions not implemented")
}
func (UnimplementedAggregatorServer) GetExecution(context.Context, *ExecutionReq) (*Execution, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExecution not implemented")
}
func (UnimplementedAggregatorServer) GetExecutionStatus(context.Context, *ExecutionReq) (*ExecutionStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExecutionStatus not implemented")
}
func (UnimplementedAggregatorServer) CancelTask(context.Context, *IdReq) (*wrapperspb.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelTask not implemented")
}
func (UnimplementedAggregatorServer) DeleteTask(context.Context, *IdReq) (*wrapperspb.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}
func (UnimplementedAggregatorServer) TriggerTask(context.Context, *UserTriggerTaskReq) (*UserTriggerTaskResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TriggerTask not implemented")
}
func (UnimplementedAggregatorServer) CreateSecret(context.Context, *CreateOrUpdateSecretReq) (*wrapperspb.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSecret not implemented")
}
func (UnimplementedAggregatorServer) DeleteSecret(context.Context, *DeleteSecretReq) (*wrapperspb.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSecret not implemented")
}
func (UnimplementedAggregatorServer) ListSecrets(context.Context, *ListSecretsReq) (*ListSecretsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSecrets not implemented")
}
func (UnimplementedAggregatorServer) UpdateSecret(context.Context, *CreateOrUpdateSecretReq) (*wrapperspb.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSecret not implemented")
}
func (UnimplementedAggregatorServer) mustEmbedUnimplementedAggregatorServer() {}

// UnsafeAggregatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AggregatorServer will
// result in compilation errors.
type UnsafeAggregatorServer interface {
	mustEmbedUnimplementedAggregatorServer()
}

func RegisterAggregatorServer(s grpc.ServiceRegistrar, srv AggregatorServer) {
	s.RegisterService(&Aggregator_ServiceDesc, srv)
}

func _Aggregator_GetKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).GetKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aggregator.Aggregator/GetKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).GetKey(ctx, req.(*GetKeyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Aggregator_GetNonce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NonceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).GetNonce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aggregator.Aggregator/GetNonce",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).GetNonce(ctx, req.(*NonceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Aggregator_GetWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWalletReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).GetWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aggregator.Aggregator/GetWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).GetWallet(ctx, req.(*GetWalletReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Aggregator_ListWallets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWalletReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).ListWallets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aggregator.Aggregator/ListWallets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).ListWallets(ctx, req.(*ListWalletReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Aggregator_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aggregator.Aggregator/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).CreateTask(ctx, req.(*CreateTaskReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Aggregator_ListTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTasksReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).ListTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aggregator.Aggregator/ListTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).ListTasks(ctx, req.(*ListTasksReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Aggregator_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aggregator.Aggregator/GetTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).GetTask(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Aggregator_ListExecutions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListExecutionsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).ListExecutions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aggregator.Aggregator/ListExecutions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).ListExecutions(ctx, req.(*ListExecutionsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Aggregator_GetExecution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecutionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).GetExecution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aggregator.Aggregator/GetExecution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).GetExecution(ctx, req.(*ExecutionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Aggregator_GetExecutionStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecutionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).GetExecutionStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aggregator.Aggregator/GetExecutionStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).GetExecutionStatus(ctx, req.(*ExecutionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Aggregator_CancelTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).CancelTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aggregator.Aggregator/CancelTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).CancelTask(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Aggregator_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aggregator.Aggregator/DeleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).DeleteTask(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Aggregator_TriggerTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserTriggerTaskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).TriggerTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aggregator.Aggregator/TriggerTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).TriggerTask(ctx, req.(*UserTriggerTaskReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Aggregator_CreateSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrUpdateSecretReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).CreateSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aggregator.Aggregator/CreateSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).CreateSecret(ctx, req.(*CreateOrUpdateSecretReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Aggregator_DeleteSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSecretReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).DeleteSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aggregator.Aggregator/DeleteSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).DeleteSecret(ctx, req.(*DeleteSecretReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Aggregator_ListSecrets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSecretsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).ListSecrets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aggregator.Aggregator/ListSecrets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).ListSecrets(ctx, req.(*ListSecretsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Aggregator_UpdateSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrUpdateSecretReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).UpdateSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aggregator.Aggregator/UpdateSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).UpdateSecret(ctx, req.(*CreateOrUpdateSecretReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Aggregator_ServiceDesc is the grpc.ServiceDesc for Aggregator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Aggregator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aggregator.Aggregator",
	HandlerType: (*AggregatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKey",
			Handler:    _Aggregator_GetKey_Handler,
		},
		{
			MethodName: "GetNonce",
			Handler:    _Aggregator_GetNonce_Handler,
		},
		{
			MethodName: "GetWallet",
			Handler:    _Aggregator_GetWallet_Handler,
		},
		{
			MethodName: "ListWallets",
			Handler:    _Aggregator_ListWallets_Handler,
		},
		{
			MethodName: "CreateTask",
			Handler:    _Aggregator_CreateTask_Handler,
		},
		{
			MethodName: "ListTasks",
			Handler:    _Aggregator_ListTasks_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _Aggregator_GetTask_Handler,
		},
		{
			MethodName: "ListExecutions",
			Handler:    _Aggregator_ListExecutions_Handler,
		},
		{
			MethodName: "GetExecution",
			Handler:    _Aggregator_GetExecution_Handler,
		},
		{
			MethodName: "GetExecutionStatus",
			Handler:    _Aggregator_GetExecutionStatus_Handler,
		},
		{
			MethodName: "CancelTask",
			Handler:    _Aggregator_CancelTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _Aggregator_DeleteTask_Handler,
		},
		{
			MethodName: "TriggerTask",
			Handler:    _Aggregator_TriggerTask_Handler,
		},
		{
			MethodName: "CreateSecret",
			Handler:    _Aggregator_CreateSecret_Handler,
		},
		{
			MethodName: "DeleteSecret",
			Handler:    _Aggregator_DeleteSecret_Handler,
		},
		{
			MethodName: "ListSecrets",
			Handler:    _Aggregator_ListSecrets_Handler,
		},
		{
			MethodName: "UpdateSecret",
			Handler:    _Aggregator_UpdateSecret_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/avs.proto",
}
