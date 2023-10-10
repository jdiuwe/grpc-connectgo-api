package handlers

import (
	"context"
	"database/sql"
	"errors"

	"connectrpc.com/connect"
	"github.com/gofrs/uuid"
	db "github.com/grpc-connectgo-api-demo/wallet/db/sqlc"
	v1 "github.com/grpc-connectgo-api-demo/wallet/gen/go/api/user/v1"
	"github.com/grpc-connectgo-api-demo/wallet/gen/go/api/user/v1/userv1connect"
	"github.com/grpc-connectgo-api-demo/wallet/pkg/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	userv1connect.UnimplementedUserServiceHandler
	store      *db.Store
	jwtManager *jwt.Manager
}

func NewUserHandler(st *db.Store, jwtManager *jwt.Manager) *UserHandler {
	return &UserHandler{store: st, jwtManager: jwtManager}
}

func (h UserHandler) RegisterUser(ctx context.Context, request *connect.Request[v1.RegisterUserRequest]) (*connect.Response[v1.RegisterUserResponse], error) {
	if err := request.Msg.Validate(); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	req := request.Msg

	// check if user exists
	user, err := h.store.UserEmailExists(ctx, req.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if u, ok := user.(int64); ok && u == 1 {
		return nil, connect.NewError(connect.CodeAlreadyExists, ErrUserExists)
	}

	uid, err := uuid.NewV4()
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	hashedPassword, err := HashPassword(req.GetPassword())
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, ErrFailedToHashPass)
	}

	// create user
	newUser, err := h.store.CreateUser(ctx, db.CreateUserParams{
		Uuid:           uid.String(),
		Firstname:      req.GetFirstName(),
		Lastname:       req.GetLastName(),
		Email:          req.GetEmail(),
		HashedPassword: hashedPassword,
	})
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&v1.RegisterUserResponse{
		Message: "user account created",
		Status:  "success",
		Uuid:    newUser.Uuid,
	}), nil
}

func (h UserHandler) LoginUser(ctx context.Context, request *connect.Request[v1.LoginUserRequest]) (*connect.Response[v1.LoginUserResponse], error) {
	req := request.Msg

	if err := req.Validate(); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	user, err := h.store.GetUserByEmail(ctx, req.GetEmail())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, ErrUserNotFound.Error())
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if err = CheckPassword(req.GetPassword(), user.HashedPassword); err != nil {
		return nil, status.Error(codes.Unauthenticated, ErrBadPassword.Error())
	}

	accessToken, err := h.jwtManager.Generate(user.Uuid)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&v1.LoginUserResponse{
		Message:     "user logged in",
		Status:      "success",
		AccessToken: accessToken,
	}), nil
}

func (h UserHandler) GetUserAccount(ctx context.Context, request *connect.Request[v1.GetUserAccountRequest]) (*connect.Response[v1.GetUserAccountResponse], error) {
	req := request.Msg

	if err := req.Validate(); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	user, err := h.store.GetUserByUUID(ctx, req.GetUuid())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, ErrUserNotFound.Error())
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&v1.GetUserAccountResponse{
		UserAccount: &v1.UserAccount{
			Uuid:          user.Uuid,
			FirstName:     user.Firstname,
			LastName:      user.Lastname,
			Email:         user.Email,
			EmailVerified: user.Verified.Bool,
		},
		Status: "success",
	}), nil
}
