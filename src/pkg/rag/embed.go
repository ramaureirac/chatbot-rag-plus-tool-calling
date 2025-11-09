package rag

import (
	"context"
	"os"

	"fmt"
	"path/filepath"
	"strings"

	api "github.com/ollama/ollama/api"

	pdf "github.com/ledongthuc/pdf"
)

type EmbeddingDocument struct {
	Id        string
	Chunk     string
	Embedding []float32
}

// type Document struct {
// 	Name       string
// 	Content    string
// 	Embeddings []EmbeddingDocument
// }

func (r *RAG) embedPDF(fp string) ([]EmbeddingDocument, error) {
	file, reader, err := pdf.Open(fp)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var fileContent strings.Builder
	pages := reader.NumPage()

	for idx := 1; idx <= pages; idx++ {
		pageContent := reader.Page(idx)
		if pageContent.V.IsNull() {
			continue
		}
		text, err := pageContent.GetPlainText(nil)
		if err != nil {
			continue
		}
		fileContent.WriteString(text)
	}

	const chunkSize = 150
	var words []string = strings.Fields(fileContent.String())
	var embeddings []EmbeddingDocument

	for idx := 0; idx < len(words); idx += chunkSize {
		end := idx + chunkSize
		if end > len(words) {
			end = len(words)
		}
		chunk := strings.Join(words[idx:end], " ")
		if len(strings.TrimSpace(chunk)) > 0 {
			embedding, err := r.generateEmbeddings(chunk)
			if err != nil {
				return nil, err
			}

			filename := strings.ReplaceAll(filepath.Base(fp), ".", "_")
			embDoc := EmbeddingDocument{
				Id:        fmt.Sprintf("%s_%d_%d", filename, idx, end),
				Chunk:     chunk,
				Embedding: embedding,
			}
			embeddings = append(embeddings, embDoc)
		}
	}

	return embeddings, nil
}

func (r *RAG) generateEmbeddings(text string) ([]float32, error) {
	res, error := r.ollama.Embed(context.Background(), &api.EmbedRequest{
		Model: os.Getenv("OLLAMA_EMBED_MODEL"),
		Input: text,
	})
	return res.Embeddings[0], error
}
