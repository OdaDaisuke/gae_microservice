package gateway

import (
	"github.com/OdaDaisuke/gae_sand/services/main/config"
	"google.golang.org/grpc"
)

type GrpcGatewayFactory struct {
	AccountConn *grpc.ClientConn
	ContentConn *grpc.ClientConn
}

func NewGrpcGatewayFactory(hosts *config.GrpcHosts) (*GrpcGatewayFactory, error) {
	accountConn, err := grpc.Dial(hosts.Account.Build(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	contentConn, err := grpc.Dial(hosts.Content.Build(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &GrpcGatewayFactory{
		AccountConn: accountConn,
		ContentConn: contentConn,
	}, nil
}

