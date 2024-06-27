package internals

import (
	"encoding/json"
	"os"
)

type PersistedHost struct {
	Data map[string]HostData
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

	return &hostData, nil
}

func PersistHosts(file_path string, hostData *PersistedHost) error {
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
		newStore := PersistedHost{Data: map[string]HostData{}}
		err = PersistHosts(file_path, &newStore)
		store = &newStore
	}

	return store, err
}
