package httpsvr

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/thrillee/namecheap-dns-manager/internals"
)

type HttpAPIServer struct {
	ListenAddr  string
	ProdService internals.HostManger
	DevService  internals.HostManger
}

func (h HttpAPIServer) getHostManagerService(isLive bool) internals.HostManger {
	if isLive {
		return h.ProdService
	}
	return h.DevService
}

func (h HttpAPIServer) Run() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Use(cors.Handler(
		cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"*"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		},
	))

	router := chi.NewRouter()
	router.Post("/add-sd", middlewareWhitelistedIP(h.addSubDomain))
	router.Post("/delete-sd", middlewareWhitelistedIP(h.addSubDomain))

	r.Mount("/api/{env}", router)

	srv := &http.Server{
		Handler: r,
		Addr:    h.ListenAddr,
	}
	fmt.Printf("Starting HTTP Server ON %s...", h.ListenAddr)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
