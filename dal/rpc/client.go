package rpc

import (
	plog_gateway "plog_client/proto"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	addr = ":9999"
)

var (
	plogClient plog_gateway.PLogGatewayClient
)

func init() {
	plogClient = plog_gateway.NewPLogGatewayClient(getConn(addr))
}

func getConn(addr string) *grpc.ClientConn {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		logrus.Fatalf("failed to dial: %+v", err)
	}
	return conn
}
