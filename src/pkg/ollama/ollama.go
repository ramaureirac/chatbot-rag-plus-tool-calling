package ollamaclient

import (
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/ollama/ollama/api"
)

func NewOllamaApiClient() (*api.Client, error) {
	apiurl := os.Getenv("OLLAMA_URL")
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	url, err := url.ParseRequestURI(apiurl)
	if err != nil {
		return nil, err
	}
	apiclient := api.NewClient(url, &client)
	return apiclient, nil
}
