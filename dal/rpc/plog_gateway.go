package rpc

import (
	"context"
	plog_gateway "plog_client/proto"

	"github.com/sirupsen/logrus"
)

func UploadLog(ctx context.Context, log *plog_gateway.Log) (resp *plog_gateway.UploadLogResponse, err error) {
	req := &plog_gateway.UploadLogRequest{
		Log: log,
	}
	resp, err = plogClient.UploadLog(ctx, req)
	if err != nil {
		logrus.Warnf("FAILED TO UploadLog: %+v", err)
		return
	}
	logrus.Infof("UploadLog, log: %+v\n", log)
	return
}
