package i8s

type Host struct {
	ID                  string
	Name                string
	Comment             string
	Address             string
	Cpu                 Cpu
	ExternalStatus      string
	Hardware            Hardware
	MaxSchedulingMemory int64
	Memory              int64
	Os                  Os
	Port                int
	Status              string
	Summary             int
	Type                string
	Version             string // `xml:"full_version"`
	Cluster             string
	Vms                 []*Vm
}

func HostFromXml(xHost XMLHost) Host {
	return Host{
		ID:                  xHost.ID,
		Name:                xHost.Name,
		Comment:             xHost.Comment,
		Address:             xHost.Address,
		Cpu:                 CpuFromXml(xHost.Cpu),
		ExternalStatus:      xHost.Name,
		Hardware:            xHost.HardwareInformation.Hardware,
		MaxSchedulingMemory: xHost.MaxSchedulingMemory,
		Memory:              xHost.Memory,
		Os:                  OsFromXml(xHost.Os),
		Port:                xHost.Port,
		Status:              xHost.Status,
		Summary:             xHost.Summary.Active,
		Type:                xHost.Type,
		Version:             xHost.Version.FullVersion,
		Cluster:             xHost.Cluster.ID,
		Vms:                 make([]*Vm, 0),
	}
}

type XMLHosts struct {
	Hosts []XMLHost `xml:"host"`
}

type XMLHost struct {
	ID             string `xml:"id,attr"` //="1d312aac-c09a-4231-8737-249cef1bae7d"
	Name           string `xml:"name"`
	Comment        string `xml:"comment"`
	Address        string `xml:"address"`
	AutoNumaStatus string `xml:"auto_numa_status"`
	Certificate    struct {
		Organization string `xml:"organization"`
		Subject      string `xml:"subject"`
	} `xml:"certificate"`
	Cpu               XMLHostCpu `xml:"cpu"`
	DevicePassthrough struct {
		Enabled string `xml:"enabled"`
	} `xml:"device_passthrough"`
	ExternalStatus      string          `xml:"external_status"`
	HardwareInformation XMLHostHardware `xml:"hardware_information"`
	Iscsi               struct {
		Initiator string `xml:"initiator"`
	} `xml:"iscsi"`
	KdumpStatus string `xml:"kdump_status"`
	Ksm         struct {
		Enabled string `xml:"enabled"`
	} `xml:"ksm"`
	LibvirtVersion struct {
		Build       string `xml:"build"`
		FullVersion string `xml:"full_version"`
		Major       string `xml:"major"`
		Minor       string `xml:"minor"`
		Revision    string `xml:"revision"`
	} `xml:"libvirt_version"`
	MaxSchedulingMemory int64     `xml:"max_scheduling_memory"`
	Memory              int64     `xml:"memory"`
	NumaSupported       string    `xml:"numa_supported"`
	Os                  XMLHostOs `xml:"os"`
	Port                int       `xml:"port"`
	PowerManagement     struct {
		AutomaticPmEnabled string `xml:"automatic_pm_enabled"`
		Enabled            string `xml:"enabled"`
		KdumpDetection     string `xml:"kdump_detection"`
		PmProxies          string `xml:"pm_proxies"`
	} `xml:"power_management"`
	Protocol string `xml:"protocol"`
	SeLinux  struct {
		Mode string `xml:"mode"`
	} `xml:"se_linux"`
	Spm struct {
		Priority string `xml:"priority"`
		Status   string `xml:"status"`
	} `xml:"spm"`
	Ssh struct {
		Fingerprint string `xml:"fingerprint"`
		Port        string `xml:"port"`
	} `xml:"ssh"`
	Status  string `xml:"status"`
	Summary struct {
		Active    int `xml:"active"`
		Migrating int `xml:"migrating"`
		Total     int `xml:"total"`
	} `xml:"summary"`
	TransparentHugepages struct {
		Enabled string `xml:"enabled"`
	} `xml:"transparent_hugepages"`
	Type            string `xml:"type"`
	UpdateAvailable string `xml:"update_available"`
	Version         struct {
		Build       string `xml:"build"`
		FullVersion string `xml:"full_version"`
		Major       int    `xml:"major"`
		Minor       int    `xml:"minor"`
		Revision    string `xml:"revision"`
	} `xml:"version"`
	Cluster struct {
		Href string `xml:"href,attr"`
		ID   string `xml:"id,attr"`
	} `xml:"cluster"`
}

type Cpu struct {
	Name    string
	Speed   string
	Cores   string
	Sockets string
	Threads string
}

type XMLHostCpu struct {
	Name     string `xml:"name"`
	Speed    string `xml:"speed"`
	Topology struct {
		Cores   string `xml:"cores"`
		Sockets string `xml:"sockets"`
		Threads string `xml:"threads"`
	} `xml:"topology"`
}

func CpuFromXml(xCpu XMLHostCpu) Cpu {
	cpu := Cpu{
		Name:    xCpu.Name,
		Speed:   xCpu.Speed,
		Cores:   xCpu.Topology.Cores,
		Sockets: xCpu.Topology.Sockets,
		Threads: xCpu.Topology.Threads,
	}

	return cpu
}

type Hardware struct {
	Manufacturer string `xml:"manufacturer"`
	ProductName  string `xml:"product_name"`
	SerialNumber string `xml:"serial_number"`
	Uuid         string `xml:"uuid"`
	Version      string `xml:"version"`
}

type XMLHostHardware struct {
	Hardware
	Family              string `xml:"family"`
	SupportedRngSources struct {
		Text               string `xml:",chardata"`
		SupportedRngSource string `xml:"supported_rng_source"`
	} `xml:"supported_rng_sources"`
}

type Os struct {
	Type        string
	FullVersion string
}

func OsFromXml(xOs XMLHostOs) Os {
	os := Os{
		Type:        xOs.Type,
		FullVersion: xOs.Version.FullVersion,
	}

	return os
}

type XMLHostOs struct {
	CustomKernelCmdline   string `xml:"custom_kernel_cmdline"`
	ReportedKernelCmdline string `xml:"reported_kernel_cmdline"`
	Type                  string `xml:"type"`
	Version               struct {
		Text        string `xml:",chardata"`
		FullVersion string `xml:"full_version"`
		Major       string `xml:"major"`
	} `xml:"version"`
}
