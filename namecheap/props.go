package namecheap

import "os"

type Properties struct {
	DEV_NAMECHEAP_API_KEY       string
	DEV_NAMECHEAP_API_USERNAME  string
	PROD_NAMECHEAP_API_KEY      string
	PROD_NAMECHEAP_API_USERNAME string
	HOST_IP                     string
	prod_url                    string
	dev_url                     string
}

func CreateProperties() *Properties {
	props := Properties{
		HOST_IP:                     os.Getenv("HOST_IP"),
		DEV_NAMECHEAP_API_KEY:       os.Getenv("DEV_NAMECHEAP_API_KEY"),
		DEV_NAMECHEAP_API_USERNAME:  os.Getenv("DEV_NAMECHEAP_API_USERNAME"),
		PROD_NAMECHEAP_API_KEY:      os.Getenv("PROD_NAMECHEAP_API_KEY"),
		PROD_NAMECHEAP_API_USERNAME: os.Getenv("PROD_NAMECHEAP_API_USERNAME"),
		prod_url:                    "https://api.namecheap.com/xml.response",
		dev_url:                     "https://api.sandbox.namecheap.com/xml.response",
	}

	return &props
}
