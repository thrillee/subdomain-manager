package namecheap

import (
	"fmt"

	"github.com/thrillee/namecheap-dns-manager/internals"
)

type Namecheap struct {
	*internals.PersistedHost
}

func formDomain(sld, tld string) string {
	return fmt.Sprintf("%s.%s", sld, tld)
}

func (nc *Namecheap) AddSubDomain(data internals.HostData) (internals.HostResponse, error) {
	newDomain := formDomain(data.SLD, data.TLD)

	hostRecords, ok := nc.Data[newDomain]
	if !ok {
		hostRecords = internals.HostData{
			SLD:     data.SLD,
			TLD:     data.TLD,
			Records: []internals.HostRecord{},
		}
	}

	duplicatedRecords := []string{}
	for _, storedRecords := range hostRecords.Records {
		for _, dataRecord := range data.Records {
			if storedRecords.HostName == dataRecord.HostName {
				duplicatedRecords = append(duplicatedRecords, dataRecord.HostName)
			}
		}
	}

	if len(duplicatedRecords) > 0 {
		return internals.HostResponse{}, fmt.Errorf("%v Already Exists", duplicatedRecords)
	}

	return internals.HostResponse{
		Success: false,
		Message: "Something Went Wrong",
	}, nil
}

func (nc *Namecheap) DeleteSubDomain(data internals.HostData) (internals.HostResponse, error) {
	newDomain := formDomain(data.SLD, data.TLD)

	hostRecords, ok := nc.Data[newDomain]
	if !ok {
		hostRecords = internals.HostData{
			SLD:     data.SLD,
			TLD:     data.TLD,
			Records: []internals.HostRecord{},
		}
	}

	notFoundRecords := data.Records
	newHostRecords := hostRecords.Records
	for storedRecordIdx, storedRecords := range hostRecords.Records {
		for dataIdx, dataRecord := range data.Records {
			if storedRecords.HostName == dataRecord.HostName {
				notFoundRecords = internals.DeleteHostRecord(notFoundRecords, dataIdx)
				newHostRecords = internals.DeleteHostRecord(newHostRecords, storedRecordIdx)
			}
		}
	}

	if len(notFoundRecords) > 0 {
		return internals.HostResponse{}, fmt.Errorf("%v Not Found", notFoundRecords)
	}

	return internals.HostResponse{
		Success: false,
		Message: "Something Went Wrong",
	}, nil
}
