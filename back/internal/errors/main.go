package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"servermap/i8s"

	log "github.com/sirupsen/logrus"
)

type Service struct {
}

type Message struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

func New() *Service {
	return &Service{}
}

func (s *Service) HandleError(w http.ResponseWriter, err error) {
	_, filename, lineno, ok := runtime.Caller(1)
	message := ""
	if ok {
		message = fmt.Sprintf("%v:%v: %v\n", filename, lineno, err)
	}
	log.Warning(message, err)
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(Message{Text: fmt.Sprintf("%s", err), Type: i8s.TypeDANGER})
}

func (s *Service) Set401(w http.ResponseWriter, str string) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(str))
}
