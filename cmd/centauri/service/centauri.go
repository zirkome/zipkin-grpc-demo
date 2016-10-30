package service

import (
	"golang.org/x/net/context"

	"github.com/Sirupsen/logrus"
	beta "github.com/kokaz/zipkin-grpc-demo/cmd/beta/service"
)

type CentauriServer struct {
	betaClient beta.BetaServiceClient
}

func NewCentauriServer(betaClient beta.BetaServiceClient) *CentauriServer {
	return &CentauriServer{
		betaClient: betaClient,
	}
}

func (s *CentauriServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	res, err := s.betaClient.Set(ctx, &beta.SetRequest{
		Data: req.Data,
	})
	if err != nil {
		logrus.WithError(err).Error("s.betaClient.Set")
	}

	return &GetResponse{
		Data: res.Data + " GETCENTAURI",
	}, nil
}
