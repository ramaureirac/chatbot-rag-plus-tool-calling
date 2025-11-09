package main

import (
	"fmt"
	"log"
	"os"

	godotenv "github.com/joho/godotenv"
	agent "github.com/ramaureirac/devops-ragbot/src/pkg/agent"
	rag "github.com/ramaureirac/devops-ragbot/src/pkg/rag"
	router "github.com/ramaureirac/devops-ragbot/src/server"
)

func main() {

	godotenv.Load(".env")

	argsLen := len(os.Args)
	if argsLen > 1 {
		switch os.Args[1] {

		case "serve":
			log.Println("Now running in server mode")
			r := router.NewRouterApp()
			p := os.Getenv("GIN_PORT")
			log.Println("Listening on http://localhost:" + p)
			r.Run("0.0.0.0:" + p)

		case "embed":
			if argsLen > 2 {
				log.Println("Now populating RAG")
				rag.LoadSources(os.Args[2])
			} else {
				fmt.Println("error: missing directory")
			}
		}
	} else {
		log.Println("Now running in local mode")
		agent.NewAgentApp()
	}

}
