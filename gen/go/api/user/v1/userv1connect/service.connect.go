// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/user/v1/service.proto

package userv1connect

import (
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"

	connect "connectrpc.com/connect"
	v1 "github.com/grpc-connectgo-api-demo/wallet/gen/go/api/user/v1"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// UserServiceName is the fully-qualified name of the UserService service.
	UserServiceName = "api.user.v1.UserService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// UserServiceRegisterUserProcedure is the fully-qualified name of the UserService's RegisterUser
	// RPC.
	UserServiceRegisterUserProcedure = "/api.user.v1.UserService/RegisterUser"
	// UserServiceLoginUserProcedure is the fully-qualified name of the UserService's LoginUser RPC.
	UserServiceLoginUserProcedure = "/api.user.v1.UserService/LoginUser"
	// UserServiceGetUserAccountProcedure is the fully-qualified name of the UserService's
	// GetUserAccount RPC.
	UserServiceGetUserAccountProcedure = "/api.user.v1.UserService/GetUserAccount"
)

// UserServiceClient is a client for the api.user.v1.UserService service.
type UserServiceClient interface {
	RegisterUser(context.Context, *connect.Request[v1.RegisterUserRequest]) (*connect.Response[v1.RegisterUserResponse], error)
	LoginUser(context.Context, *connect.Request[v1.LoginUserRequest]) (*connect.Response[v1.LoginUserResponse], error)
	GetUserAccount(context.Context, *connect.Request[v1.GetUserAccountRequest]) (*connect.Response[v1.GetUserAccountResponse], error)
}

// NewUserServiceClient constructs a client for the api.user.v1.UserService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) UserServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userServiceClient{
		registerUser: connect.NewClient[v1.RegisterUserRequest, v1.RegisterUserResponse](
			httpClient,
			baseURL+UserServiceRegisterUserProcedure,
			opts...,
		),
		loginUser: connect.NewClient[v1.LoginUserRequest, v1.LoginUserResponse](
			httpClient,
			baseURL+UserServiceLoginUserProcedure,
			opts...,
		),
		getUserAccount: connect.NewClient[v1.GetUserAccountRequest, v1.GetUserAccountResponse](
			httpClient,
			baseURL+UserServiceGetUserAccountProcedure,
			opts...,
		),
	}
}

// userServiceClient implements UserServiceClient.
type userServiceClient struct {
	registerUser   *connect.Client[v1.RegisterUserRequest, v1.RegisterUserResponse]
	loginUser      *connect.Client[v1.LoginUserRequest, v1.LoginUserResponse]
	getUserAccount *connect.Client[v1.GetUserAccountRequest, v1.GetUserAccountResponse]
}

// RegisterUser calls api.user.v1.UserService.RegisterUser.
func (c *userServiceClient) RegisterUser(ctx context.Context, req *connect.Request[v1.RegisterUserRequest]) (*connect.Response[v1.RegisterUserResponse], error) {
	return c.registerUser.CallUnary(ctx, req)
}

// LoginUser calls api.user.v1.UserService.LoginUser.
func (c *userServiceClient) LoginUser(ctx context.Context, req *connect.Request[v1.LoginUserRequest]) (*connect.Response[v1.LoginUserResponse], error) {
	return c.loginUser.CallUnary(ctx, req)
}

// GetUserAccount calls api.user.v1.UserService.GetUserAccount.
func (c *userServiceClient) GetUserAccount(ctx context.Context, req *connect.Request[v1.GetUserAccountRequest]) (*connect.Response[v1.GetUserAccountResponse], error) {
	return c.getUserAccount.CallUnary(ctx, req)
}

// UserServiceHandler is an implementation of the api.user.v1.UserService service.
type UserServiceHandler interface {
	RegisterUser(context.Context, *connect.Request[v1.RegisterUserRequest]) (*connect.Response[v1.RegisterUserResponse], error)
	LoginUser(context.Context, *connect.Request[v1.LoginUserRequest]) (*connect.Response[v1.LoginUserResponse], error)
	GetUserAccount(context.Context, *connect.Request[v1.GetUserAccountRequest]) (*connect.Response[v1.GetUserAccountResponse], error)
}

// NewUserServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserServiceHandler(svc UserServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	userServiceRegisterUserHandler := connect.NewUnaryHandler(
		UserServiceRegisterUserProcedure,
		svc.RegisterUser,
		opts...,
	)
	userServiceLoginUserHandler := connect.NewUnaryHandler(
		UserServiceLoginUserProcedure,
		svc.LoginUser,
		opts...,
	)
	userServiceGetUserAccountHandler := connect.NewUnaryHandler(
		UserServiceGetUserAccountProcedure,
		svc.GetUserAccount,
		opts...,
	)
	return "/api.user.v1.UserService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case UserServiceRegisterUserProcedure:
			userServiceRegisterUserHandler.ServeHTTP(w, r)
		case UserServiceLoginUserProcedure:
			userServiceLoginUserHandler.ServeHTTP(w, r)
		case UserServiceGetUserAccountProcedure:
			userServiceGetUserAccountHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedUserServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUserServiceHandler struct{}

func (UnimplementedUserServiceHandler) RegisterUser(context.Context, *connect.Request[v1.RegisterUserRequest]) (*connect.Response[v1.RegisterUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.user.v1.UserService.RegisterUser is not implemented"))
}

func (UnimplementedUserServiceHandler) LoginUser(context.Context, *connect.Request[v1.LoginUserRequest]) (*connect.Response[v1.LoginUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.user.v1.UserService.LoginUser is not implemented"))
}

func (UnimplementedUserServiceHandler) GetUserAccount(context.Context, *connect.Request[v1.GetUserAccountRequest]) (*connect.Response[v1.GetUserAccountResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.user.v1.UserService.GetUserAccount is not implemented"))
}
