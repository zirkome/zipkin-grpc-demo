package main

import (
	"fmt"
	"net"
	"os"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	opentracing "github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
	"google.golang.org/grpc"

	"github.com/Sirupsen/logrus"
	"github.com/kokaz/zipkin-grpc-demo/cmd/beta/service"
)

const (
	// Our service name.
	serviceName = "beta"

	// Host + port of our service.
	hostPort = "127.0.0.1:9001"

	// Endpoint to send Zipkin spans to.
	zipkinHTTPEndpoint = "http://localhost:9411/api/v1/spans"

	// Debug mode.
	debug = true

	// same span can be set to true for RPC style spans (Zipkin V1) vs Node style (OpenTracing)
	sameSpan = false
)

func main() {
	collector, err := zipkin.NewHTTPCollector(zipkinHTTPEndpoint)
	if err != nil {
		fmt.Printf("unable to create Zipkin HTTP collector: %+v", err)
		os.Exit(-1)
	}

	// create recorder.
	recorder := zipkin.NewRecorder(collector, debug, hostPort, serviceName)

	// create tracer.
	tracer, err := zipkin.NewTracer(
		recorder, zipkin.ClientServerSameSpan(sameSpan),
	)
	if err != nil {
		fmt.Printf("unable to create Zipkin tracer: %+v", err)
		os.Exit(-1)
	}

	// explicitely set our tracer to be the default tracer.
	opentracing.InitGlobalTracer(tracer)

	svc := &service.BetaServer{}
	srv := grpc.NewServer(grpc.UnaryInterceptor(
		otgrpc.OpenTracingServerInterceptor(tracer, otgrpc.LogPayloads())))
	service.RegisterBetaServiceServer(srv, svc)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9001))
	if err != nil {
		logrus.WithError(err).Error("net.Listen")
	}
	err = srv.Serve(lis)
	if err != nil {
		logrus.WithError(err).Error("srv.Serve")
	}
}
