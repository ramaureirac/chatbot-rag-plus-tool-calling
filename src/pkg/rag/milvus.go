package rag

// docs: https://milvus.io/docs/es/insert-update-delete.md

import (
	"context"
	"os"

	"github.com/milvus-io/milvus/client/v2/milvusclient"
)

type DataBase struct {
	Client *milvusclient.Client
}

func newDataBaseClient() (*milvusclient.Client, error) {
	uri := os.Getenv("MILVUS_URI")

	cfg := milvusclient.ClientConfig{
		Address: uri,
	}

	ctx := context.Background()
	mc, err := milvusclient.New(ctx, &cfg)

	return mc, err
}

// func (r *RAG) ListEmbeddings() error {
// 	outs := []string{"id"}
// 	ctx := context.Background()

// 	opt := milvusclient.NewQueryOption("rag").
// 		WithFilter("").
// 		WithLimit(10).
// 		WithOutputFields(outs...)

// 	rs, err := r.Milvus.Query(ctx, opt)
// 	if err != nil {
// 		return err
// 	}

// 	for _, field := range outs {
// 		col := rs.GetColumn(field)
// 		log.Printf("Field %s: %v\n", field, col.FieldData())
// 	}

// 	return nil
// }
