package main

import (
	"ImageSlip/commonutils"
	"context"
	"errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"os"
)

var log *zap.SugaredLogger

func init() {
	log = commonutils.Log
}

func main() {
	log.Debug("starting client connection to server")
	filePath := os.Args[1]
	req := CreateRequest(filePath)
	log.Info("Sending message to server")
	SendToServer(req)
}

/*
	SendToServer... boiler plate code to connect to gRPC server
*/
func SendToServer(req *commonutils.MessageRequest) {
	serverAddress := ":39298"
	log.Debug("listening on ", serverAddress)
	conn, e := grpc.Dial(serverAddress, grpc.WithInsecure())

	if e != nil {
		panic(e)
	}
	defer conn.Close()

	client := commonutils.NewPostMessageClient(conn)

	if respMsg, e := client.SendMessage(context.Background(), req); e != nil {
		log.Error("not able to send message to server")
	} else {
		log.Info("response message received is %#v \n", respMsg)
		log.Info(respMsg.GetResponse())
	}
}

/*
	CreateRequest... constructs MessageRequest defined in protobuf
	Reads file from given filepath into MessageRequest
	invokes getMetaData()
*/
func CreateRequest(filePath string) *commonutils.MessageRequest {
	log.Debug("creating message request")
	err, mData, buf := commonutils.GetFileContent(filePath)
	err = validate(buf)
	if err != nil {
		panic(err)
	}

	req := commonutils.MessageRequest{File: buf, MetaData: mData}
	return &req
}

/*
	validate... perform validations
*/
func validate(buf []byte) error {
	//isValid:= isValidImageType(mData.FileName)
	isValid := commonutils.IsValidContent(buf)
	if !isValid {
		return errors.New("validation failed : invalid content type, accepts only png | jpg | gif")
	}
	return nil
}
