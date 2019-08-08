package main

import (
	context "context"
	"encoding/json"
	config "github.com/Basic-Components/qrcodegenerator/config"
	pb "github.com/Basic-Components/qrcodegenerator/qrcodegeneratordeclare"
	"github.com/Basic-Components/qrcodegenerator/generator"
	"github.com/Basic-Components/qrcodegenerator/logger"
	"net"

	grpc "google.golang.org/grpc"
	logrus "github.com/sirupsen/logrus"
)
type rpcserver struct {
}

var encodelog *logrus.Entry= logger.Log.WithFields(logrus.Fields{
	"API": "Encode",
})


func (s *rpcserver) Encode(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	encodelog.WithFields(logrus.Fields{
		"data": in.Data,
	}).Debug("get request")
	result, err := generator.Encode(in.Data)
	var status pb.StatusData
	if err != nil {
		status = pb.StatusData{
			Status: pb.StatusData_ERROR,
			Msg:    err.Error()}
		encodelog.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Debug("send response")
		return &pb.Response{
			Status: &status}, nil
	}
	status = pb.StatusData{
		Status: pb.StatusData_SUCCEED,
		Msg:    "success"}
	encodelog.WithFields(logrus.Fields{
		"response": string(result),
	}).Debug("send response")
	return &pb.Response{
		Status: &status,
		Data:   result}, nil
	
}

// Run 执行签名验签服务
func Run(conf config.ConfigType) {
	rpc := rpcserver{}
	listener, err := net.Listen("tcp", conf.Address)
	if err != nil {
		logger.Log.Fatalf("failed to listen: %v", err)
		return
	}
	logger.Log.Info("server started @", conf.Address)
	server := grpc.NewServer()
	pb.RegisterQrServiceServer(server, &rpc)
	if err := server.Serve(listener); err != nil {
		logger.Log.Fatalf("failed to serve: %v", err)
		return
	}
}

func main() {
	conf,err := config.Init()
	if err != nil{
		logger.Log.Fatalf("load config error : %v", err)
	}else{
		logger.Log.Info("start server config %v", conf)
		Run(conf)
	}
}