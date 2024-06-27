package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/thrillee/namecheap-dns-manager/httpsvr"
	"github.com/thrillee/namecheap-dns-manager/internals"
	"github.com/thrillee/namecheap-dns-manager/namecheap"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dev_store_fs := os.Getenv("DEV_HOST_STORE")
	dev_hosts, err := internals.GetOrNewHostStore(dev_store_fs)
	if err != nil {
		log.Fatal(fmt.Printf("Failed To Load Dev Store: %v", err))
	}

	prod_hosts, err := internals.GetOrNewHostStore(dev_store_fs)
	if err != nil {
		log.Fatal(fmt.Printf("Failed To Load Prod Store: %v", err))
	}

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		log.Fatal("PORT is not found in the env")
	}

	devNcManager := namecheap.CreateNameCheapHostManager(dev_hosts, false)
	prodNcManager := namecheap.CreateNameCheapHostManager(prod_hosts, true)

	httpServer := httpsvr.HttpAPIServer{
		ListenAddr:  ":" + port,
		DevService:  devNcManager,
		ProdService: prodNcManager,
	}
	httpServer.Run()
}
