package gapi

import (
	"context"
	"database/sql"
	"time"

	db "github.com/eugenius-watchman/golang_simplebank/db/sqlc"
	"github.com/eugenius-watchman/golang_simplebank/pb"
	"github.com/eugenius-watchman/golang_simplebank/util"
	"github.com/eugenius-watchman/golang_simplebank/val"

	// "github.com/lib/pq"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	// authorization
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	// validations
	if authPayload.Username != req.GetUsername() {
		return nil, status.Errorf(codes.PermissionDenied, "unauthorized: users can only update their own profile")
	}

	violations := validateUpdateUserRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	arg := db.UpdateUserParams{
		Username:       req.GetUsername(),
		FullName: sql.NullString{
			String: req.GetFullname(),
			Valid: req.Fullname != nil,
		},
		Email: sql.NullString{
			String: req.GetEmail(),
			Valid: req.Email != nil,
		},
	}

	if req.Password != nil {
		// hashed password
		hashedPassword, err := util.HashPassword(req.GetPassword())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
		}

		arg.HashedPassword = sql.NullString{
			String: hashedPassword,
			Valid: true,
		}

		arg.PasswordChangedAt = sql.NullTime{
			Time: time.Now(),
			Valid: true,
		}
	}

	user, err := server.store.UpdateUser(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to update user: %s", err)
	}

	rsp := &pb.UpdateUserResponse{
		User: convertUser(user),
	}
	return rsp, nil
}


func validateUpdateUserRequest(req *pb.UpdateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	if req.Password != nil {
		if err := val.ValidatePassword(req.GetPassword()); err != nil {
			violations = append(violations, fieldViolation("password", err))
		}
	}
	
	if req.Fullname != nil {
		if err := val.ValidateFullname(req.GetFullname()); err != nil {
			violations = append(violations, fieldViolation("full_name", err))
		}
	}
	
	if req.Email != nil {
		if err := val.ValidateEmail(req.GetEmail()); err != nil {
			violations = append(violations, fieldViolation("email", err))
		}
	}
	
	return violations
}

