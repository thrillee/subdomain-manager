package internals

import (
	"encoding/json"
	"os"
)

type PersistedHost struct {
	Data     map[string]HostData
	Filepath string
}

func (p *PersistedHost) Save() {
	persistHosts(p.Filepath, p)
}

func GetPersistedHosts(file_path string) (*PersistedHost, error) {
	data, err := os.ReadFile(file_path)
	if err != nil {
		return nil, err
	}

	var hostData PersistedHost
	err = json.Unmarshal(data, &hostData)
	if err != nil {
		return nil, err
	}

	hostData.Filepath = file_path

	return &hostData, nil
}

func persistHosts(file_path string, hostData *PersistedHost) error {
	data, err := json.Marshal(hostData)
	if err != nil {
		return err
	}

	err = os.WriteFile(file_path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func GetOrNewHostStore(file_path string) (*PersistedHost, error) {
	store, err := GetPersistedHosts(file_path)
	if err != nil {
		newStore := PersistedHost{Data: map[string]HostData{}, Filepath: file_path}
		err = persistHosts(file_path, &newStore)
		store = &newStore
	}

	return store, err
}
