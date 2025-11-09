package agentclient

import (
	"errors"
	"log"
	"strings"

	api "github.com/ollama/ollama/api"
)

func (agt *Agent) execute(cr *api.Message) (string, error) {
	var err error = nil
	var res string = cr.Content
	if len(cr.ToolCalls) > 0 {
		switch cr.ToolCalls[0].Function.Name {
		case "CreateGitLabRepository":
			log.Println("calling gitlab service")
			err = agt.gitLab.CreateRepository(cr.ToolCalls[0].Function.Arguments)
			if err != nil {
				res = agt.generateErroMessage(err.Error())
				log.Println("gitlab repo: " + err.Error())
			} else {
				res = "¡Repositorio creado!"
				log.Println("gitlab repo: " + cr.ToolCalls[0].Function.Arguments.String())
			}
		case "InvokeRAG":
			log.Println("calling rag service")
			query, ok := cr.ToolCalls[0].Function.Arguments["query"].(string)
			if !ok {
				errtxt := "missing args: query (string) from user"
				return agt.generateErroMessage(errtxt), errors.New(errtxt)
			}
			info, err := agt.raginstance.Search(query)
			if err != nil {
				log.Println("rag system: " + err.Error())
				res = agt.generateErroMessage("error when connecting with rag service databse")
			} else {
				log.Println("rag system: OK")
				res = agt.generateAnswer(info, query)
			}
		default:
			res = "Ha habido un error al procesar lo solicitado"
			log.Println("exec: wrong tool calling")
		}
		agt.chatRequest.Messages = []api.Message{}
	} else {
		agt.appendMessage(*cr)
	}
	return strings.TrimSuffix(res, "\n"), err
}
