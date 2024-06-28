package hostfactory

import (
	"fmt"
	"log"

	"github.com/thrillee/namecheap-dns-manager/internals"
)

type HostFactory struct {
	handlers map[string]internals.HostManger
}

func (h HostFactory) RegisterNewHostManager(hm internals.HostManger) {
	h.handlers[hm.GetFactoryKey()] = hm
	log.Printf("Registered Host Manager %v", hm)
}

func (h HostFactory) GetManager(managerKey string) (internals.HostManger, error) {
	hm, ok := h.handlers[managerKey]
	if !ok {
		return nil, fmt.Errorf("Operation Failed: %v not found", managerKey)
	}
	return hm, nil
}

func CreateNewHostFactory() *HostFactory {
	return &HostFactory{
		handlers: map[string]internals.HostManger{},
	}
}
