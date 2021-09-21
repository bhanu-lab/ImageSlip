package main

import (
	"ImageSlip/commonutils"
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

type PostMessageServer struct {
	commonutils.UnimplementedPostMessageServer
}

var log *zap.SugaredLogger
var config *commonutils.Config

func init() {
	log = commonutils.Log
	localConfig, err := commonutils.GetConfig()
	config = localConfig
	if err != nil {
		log.Error("unable to read config")
		panic("failed to read config")
	}
}

/*
	SendMessage... implementation for grpc server SendMessage refer
	commonutils/Botservice_grpc.pb.go
*/
func (pms *PostMessageServer) SendMessage(ctx context.Context, req *commonutils.MessageRequest) (*commonutils.MessageResponse, error) {
	res := new(commonutils.MessageResponse)
	fileName := req.MetaData.FileName
	log.Debug("recevied file name is %s and size\n", fileName)
	fileID, err := createFile(fileName, req)
	if err != nil {
		//fmt.Errorf("failed while writing data to file")
		res.Response = fmt.Sprintf("failed to upload %s", fileName)
	} else {
		fileShareURL := config.FileHost
		res.Response = fmt.Sprintf("uploaded %s successfully, use URl:%s for downloading", fileName, fileShareURL+fileID)
	}
	return res, nil
}

/*
	createFile... creates file with content received in fileStorePath
	inside directory with name <RANDOM_EIGHT_DIGIT> num and file name received
	/fileStorePath/<RANDOM_EIGHT_DIGIT>/file_name
*/
func createFile(fileName string, req *commonutils.MessageRequest) (string, error) {
	randomDirName := genRandNum()
	path := config.FileStorePath + randomDirName
	log.Info("creating file " + path + "/" + fileName)
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			log.Errorf("couldnt create directory", config.FileStorePath+randomDirName)
			return "", err
		}
	}
	f, err := os.Create(path + "/" + fileName)
	if err != nil {
		log.Fatal("failed to create file", err.Error())
		return "", err
	}
	n2, err := f.Write(req.File)
	log.Debugf("wrote %d bytes\n", n2)
	return randomDirName + "/" + fileName, err
}

/*
	getRandNum... generates 8 digit random number
*/
func genRandNum() string {
	min := 10000000
	max := 99999999
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(max-min) + min)
}

func main() {
	log.Debug("in grpc server.... trying to start server")
	grpcPort := config.GRPCPort
	netListener := getNetListener(grpcPort)
	grpcServer := grpc.NewServer()

	pms := new(PostMessageServer)
	commonutils.RegisterPostMessageServer(grpcServer, pms)

	//start the server
	log.Info("starting grpc server")
	if err := grpcServer.Serve(netListener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

/*
	getNetListener... boiler plate code for returning net.Listener
*/
func getNetListener(port int) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
	}
	return lis
}
