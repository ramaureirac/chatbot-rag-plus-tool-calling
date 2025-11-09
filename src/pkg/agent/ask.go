package agentclient

import (
	"context"
	"log"
	"time"

	api "github.com/ollama/ollama/api"
)

func (agt *Agent) AskQuestion(question string) (string, error) {
	agt.LastRequest = time.Now()
	agt.appendMessage(api.Message{Role: "user", Content: question})
	var res string
	var err error
	err = agt.ollama.Chat(
		context.Background(),
		agt.chatRequest,
		func(cr api.ChatResponse) error {
			if cr.Done {
				res, err = agt.execute(&cr.Message)
				if err != nil {
					log.Println("exec: " + err.Error())
				}
			}
			return nil
		},
	)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return res, nil
}
