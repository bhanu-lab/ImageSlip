package commonutils

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"strings"

	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

func init() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // flushes buffer, if any
	Log = logger.Sugar()
}

/*
	GetFileContent... reads file and returns its content as bytes
	and file info as MetaData
	returns Metadata, []byte
*/
func GetFileContent(filePath string) (error, *MetaData, []byte) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 643)
	if !IsError(err) {
		panic(err)
	}
	defer file.Close()

	info, err := file.Stat()
	if !IsError(err) {
		panic(err)
	}
	if info.IsDir() {
		panic("Cant upload a directory. Please select a file")
	}
	mData := getMetaData(info)
	buf := getFileBytes(file)
	return err, mData, buf
}

/*
	getFileBytes... convert file given into bytes for transport
	---> file *os.File
	<--- []byte
*/
func getFileBytes(file *os.File) []byte {
	var buf []byte
	r := bufio.NewReader(file)
	b := make([]byte, 10)
	for {
		n, err := r.Read(b)
		if err == io.EOF { //handling end of file error
			break
		}
		if err != nil { // avoid
			Log.Error("Error reading file:", err)
			break
		}
		buf = append(buf, b[0:n]...)
	}
	return buf
}

/*
	getMetaData... collects metadata of the file passed
	fills *commonutils.MetaData with
			FileName
			Size
			LastModified
*/
func getMetaData(info os.FileInfo) *MetaData {
	mData := MetaData{}
	mData.FileName = info.Name()
	mData.Size = info.Size()

	Log.Debug("FileName: [%s] FileSize: [%d] FileModifiedTime: [%v] \n", mData.FileName, mData.Size, mData.Lastmodified)
	return &mData
}

/*
	IsError... if error is not nil prints error
	returns True : if error is nil
			False : if error is not nil
*/
func IsError(err error) bool {
	if err != nil {
		Log.Error(err.Error())
	}
	return err == nil
}

/*
	isValidContent... checks for file content
	returns True : if content is image and type png | jpg | gif
			False : if any other content type
*/
func IsValidContent(buf []byte) bool {
	contentType := http.DetectContentType(buf)
	Log.Info("content type detected is ", contentType)
	if contentType == "image/png" || contentType == "image/jpg" {
		return true
	}
	return false
}

/*
	returns True : if image is of type 'png' | 'jpg' | 'gif'
			False : if not of above-mentioned types
*/
func IsValidImageType(fileName string) bool {
	s := strings.Split(fileName, ".")
	if (len(s) > 1 && s[1] == "") || len(s) == 1 {
		Log.Debug("couldn't detect file type")
		return false
	}

	if len(s) > 1 {
		ext := s[len(s)-1]
		if ext == "png" || ext == "jpg" || ext == "gif" {
			Log.Debug("received file type ", ext)
			return true
		}
	}

	return false

}
