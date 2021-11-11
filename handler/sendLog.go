package handler

import (
	"context"
	"plog_client/dal/rpc"
	plog_gateway "plog_client/proto"

	"github.com/sirupsen/logrus"
)

func SendLog(msg string) (resp *plog_gateway.UploadLogResponse, err error) {
	ctx := context.Background()
	log := CreateLog(msg)
	resp, err = rpc.UploadLog(ctx, log)
	if err != nil {
		logrus.Errorf("SendLog failed: %+v", err)
	}
	return
}

func CreateLog(msg string) *plog_gateway.Log {
	return &plog_gateway.Log{
		Message: msg,
	}
}
