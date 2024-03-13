package datamodels

import (
	"io"
	"mime/multipart"
)

type Payload struct {
	URL   string `json:"url"`
	Depth string `json:"depth"`
	File  multipart.File
}

func NewPayloadFile(file multipart.File) *Payload {
	return &Payload{File: file}
}

type FinalResponse struct {
	Request *MyRequest  `json:"request"`
	Person  *string     `json:"person"`
	Results []TagResult `json:"results"`
	//Doc       `json:"-"`
}

type MyRequest struct {
	URL      string    `json:"url" validate:"required,url"`
	Depth    int       `json:"depth" validate:"gt=-1,lte=2"`
	File     io.Reader `json:"file"`
	FileName string    `json:"fileName"`
}

type Result struct {
	Guideline string   `json:"guideline"`
	Rules     []string `json:"rules"`
}

type TagResult struct {
	Tag    string   `json:"tag"`
	Result []Result `json:"result"`
}

func NewPayload(URL string, depth string) *Payload {
	return &Payload{URL: URL, Depth: depth}
}
