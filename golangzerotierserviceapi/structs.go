package golangzerotierserviceapi

type Join struct {
	AllowDNS     bool `json:"allowDNS"`
	AllowDefault bool `json:"allowDefault"`
	AllowManaged bool `json:"allowManaged"`
	AllowGlobal  bool `json:"allowGlobal"`
}

type LeaveNetwork struct {
	Result bool `json:"result"`
}

type Network struct {
	AllowDNS          bool     `json:"allowDNS"`
	AllowDefault      bool     `json:"allowDefault"`
	AllowManaged      bool     `json:"allowManaged"`
	AllowGlobal       bool     `json:"allowGlobal"`
	AssignedAddresses []string `json:"assignedAddresses"`
	Bridge            bool     `json:"bridge"`
	BroadcastEnabled  bool     `json:"broadcastEnabled"`
	DNS               struct {
		Domain  string   `json:"domain"`
		Servers []string `json:"servers"`
	} `json:"dns"`
	ID                     string `json:"id"`
	Mac                    string `json:"mac"`
	Mtu                    int    `json:"mtu"`
	MulticastSubscriptions []struct {
		Adi int    `json:"adi"`
		Mac string `json:"mac"`
	} `json:"multicastSubscriptions"`
	AuthenticationURL        string `json:"authenticationURL"`
	AuthenticationExpiryTime int    `json:"authenticationExpiryTime"`
	Name                     string `json:"name"`
	NetconfRevision          int    `json:"netconfRevision"`
	PortDeviceName           string `json:"portDeviceName"`
	PortError                int    `json:"portError"`
	Routes                   []struct {
		Flags  int    `json:"flags"`
		Metric int    `json:"metric"`
		Target string `json:"target"`
		Via    string `json:"via"`
	} `json:"routes"`
	Status    string      `json:"status"`
	Type      string      `json:"type"`
	Property1 interface{} `json:"property1"`
	Property2 interface{} `json:"property2"`
}

type Status struct {
	Address string `json:"address"`
	Clock   int64  `json:"clock"`
	Config  struct {
		Settings struct {
			AllowTCPFallbackRelay bool     `json:"allowTcpFallbackRelay"`
			ForceTCPRelay         bool     `json:"forceTcpRelay"`
			ListeningOn           []string `json:"listeningOn"`
			PortMappingEnabled    bool     `json:"portMappingEnabled"`
			PrimaryPort           int      `json:"primaryPort"`
			SecondaryPort         int      `json:"secondaryPort"`
			SoftwareUpdate        string   `json:"softwareUpdate"`
			SoftwareUpdateChannel string   `json:"softwareUpdateChannel"`
			SurfaceAddresses      []string `json:"surfaceAddresses"`
			TertiaryPort          int      `json:"tertiaryPort"`
		} `json:"settings"`
	} `json:"config"`
	Online               bool   `json:"online"`
	PlanetWorldID        int    `json:"planetWorldId"`
	PlanetWorldTimestamp int64  `json:"planetWorldTimestamp"`
	PublicIdentity       string `json:"publicIdentity"`
	TCPFallbackActive    bool   `json:"tcpFallbackActive"`
	Version              string `json:"version"`
	VersionBuild         int    `json:"versionBuild"`
	VersionMajor         int    `json:"versionMajor"`
	VersionMinor         int    `json:"versionMinor"`
	VersionRev           int    `json:"versionRev"`
}
