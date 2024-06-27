package namecheap

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/thrillee/namecheap-dns-manager/internals"
)

type nameCheapApiResponse struct {
	XMLName         xml.Name        `xml:"ApiResponse"`
	Status          string          `xml:"Status,attr"`
	CommandResponse commandResponse `xml:"CommandResponse"`
}

type commandResponse struct {
	DomainDNSSetHostsResult domainDNSSetHostsResult `xml:"DomainDNSSetHostsResult"`
}

type host struct {
	HostId             string `xml:"HostId,attr"`
	Name               string `xml:"Name,attr"`
	Type               string `xml:"Type,attr"`
	Address            string `xml:"Address,attr"`
	MXPref             string `xml:"MXPref,attr"`
	TTL                string `xml:"TTL,attr"`
	AssociatedAppTitle string `xml:"AssociatedAppTitle,attr"`
	FriendlyName       string `xml:"FriendlyName,attr"`
	IsActive           string `xml:"IsActive,attr"`
	IsDDNSEnabled      string `xml:"IsDDNSEnabled,attr"`
}

type domainDNSSetHostsResult struct {
	Domain        string `xml:"Domain,attr"`
	IsSuccess     string `xml:"IsSuccess,attr"`
	EmailType     string `xml:"EmailType,attr"`
	IsUsingOurDNS string `xml:"IsUsingOurDNS,attr"`
	Hosts         []host `xml:"host"`
}

type nameCheapAPI struct {
	isLive     bool
	properties *Properties
}

func (nc *nameCheapAPI) getUrl() string {
	if nc.isLive {
		return nc.properties.prod_url
	}
	return nc.properties.dev_url
}

func (nc *nameCheapAPI) getAPIKey() string {
	if nc.isLive {
		return nc.properties.PROD_NAMECHEAP_API_KEY
	}
	return nc.properties.DEV_NAMECHEAP_API_KEY
}

func (nc *nameCheapAPI) getAPIUsername() string {
	if nc.isLive {
		return nc.properties.PROD_NAMECHEAP_API_USERNAME
	}
	return nc.properties.DEV_NAMECHEAP_API_USERNAME
}

func (nc *nameCheapAPI) postHost(data internals.HostData) (nameCheapApiResponse, error) {
	baseURL := nc.getUrl()

	params := url.Values{}
	params.Add("SLD", data.SLD)
	params.Add("TLD", data.TLD)
	params.Add("apikey", nc.getAPIKey())
	params.Add("apiuser", nc.getAPIUsername())
	params.Add("username", nc.getAPIUsername())
	params.Add("ClientIp", nc.properties.HOST_IP)
	params.Add("Command", "namecheap.domains.dns.setHosts")

	for idx, record := range data.Records {
		params.Add(fmt.Sprintf("TTL%d", idx+1), "1800")
		params.Add(fmt.Sprintf("RecordType%d", idx+1), "A")
		params.Add(fmt.Sprintf("Address%d", idx+1), record.Address)
		params.Add(fmt.Sprintf("HostName%d", idx+1), record.HostName)
	}

	urlWithParams := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	fmt.Println(urlWithParams)

	req, err := http.NewRequest("GET", urlWithParams, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return nameCheapApiResponse{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return nameCheapApiResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	log.Printf("Response Body: \n %v \n", body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return nameCheapApiResponse{}, err
	}

	var response nameCheapApiResponse
	err = xml.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error unmarshalling XML: ", err)
		return nameCheapApiResponse{}, err
	}
	log.Printf("Response Body: \n %v \n", response)

	return response, err
}

func (nc *nameCheapAPI) listHost(data internals.HostData) (nameCheapApiResponse, error) {
	baseURL := nc.getUrl()

	params := url.Values{}
	params.Add("SLD", data.SLD)
	params.Add("TLD", data.TLD)
	params.Add("apikey", nc.getAPIKey())
	params.Add("apiuser", nc.getAPIUsername())
	params.Add("username", nc.getAPIUsername())
	params.Add("ClientIp", nc.properties.HOST_IP)
	params.Add("Command", "namecheap.domains.dns.setHosts")

	urlWithParams := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	fmt.Println(urlWithParams)

	req, err := http.NewRequest("GET", urlWithParams, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return nameCheapApiResponse{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return nameCheapApiResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return nameCheapApiResponse{}, err
	}

	var response nameCheapApiResponse
	err = xml.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error unmarshalling XML: ", err)
		return nameCheapApiResponse{}, err
	}

	return response, err
}

func createNameCheapAPI(isLive bool) *nameCheapAPI {
	return &nameCheapAPI{
		properties: createProperties(),
		isLive:     isLive,
	}
}
