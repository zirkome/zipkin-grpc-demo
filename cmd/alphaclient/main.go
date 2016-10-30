package main

import (
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/Sirupsen/logrus"
	alpha "github.com/kokaz/zipkin-grpc-demo/cmd/alpha/service"
)

func main() {
	alphaConn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		logrus.WithError(err).Error("grpc.Dial")
	}

	alphaClient := alpha.NewAlphaServiceClient(alphaConn)

	ctx := context.Background()
	res, err := alphaClient.Get(ctx, &alpha.GetRequest{
		Data: "hola",
	})
	if err != nil {
		logrus.WithError(err).Error("alphaClient.Get")
		return
	}

	fmt.Printf("%v", res.Data)
}
