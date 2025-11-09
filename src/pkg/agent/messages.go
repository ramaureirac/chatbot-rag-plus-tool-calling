package agentclient

import "github.com/ollama/ollama/api"

func (agt *Agent) appendMessage(msg api.Message) {
	agt.chatRequest.Messages = append(agt.chatRequest.Messages, msg)
	msgLen := len(agt.chatRequest.Messages)
	if msgLen > agt.chatLimit {
		agt.chatRequest.Messages = agt.chatRequest.Messages[msgLen-agt.chatLimit:]
	}
}
