package hostfactory

import "github.com/thrillee/namecheap-dns-manager/internals"

type AbstractHostFactory interface {
	RegisterNewHostManager(internals.HostManger)
	GetManager(string) (internals.HostManger, error)
}
