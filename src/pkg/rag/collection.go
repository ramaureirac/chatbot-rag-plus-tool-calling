package rag

import (
	"context"
	"os"

	"github.com/milvus-io/milvus/client/v2/entity"
	"github.com/milvus-io/milvus/client/v2/index"
	"github.com/milvus-io/milvus/client/v2/milvusclient"
)

func (r *RAG) createCollection() error {
	schema := entity.Schema{
		CollectionName: "rag",
		Description:    "Embeddings collection",
		Fields: []*entity.Field{
			{
				Name:       "id",
				DataType:   entity.FieldTypeVarChar,
				PrimaryKey: true,
				AutoID:     false,
				TypeParams: map[string]string{
					"max_length": "500",
				},
			},
			{
				Name:     "chunk",
				DataType: entity.FieldTypeVarChar,
				TypeParams: map[string]string{
					"max_length": "65535",
				},
			},
			{
				Name:     "embedding",
				DataType: entity.FieldTypeFloatVector,
				TypeParams: map[string]string{
					"dim": os.Getenv("OLLAMA_EMBED_SIZE"),
				},
			},
		},
	}
	ctx := context.Background()
	opts := milvusclient.NewCreateCollectionOption("rag", &schema)
	err := r.milvus.CreateCollection(ctx, opts)
	if err != nil {
		return err
	}

	idx := index.NewHNSWIndex(entity.COSINE, 32, 128)
	idxopts := milvusclient.NewCreateIndexOption("rag", "embedding", idx)
	_, err = r.milvus.CreateIndex(ctx, idxopts)
	if err != nil {
		return err
	}

	err = r.loadCollection()
	if err != nil {
		return err
	}

	return nil
}

func (r *RAG) loadCollection() error {
	ctx := context.Background()
	_, err := r.milvus.LoadCollection(ctx, milvusclient.NewLoadCollectionOption("rag"))
	if err != nil {
		return err
	}
	return nil
}
