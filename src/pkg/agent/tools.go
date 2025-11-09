package agentclient

import "github.com/ollama/ollama/api"

func getChatTools() api.Tools {
	tools := api.Tools{
		api.Tool{
			Type: "function",
			Function: api.ToolFunction{
				Name:        "CreateGitLabRepository",
				Description: "Creates a new GitLab repository by given it's group id (int) and name (string)",
				Parameters: api.ToolFunctionParameters{
					Type:     "object",
					Required: []string{"name", "group"},
					Properties: map[string]api.ToolProperty{
						"name": {
							Type:        api.PropertyType{"string"},
							Description: "Repository name",
						},
						"group": {
							Type:        api.PropertyType{"string"},
							Description: "Repository group ID",
						},
					},
				},
			},
		},
		api.Tool{
			Type: "function",
			Function: api.ToolFunction{
				Name:        "InvokeRAG",
				Description: "Invokes RAG Service to retrive a response",
				Parameters: api.ToolFunctionParameters{
					Type:     "object",
					Required: []string{"query"},
					Properties: map[string]api.ToolProperty{
						"query": {
							Type:        api.PropertyType{"string"},
							Description: "Question for vector search in RAG System",
						},
					},
				},
			},
		},
	}
	return tools
}
