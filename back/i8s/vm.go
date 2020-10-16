package i8s

func VmFromXml(xVm XMLVm, hosts map[string]*Host) Vm {
	h := hosts[xVm.Host.ID]
	return Vm{
		ID:               xVm.ID,
		Name:             xVm.Name,
		Host:             h,
		Cpu:              VCpuFromXml(xVm.Cpu),
		CpuShares:        xVm.CpuShares,
		CreationTime:     xVm.CreationTime,
		Display:          xVm.Display.Display,
		HighAvailability: xVm.HighAvailability,
		Memory:           xVm.Memory,
		MemoryPolicy:     xVm.MemoryPolicy,
		Os:               xVm.Os.Type,
		Type:             xVm.Type,
		Status:           xVm.Status,
		StopTime:         xVm.StopTime,
		StopReason:       xVm.StopReason,
		OperatingSystem:  xVm.GuestOperatingSystem.XMLOperatingSystem,
		StartTime:        xVm.StartTime,
	}
}

type Vm struct {
	ID               string
	Name             string
	Host             *Host
	Cpu              VCpu
	CpuShares        int
	CreationTime     string
	Display          Display
	HighAvailability XMLHighAvailability
	Memory           int64
	MemoryPolicy     XMLMemoryPolicy
	Os               string
	Type             string
	Status           string
	StopTime         string
	StopReason       string
	OperatingSystem  XMLOperatingSystem
	StartTime        string
}

type VCpu struct {
	Name    string
	Speed   string
	Cores   int
	Sockets int
	Threads int
}

type XMLVCpu struct {
	Architecture string `xml:"architecture"`
	Topology     struct {
		Cores   int `xml:"cores"`
		Sockets int `xml:"sockets"`
		Threads int `xml:"threads"`
	} `xml:"topology"`
	Mode string `xml:"mode"`
}

func VCpuFromXml(xCpu XMLVCpu) VCpu {
	vcpu := VCpu{
		Name:    xCpu.Architecture,
		Cores:   xCpu.Topology.Cores,
		Sockets: xCpu.Topology.Sockets,
		Threads: xCpu.Topology.Threads,
	}

	return vcpu
}

type XMLVm struct {
	ID          string `xml:"id,attr"` //="1d312aac-c09a-4231-8737-249cef1bae7d"
	Name        string `xml:"name"`
	Description string `xml:"description"`
	Comment     string `xml:"comment"`
	Bios        struct {
		Text     string `xml:",chardata"`
		BootMenu struct {
			Enabled string `xml:"enabled"`
		} `xml:"boot_menu"`
	} `xml:"bios"`
	Cpu              XMLVCpu             `xml:"cpu"`
	CpuShares        int                 `xml:"cpu_shares"`
	CreationTime     string              `xml:"creation_time"`
	DeleteProtected  string              `xml:"delete_protected"`
	Display          XMLDisplay          `xml:"display"`
	HighAvailability XMLHighAvailability `xml:"high_availability"`
	Io               struct {
		Threads int `xml:"threads"`
	} `xml:"io"`
	Memory       int64           `xml:"memory"`
	MemoryPolicy XMLMemoryPolicy `xml:"memory_policy"`
	Migration    struct {
		AutoConverge string `xml:"auto_converge"`
		Compressed   string `xml:"compressed"`
	} `xml:"migration"`
	MigrationDowntime string `xml:"migration_downtime"`
	Origin            string `xml:"origin"`
	Os                struct {
		Boot struct {
			Devices struct {
				Device []string `xml:"device"`
			} `xml:"devices"`
		} `xml:"boot"`
		Type string `xml:"type"`
	} `xml:"os"`
	PlacementPolicy struct {
		Affinity string `xml:"affinity"`
		Hosts    struct {
			Text string `xml:",chardata"`
			Host struct {
				Href string `xml:"href,attr"`
				ID   string `xml:"id,attr"`
			} `xml:"host"`
		} `xml:"hosts"`
	} `xml:"placement_policy"`
	SmallIcon struct {
		Href string `xml:"href,attr"`
		ID   string `xml:"id,attr"`
	} `xml:"small_icon"`
	Sso struct {
		Methods struct {
			Method struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
			} `xml:"method"`
		} `xml:"methods"`
	} `xml:"sso"`
	StartPaused                 string `xml:"start_paused"`
	Stateless                   string `xml:"stateless"`
	StorageErrorResumeBehaviour string `xml:"storage_error_resume_behaviour"`
	TimeZone                    struct {
		Name string `xml:"name"`
	} `xml:"time_zone"`
	Type string `xml:"type"`
	Usb  struct {
		Enabled string `xml:"enabled"`
	} `xml:"usb"`
	Cluster struct {
		Href string `xml:"href,attr"`
		ID   string `xml:"id,attr"`
	} `xml:"cluster"`
	CpuProfile struct {
		Href string `xml:"href,attr"`
		ID   string `xml:"id,attr"`
	} `xml:"cpu_profile"`
	Quota struct {
		ID string `xml:"id,attr"`
	} `xml:"quota"`
	NextRunConfigurationExists string `xml:"next_run_configuration_exists"`
	NumaTuneMode               string `xml:"numa_tune_mode"`
	Status                     string `xml:"status"`
	StopTime                   string `xml:"stop_time"`
	OriginalTemplate           struct {
		Href string `xml:"href,attr"`
		ID   string `xml:"id,attr"`
	} `xml:"original_template"`
	Template struct {
		Href string `xml:"href,attr"`
		ID   string `xml:"id,attr"`
	} `xml:"template"`
	StopReason     string `xml:"stop_reason"`
	Initialization struct {
		AuthorizedSshKeys string `xml:"authorized_ssh_keys"`
		CustomScript      string `xml:"custom_script"`
		NicConfigurations string `xml:"nic_configurations"`
		RegenerateSshKeys string `xml:"regenerate_ssh_keys"`
		UserName          string `xml:"user_name"`
		HostName          string `xml:"host_name"`
		ActiveDirectoryOu string `xml:"active_directory_ou"`
		Domain            string `xml:"domain"`
		InputLocale       string `xml:"input_locale"`
		OrgName           string `xml:"org_name"`
		SystemLocale      string `xml:"system_locale"`
		UiLanguage        string `xml:"ui_language"`
		UserLocale        string `xml:"user_locale"`
	} `xml:"initialization"`
	Fqdn                 string                  `xml:"fqdn"`
	GuestOperatingSystem XMLGuestOperatingSystem `xml:"guest_operating_system"`
	GuestTimeZone        struct {
		Name      string `xml:"name"`
		UtcOffset string `xml:"utc_offset"`
	} `xml:"guest_time_zone"`
	RunOnce   string `xml:"run_once"`
	StartTime string `xml:"start_time"`
	Host      struct {
		Href string `xml:"href,attr"`
		ID   string `xml:"id,attr"`
	} `xml:"host"`
	CustomCompatibilityVersion struct {
		Major string `xml:"major"`
		Minor string `xml:"minor"`
	} `xml:"custom_compatibility_version"`
	Domain struct {
		Name string `xml:"name"`
	} `xml:"domain"`
}

type XMLOperatingSystem struct {
	Architecture string `xml:"architecture"`
	Distribution string `xml:"distribution"`
	Family       string `xml:"family"`
	Codename     string `xml:"codename"`
}

type XMLGuestOperatingSystem struct {
	XMLOperatingSystem
	Kernel struct {
		Version struct {
			Build       string `xml:"build"`
			FullVersion string `xml:"full_version"`
			Major       string `xml:"major"`
			Minor       string `xml:"minor"`
			Revision    string `xml:"revision"`
		} `xml:"version"`
	} `xml:"kernel"`
	Version struct {
		Build       string `xml:"build"`
		FullVersion string `xml:"full_version"`
		Major       string `xml:"major"`
		Minor       string `xml:"minor"`
	} `xml:"version"`
}

type XMLMemoryPolicy struct {
	Guaranteed int64 `xml:"guaranteed"`
	Max        int64 `xml:"max"`
}

type XMLHighAvailability struct {
	Enabled  bool `xml:"enabled"`
	Priority int  `xml:"priority"`
}

type Display struct {
	Type       string `xml:"type"`
	Address    string `xml:"address"`
	SecurePort int    `xml:"secure_port"`
	Port       int    `xml:"port"`
}

type XMLDisplay struct {
	Display
	AllowOverride       string `xml:"allow_override"`
	CopyPasteEnabled    string `xml:"copy_paste_enabled"`
	DisconnectAction    string `xml:"disconnect_action"`
	FileTransferEnabled string `xml:"file_transfer_enabled"`
	Monitors            string `xml:"monitors"`
	SingleQxlPci        string `xml:"single_qxl_pci"`
	SmartcardEnabled    string `xml:"smartcard_enabled"`
}

type XMLVms struct {
	Vms []XMLVm `xml:"vm"`
}
