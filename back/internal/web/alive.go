package web

import (
	"encoding/json"
	"net/http"
)

type AliveMessage struct {
	Alive bool
}

func (s *Service) Alive(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(AliveMessage{Alive: true})
}
