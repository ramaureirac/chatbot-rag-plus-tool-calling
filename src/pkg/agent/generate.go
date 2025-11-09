package agentclient

import (
	"context"
	"log"
	"os"

	api "github.com/ollama/ollama/api"
)

func (agt *Agent) generate(prompt string) string {
	stream := false
	var res string
	err := agt.ollama.Generate(context.Background(), &api.GenerateRequest{
		Model:  os.Getenv("OLLAMA_GEN_MODEL"),
		Prompt: prompt,
		Stream: &stream,
	}, func(gr api.GenerateResponse) error {
		if gr.Done {
			res = gr.Response
		}
		return nil
	})
	if err != nil {
		log.Panicln("gen: " + err.Error())
	}
	return res
}

func (agt *Agent) generateErroMessage(err string) string {
	return agt.generate("YOU got this error: " + err + " when YOU trigger an API. write a short notification")
}

func (agt *Agent) generateAnswer(info string, query string) string {
	return agt.generate("Using this context: " + info + " answer the question: " + query)
}
