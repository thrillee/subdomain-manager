package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/thrillee/namecheap-dns-manager/internals"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dev_store_fs := os.Getenv("DEV_HOST_STORE")
	dev_hosts, err := internals.GetOrNewHostStore(dev_store_fs)
	if err != nil {
		log.Fatal("Failed To Load Dev Store: %v", err)
	}

	prod_hosts, err := internals.GetOrNewHostStore(dev_store_fs)
	if err != nil {
		log.Fatal("Failed To Load Prod Store: %v", err)
	}
}
