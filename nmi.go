package nmi

import (
	"net/http"
	"os"
)

type NMIClient struct {
	ApiKey     string
	BaseURL    string
	HttpClient *http.Client
}

func NewNMIClient() *NMIClient {
	apiKey := os.Getenv("NMI_API_KEY")
	baseURL := "https://api.nmi.com/v4"

	return &NMIClient{
		ApiKey:     apiKey,
		BaseURL:    baseURL,
		HttpClient: &http.Client{},
	}
}
