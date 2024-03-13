package server

import (
	"log"
	"net/http"
)

type RouterHandler interface {
	StartApp(rw http.ResponseWriter, r *http.Request)
	StartScan(rw http.ResponseWriter, r *http.Request)
}

type AppLogger struct {
	myLogger *log.Logger
}

func NewAppLogger(myLogger *log.Logger) RouterHandler {
	return &AppLogger{myLogger: myLogger}
}
