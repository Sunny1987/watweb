package server

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"webapp/datamodels"
)

func (l *AppLogger) ScanForFile(rw http.ResponseWriter, r *http.Request) bool {
	//client := resty.Client{}
	var fresp datamodels.FinalResponse
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(rw, "Error max file size exceeded", http.StatusBadRequest)
		return true
	}
	file, handler, err := r.FormFile("file")
	l.myLogger.Printf("Uploaded File: %+v\n", handler.Filename)
	l.myLogger.Printf("File Size: %+v\n", handler.Size)
	l.myLogger.Printf("MIME Headers: %+v\n", handler.Header)
	//payload := NewPayloadFile(file)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", handler.Filename)
	if err != nil {
		return true
	}
	_, err = io.Copy(part, file)
	err = writer.Close()
	if err != nil {
		return true
	}

	resp, err := http.Post(WATUPLOADHTMLLOCAL, writer.FormDataContentType(), body)
	defer resp.Body.Close()
	if err != nil {
		l.myLogger.Println(err)
	}
	json.NewDecoder(resp.Body).Decode(&fresp)
	if len(fresp.Results) == 1 {
		l.myLogger.Println(fresp.Request.FileName)
	} else {

	}
	return false
}

func (l *AppLogger) ScanForURL(url string, depth string) bool {
	var fresp datamodels.FinalResponse
	var jsonPayload []byte
	l.myLogger.Println(url)
	l.myLogger.Println(depth)
	payload := datamodels.NewPayload(url, depth)
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		l.myLogger.Println(err)
	}
	resp, err := http.Post(WATSCANURLLOCAL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		l.myLogger.Println(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		l.myLogger.Println("Error: Unexpected status code:", resp.StatusCode)
		return true
	}
	finalResp, err := json.MarshalIndent(&fresp, "", " ")
	if err != nil {
		l.myLogger.Println(err)
	}

	l.myLogger.Println(string(finalResp))
	return false
}
