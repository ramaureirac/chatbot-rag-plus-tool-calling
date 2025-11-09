package rag

import (
	"context"
	"strings"

	"github.com/milvus-io/milvus/client/v2/entity"
	"github.com/milvus-io/milvus/client/v2/milvusclient"
)

func (r *RAG) Search(query string) (string, error) {
	qv, err := r.generateEmbeddings(query)
	if err != nil {
		return "", err
	}

	ctx := context.Background()
	eqv := []entity.Vector{entity.FloatVector(qv)}
	outs := []string{"chunk"}
	lim := 3
	opts := milvusclient.NewSearchOption("rag", lim, eqv).
		WithOutputFields(outs...).
		WithANNSField("embedding")

	resultSets, err := r.milvus.Search(ctx, opts)
	if err != nil {
		return "", err
	}

	content := []string{}
	for idx := range lim {
		c, err := resultSets[0].GetColumn("chunk").GetAsString(idx)
		content = append(content, c)
		if err != nil {
			break
		}
	}

	ragctx := strings.Join(content, "; ")

	return ragctx, nil
}
