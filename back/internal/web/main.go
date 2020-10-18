package web

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"servermap/i8s"
)

type Service struct {
	erh      i8s.IErrorHandler
	hosts    []*i8s.Host
	vms      []*i8s.Vm
	hIndex   map[string]map[string][]*i8s.Host
	hIdIndex map[string]*i8s.Host
	vIndex   map[string]map[string][]*i8s.Vm
}

func New(erh i8s.IErrorHandler) *Service {
	return &Service{erh: erh, hosts: make([]*i8s.Host, 0), vms: make([]*i8s.Vm, 0)}
}

func (s *Service) InitData(hFile, vmFile string) {
	s.hIndex = make(map[string]map[string][]*i8s.Host)
	s.hIdIndex = make(map[string]*i8s.Host)
	s.hIndex["Name"] = make(map[string][]*i8s.Host)
	s.hIndex["Address"] = make(map[string][]*i8s.Host)
	s.hIndex["Os"] = make(map[string][]*i8s.Host)

	s.vIndex = make(map[string]map[string][]*i8s.Vm)
	s.vIndex["Name"] = make(map[string][]*i8s.Vm)
	s.vIndex["Os"] = make(map[string][]*i8s.Vm)

	s.initHosts(hFile)
	s.initVms(vmFile)
}

func (s *Service) initHosts(filename string) {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	xmlHosts := i8s.XMLHosts{}

	err = xml.Unmarshal([]byte(data), &xmlHosts)
	if err != nil {
		panic(err)
	}

	xmlHosts.Hosts = append(xmlHosts.Hosts, i8s.XMLHost{Name: "nonexist", ID: "nonexist"})

	for _, xHost := range xmlHosts.Hosts {
		h := i8s.HostFromXml(xHost)
		s.hosts = append(s.hosts, &h)
		s.hIdIndex[h.ID] = &h
		s.hIndex["Name"][h.Name] = append(s.hIndex["Name"][h.Name], &h)
		s.hIndex["Address"][h.Address] = append(s.hIndex["Address"][h.Address], &h)
		s.hIndex["Os"][h.Os.Type] = append(s.hIndex["Os"][h.Os.Type], &h)
	}
}

func (s *Service) initVms(filename string) {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	xmlVms := i8s.XMLVms{}

	err = xml.Unmarshal([]byte(data), &xmlVms)
	if err != nil {
		panic(err)
	}

	for _, xVm := range xmlVms.Vms {
		v := i8s.VmFromXml(xVm, s.hIdIndex)
		s.hIdIndex[v.Host].Vms = append(s.hIdIndex[v.Host].Vms, &v)
		fmt.Println("host for VM: ", v.ID, v.Host)

		s.vms = append(s.vms, &v)

		// jsonData, err := json.Marshal(s.hIdIndex[v.Host])
		// if err != nil {
		// 	log.Println(err)
		// }
		// fmt.Println(string(jsonData))

		s.vIndex["Name"][v.Name] = append(s.vIndex["Name"][v.Name], &v)
		s.vIndex["Os"][v.Os] = append(s.vIndex["Os"][v.Os], &v)
	}

	// jsonData, err := json.Marshal(s.hIdIndex)
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(string(jsonData))
}
