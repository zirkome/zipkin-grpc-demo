package service

import "golang.org/x/net/context"

type BetaServer struct{}

func (s *BetaServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return &GetResponse{
		Data: " GETBETA",
	}, nil
}

func (s *BetaServer) Set(ctx context.Context, req *SetRequest) (*SetResponse, error) {
	return &SetResponse{
		Data: " SETBETA",
	}, nil
}
