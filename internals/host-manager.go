package internals

type HostRecord struct {
	TTL        int    `json:"ttl"`
	HostName   string `json:"host_name"`
	RecordType string `json:"record_type"`
	Address    string `json:"address"`
}

type HostData struct {
	SLD     string       `json:"sld"`
	TLD     string       `json:"tld"`
	Records []HostRecord `json:"records"`
}

func (h *HostData) RecordExists(record HostRecord) int {
	result := -1
	for idx, r := range h.Records {
		if r.HostName == record.HostName {
			return idx
		}
	}
	return result
}

func (h *HostData) DeleteRecord(record HostRecord) {
	for idx, r := range h.Records {
		if r.HostName == record.HostName {
			h.Records = append(h.Records[:idx], h.Records[idx+1:]...)
			break
		}
	}
}

type HostResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  interface{}
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
	AddSubDomain(*HostData) (HostResponse, error)
	DeleteSubDomain(*HostData) (HostResponse, error)
	ListSubDomain(*HostData) (HostResponse, error)
	GetFactoryKey() string
}

func DeleteHostRecord(slice []HostRecord, index int) []HostRecord {
	return append(slice[:index], slice[index+1:]...)
}
