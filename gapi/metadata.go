package gapi

import (
	"context"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	grpcGatewayUserAgentHeader  = "grpcgateway-user-agent"
	userAgentHeader 			= "user-agent"
	xForwardedForHeader 		= "x-forwarded-for"
)

type Metadata struct {
	UserAgent string
	ClinetIP  string
}

func (server *Server) extraMetadata(ctx context.Context) *Metadata {
	mtdt := &Metadata{}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if UserAgents := md.Get(grpcGatewayUserAgentHeader); len(UserAgents) > 0 {
			mtdt.UserAgent = UserAgents[0]
		}

		if UserAgents := md.Get(userAgentHeader); len(UserAgents) > 0 {
			mtdt.UserAgent = UserAgents[0]
		}


		if clientIPs := md.Get(xForwardedForHeader); len(clientIPs) > 0 {
			mtdt.ClinetIP = clientIPs[0]
		}
	}

	// getting clientIP for grpc
	if p, ok := peer.FromContext(ctx); ok {
		mtdt.ClinetIP = p.Addr.String()
	}

	return mtdt
}