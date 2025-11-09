package rag

import (
	"log"
	"os"
	"strconv"

	"github.com/milvus-io/milvus/client/v2/milvusclient"
	"github.com/ollama/ollama/api"
	ollamaclient "github.com/ramaureirac/devops-ragbot/src/pkg/ollama"
)

type RAG struct {
	ollama    *api.Client
	milvus    *milvusclient.Client
	vectorDim int
}

func NewRag() (*RAG, error) {
	dim, err := strconv.Atoi(os.Getenv("OLLAMA_EMBED_SIZE"))
	if err != nil {
		log.Fatalln("rag: " + err.Error())
	}
	ol, err := ollamaclient.NewOllamaApiClient()
	if err != nil {
		log.Fatalln("ollama: " + err.Error())
	}

	db, err := newDataBaseClient()
	if err != nil {
		return &RAG{}, err
	}
	r := RAG{
		ollama:    ol,
		milvus:    db,
		vectorDim: dim,
	}
	return &r, nil
}
