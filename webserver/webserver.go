package main

import (
	"ImageSlip/commonutils"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

var log *zap.SugaredLogger

func init() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	log = logger.Sugar()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Home).Methods("GET")
	r.HandleFunc("/file/{key}/{value}", DownloadFile).Methods("GET")
	log.Infof("listening for http requests on %d", 8080)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Error("failed to start web server")
		panic(err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	log.Info("Received home url request")
	res := "Welcome to File Downloader. Add /file/<id> path to download file"
	_, err := w.Write([]byte(res))
	if err != nil {
		log.Error("fatal: error while writing response")
	}
}

/*
	DownloadFile... downloads file requested via key
*/
func DownloadFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileDir := vars["key"]
	fileName := vars["value"]
	log.Info("Received file download request for", fileDir,"/", fileName)

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "Image/png")

	_, _, buf := commonutils.GetFileContent("/tmp/" + fileDir + "/" + fileName)
	_, err := w.Write(buf)
	if !commonutils.IsError(err) {
		log.Error("error occured while writing response")
	}
	log.Info("Sending file")
}
