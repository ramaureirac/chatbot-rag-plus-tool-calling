package rag

import (
	"context"

	"github.com/milvus-io/milvus/client/v2/milvusclient"
)

func (r *RAG) insertEmbeddings(doc *EmbeddingDocument) error {
	ctx := context.Background()
	inopts := milvusclient.NewColumnBasedInsertOption("rag").
		WithVarcharColumn("id", []string{doc.Id}).
		WithVarcharColumn("chunk", []string{doc.Chunk}).
		WithFloatVectorColumn("embedding", r.vectorDim, [][]float32{doc.Embedding})
	_, err := r.milvus.Insert(ctx, inopts)
	if err != nil {
		return err
	}
	r.loadCollection()
	return nil
}
