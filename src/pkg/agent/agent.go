package agentclient

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	api "github.com/ollama/ollama/api"
	gitlabclient "github.com/ramaureirac/devops-ragbot/src/pkg/gitlab"
	ollamaclient "github.com/ramaureirac/devops-ragbot/src/pkg/ollama"
	rag "github.com/ramaureirac/devops-ragbot/src/pkg/rag"
)

type Agent struct {
	gitLab      *gitlabclient.GitLab
	ollama      *api.Client
	raginstance *rag.RAG
	chatRequest *api.ChatRequest
	chatLimit   int
	LastRequest time.Time
}

func NewAgent() (*Agent, error) {
	gl, err := gitlabclient.NewGitLab()
	if err != nil {
		log.Fatalln("gitlab error: " + err.Error())
	}
	ol, err := ollamaclient.NewOllamaApiClient()
	if err != nil {
		log.Fatalln("ollama error: " + err.Error())
	}
	r, err := rag.NewRag()
	if err != nil {
		log.Fatal("rag: unable to instanciate rag " + err.Error())
	}
	think := false
	chat := &api.ChatRequest{
		Model:    os.Getenv("OLLAMA_CHAT_MODEL"),
		Stream:   &think,
		Think:    &api.ThinkValue{Value: false},
		Messages: []api.Message{},
		Tools:    getChatTools(),
	}
	dt := time.Now()
	agt := &Agent{
		gitLab:      gl,
		ollama:      ol,
		LastRequest: dt,
		raginstance: r,
		chatRequest: chat,
		chatLimit:   15,
	}
	return agt, nil
}

func NewAgentApp() {
	agent, err := NewAgent()
	reader := bufio.NewReader(os.Stdin)

	if err != nil {
		log.Fatalln("agent: " + err.Error())
	}

	for {
		fmt.Print("user $> ")
		query, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln("input: " + err.Error())
		}
		res, err := agent.AskQuestion(query)
		if err != nil {
			log.Fatalln("agent: " + err.Error())
		}
		fmt.Println("agnt $> " + res)
	}
}
