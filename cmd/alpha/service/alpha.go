package service

import (
	"time"

	"golang.org/x/net/context"

	"github.com/Sirupsen/logrus"
	beta "github.com/kokaz/zipkin-grpc-demo/cmd/beta/service"
	centauri "github.com/kokaz/zipkin-grpc-demo/cmd/centauri/service"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

type AlphaServer struct {
	betaClient     beta.BetaServiceClient
	centauriClient centauri.CentauriServiceClient
}

func NewAlphaServer(betaClient beta.BetaServiceClient, centauriClient centauri.CentauriServiceClient) *AlphaServer {
	return &AlphaServer{
		betaClient:     betaClient,
		centauriClient: centauriClient,
	}
}

func (s *AlphaServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	span := opentracing.SpanFromContext(ctx)

	span.LogFields(log.Int64("now", time.Now().Unix()))

	chBeta := make(chan *beta.GetResponse)
	chCentauri := make(chan *centauri.GetResponse)

	go func() {
		res, err := s.betaClient.Get(ctx, &beta.GetRequest{
			Data: req.Data,
		})
		if err != nil {
			logrus.WithError(err).Error("s.betaClient.Get")
		}
		chBeta <- res
	}()

	go func() {
		res, err := s.centauriClient.Get(ctx, &centauri.GetRequest{
			Data: req.Data,
		})
		if err != nil {
			logrus.WithError(err).Error("s.centauriClient.Get")
		}
		chCentauri <- res
	}()

	resBeta := <-chBeta
	resCentauri := <-chCentauri

	return &GetResponse{
		Data: resBeta.Data + " " + resCentauri.Data + " GETALPHA",
	}, nil
}
