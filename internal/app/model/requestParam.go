package model

import (
	"github.com/gorilla/mux"
	"net/http"
)

type RequestData struct {
	Address  string
	User     string
	Password string
}

func Extract(request *http.Request) *RequestData {
	return &RequestData{
		mux.Vars(request)["address"],
		request.Header.Get("user"),
		request.Header.Get("password"),
	}
}
