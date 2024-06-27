package internals

type HostRecord struct {
	TTL        int
	HostName   string
	RecordType string
	Address    string
}

type HostData struct {
	SLD     string
	TLD     string
	Records []HostRecord
}

type HostResponse struct {
	Success bool   `json:"success"`
	Domain  string `json:"domain"`
	Message string `json:"message"`
}

type Host struct {
	HostId             string
	Name               string
	Type               string
	Address            string
	MXPref             string
	TTL                string
	AssociatedAppTitle string
	FriendlyName       string
	IsActive           string
	IsDDNSEnabled      string
}

type HostManger interface {
	AddSubDomain(HostData) (HostResponse, error)
	DeleteSubDomain(HostData) (HostResponse, error)
	GetHosts() []Host
	*PersistedHost
}

func DeleteHostRecord(slice []HostRecord, index int) []HostRecord {
	return append(slice[:index], slice[index+1:]...)
}
