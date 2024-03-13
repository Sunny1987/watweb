package server

import (
	"html/template"
	"net/http"
)

const (
	WATSCANURL         = "https://watserver-1.onrender.com/scan"
	WATSCANURLLOCAL    = "http://localhost:8080/scan"
	WATUPLOADHTMLLOCAL = "http://localhost:8080/uploadhtml"
)

func (l *AppLogger) StartApp(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	err := tmpl.Execute(rw, nil)
	if err != nil {
		l.myLogger.Println(err)
	}
}

func (l *AppLogger) StartScan(rw http.ResponseWriter, r *http.Request) {
	url := r.PostFormValue("urlValue")
	depth := r.PostFormValue("depth")

	if url != "" {
		if l.ScanForURL(url, depth) {
			return
		}
	} else {
		if l.ScanForFile(rw, r) {
			return
		}
	}

}
