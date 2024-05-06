package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"net/http"
	"net/url"
)

type SynClient struct {
	BaseUrl     string
	Client      http.Client
	AccessToken string
}

func rest_get(client SynClient, endpoint string) (map[string]interface{}, error) {
	base_url, err := url.Parse(client.BaseUrl)
	if err != nil {
		fmt.Println("Malformed URL:", err)
		return nil, err
	}
	full_url := base_url.ResolveReference(&url.URL{Path: endpoint})
	fmt.Println("Full URL: " + full_url.String())

	var bearer = "Bearer " + client.AccessToken

	req, _ := http.NewRequest("GET", full_url.String(), nil)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	resp, err := client.Client.Do(req)
	if err != nil {
		fmt.Printf("Error %s", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	m := make(map[string]interface{})
	foo := json.Unmarshal(body, &m)
	fmt.Println(m)
	if foo != nil {
		log.Fatal(err)
	}
	// fmt.Printf("Body : %s", m)
	return m, nil // Return the map and nil error
}

func main() {
	// The base URL has to end with a trailing slash
	client := SynClient{
		BaseUrl:     "https://repo-prod.prod.sagebase.org/repo/v1/",
		Client:      http.Client{Timeout: time.Duration(1) * time.Second},
		AccessToken: os.Getenv("TOKEN"),
	}
	// result, _ := rest_get(client, "entity/syn53013218")
	result, _ := rest_get(client, "entity/syn4990358")

	fmt.Println(result["name"])
}
