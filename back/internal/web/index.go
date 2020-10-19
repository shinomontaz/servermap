package web

import (
	"encoding/json"
	"net/http"
	"servermap/i8s"
)

type Message struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

func (s *Service) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// return json tree for hosts with attached vms each

	vmTypes, ok := r.URL.Query()["vmtype"]
	var vmType string

	if ok {
		vmType = vmTypes[0]
	}

	hosts := s.filterByType(vmType)

	json.NewEncoder(w).Encode(hosts)
}

func (s *Service) filterByType(vmType string) []*i8s.Host {
	res := make([]*i8s.Host, 0, len(s.hosts))

	for _, host := range s.hosts {
		filteredHost := host.Filter(vmType)
		if vmType == "" || len(filteredHost.Vms) > 0 {
			res = append(res, &filteredHost)
		}
	}

	return res
}
