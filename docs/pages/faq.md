# F.A.Q

### How does the server manage sessions?

There isn't a properlu way to handle sessions right now. At this time each time a client performs a POST to `/login` the server internally will create a random UUID and allocate an Agent Instance in a Sessions map (where UUID is the key). There also a GO ROUTINE for dropping old sessions without activity. 

Sources: `src/server/sessions.go`, `src/server/middleware.go`

### How does the agent know what to do?

Currently the main agent is a `qwen3:14b` instance with prompts configured to only attend incoming questions about GitLab repositories. When there’s a match, the agent invokes the appropriate function. If the agent cannot produce a definitive answer, it falls back to a second `model—embeddinggemma:300m`for vector search and then uses `gemma3:4b` to generate the answer.

Sources: `model/devops.modelfile.md`, `model/rag.modelfile.md`, `src/agent/execute.go`, `src/agent/toos.go`

### How can I add new tools?

All tools are available inside `src/internal` folder. You can always create a folder inside with your code. For reference you can look `/src/internal/gitlab/repos.go` source code and the function `CreateRepository()`. Once you have your function declared and know is working fine you MUST introduce the function to the Agent Model: `model/devops.modelfile.md` following the same GitLab Example. Finally once you function is declared and properly introduced to the model you can attach it as a tool in `src/agent/tools.go`.

Sources: `model/devops.modelfile.md`, `agent/tools.go`


