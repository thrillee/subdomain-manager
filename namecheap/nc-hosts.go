package namecheap

import (
	"fmt"
	"log"

	"github.com/thrillee/namecheap-dns-manager/internals"
)

type Namecheap struct {
	db         *internals.PersistedHost
	apiManager nameCheapAPI
	managerKey string
}

func CreateNameCheapHostManager(db *internals.PersistedHost, isLive bool, key string) *Namecheap {
	return &Namecheap{db: db, apiManager: *createNameCheapAPI(isLive), managerKey: key}
}

func formDomain(sld, tld string) string {
	return fmt.Sprintf("%s.%s", sld, tld)
}

func (nc *Namecheap) GetFactoryKey() string {
	return nc.managerKey
}

func (nc *Namecheap) AddSubDomain(data internals.HostData) (internals.HostResponse, error) {
	newDomain := formDomain(data.SLD, data.TLD)

	hostRecords, ok := nc.db.Data[newDomain]
	if !ok {
		hostRecords = internals.HostData{
			SLD:     data.SLD,
			TLD:     data.TLD,
			Records: []internals.HostRecord{},
		}
	}

	duplicatedRecords := []string{}
	persistedRecords := hostRecords.Records
	for _, storedRecords := range hostRecords.Records {
		for _, dataRecord := range data.Records {
			if storedRecords.HostName == dataRecord.HostName {
				duplicatedRecords = append(duplicatedRecords, dataRecord.HostName)
			} else {
				persistedRecords = append(persistedRecords, dataRecord)
			}
		}
	}

	if len(duplicatedRecords) > 0 {
		return internals.HostResponse{}, fmt.Errorf("%v Already Exists", duplicatedRecords)
	}

	hostRecords.Records = persistedRecords

	nameCheapApiResponse, err := nc.apiManager.postHost(hostRecords)
	if err != nil {
		return internals.HostResponse{}, err
	}

	success := nameCheapApiResponse.CommandResponse.DomainDNSSetHostsResult.IsSuccess == "true"
	if success {
		nc.db.Data[newDomain] = hostRecords
		log.Printf("%d Host Record Added: %d Total Hosts", len(data.Records), len(hostRecords.Records))
	}

	return internals.HostResponse{
		Success: success,
		Message: nameCheapApiResponse.Status,
	}, nil
}

func (nc *Namecheap) DeleteSubDomain(data internals.HostData) (internals.HostResponse, error) {
	newDomain := formDomain(data.SLD, data.TLD)

	hostRecords, ok := nc.db.Data[newDomain]
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

	nameCheapApiResponse, err := nc.apiManager.postHost(hostRecords)
	if err != nil {
		return internals.HostResponse{}, err
	}

	success := nameCheapApiResponse.CommandResponse.DomainDNSSetHostsResult.IsSuccess == "true"
	if success {
		nc.db.Data[newDomain] = hostRecords
		log.Printf("%d Host Record Deleted: %d Total Hosts", len(data.Records), len(hostRecords.Records))
	}

	return internals.HostResponse{
		Success: success,
		Message: nameCheapApiResponse.Status,
	}, nil
}
