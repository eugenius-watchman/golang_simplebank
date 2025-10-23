package gapi

import (
	// "context"
	"fmt"

	db "github.com/eugenius-watchman/golang_simplebank/db/sqlc"
	"github.com/eugenius-watchman/golang_simplebank/pb"
	"github.com/eugenius-watchman/golang_simplebank/token"
	"github.com/eugenius-watchman/golang_simplebank/util"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
 )

// Server to serve HTTP requests for bank services
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer instance to setup/create a new HTTP/API server routing for services on the server
func NewServer(config util.Config, store db.Store) (*Server, error) {
	// tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}

// CreateUser implements the CreateUser RPC
// func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
// 	// For now, return unimplemented - you can add the actual logic later
// 	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
// }

// // LoginUser implements the LoginUser RPC
// func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
// 	// For now, return unimplemented - you can add the actual logic later
// 	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
// }