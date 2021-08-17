package handler

import (
	"encoding/json"
	"experiments/internal/app/model"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type Handler struct {
	logger *logrus.Logger
}

func New() *Handler {
	return &Handler{
		logger: logrus.New(),
	}
}

func (handler *Handler) HandleRequestAddressHealth(writer http.ResponseWriter, request *http.Request) {
	handler.logger.Info(fmt.Sprintf("Server accepted connection from [%v] method [%s] uri [%s]", request.RemoteAddr, request.Method, request.RequestURI))
	_ = model.Extract(request)

	handler.sendResponse(writer, http.StatusOK, nil, "OK")
}

func (handler *Handler) HandleHealth(writer http.ResponseWriter, request *http.Request) {
	handler.logger.Info(fmt.Sprintf("Server accepted connection from [%v] method [%s]", request.RemoteAddr, request.Method))
	_, err := io.WriteString(writer, "OK")
	if err != nil {
		handler.sendResponse(writer, http.StatusInternalServerError, err, nil)
	}
}

func (handler *Handler) sendResponse(writer http.ResponseWriter, code int, err error, result interface{}) {
	js, _ := json.Marshal(result)
	if err != nil {
		http.Error(writer, err.Error(), code)
		handler.logger.Error(err)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	_, err = writer.Write(js)
	if err != nil {
		handler.logger.Error(err)
		return
	}
}
