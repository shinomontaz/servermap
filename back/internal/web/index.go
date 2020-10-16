package web

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

func (s *Service) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// return json tree for hosts with attached vms each

	json.NewEncoder(w).Encode(s.hosts)

	//	json.NewEncoder(w).Encode(Message{Text: "Ovirt viewer", Type: i8s.TypeSUCCESS})
}
