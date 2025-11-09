GO_BINARY_FOLDER = bin
GO_BINARY_NAME   = devops-ragbot
GO_SOURCES       = src/cmd
GO_VULN			 = ~/go/bin/govulncheck

VUE_SOURCES      = src/public/vue
VUE_DIST         = src/public/dist

DOCKER_TAG		 = ramaureirac/devops-ragbot
DOCKER_FILE		 = docker/Dockerfile
DOCKER_COMPOSE	 = docker/Compose.yml
DOCKER_VOLUMES	 = docker/volumes

MILVUS_COMPOSE	 = docker/Milvus.Compose.yml


## dev mode run local mode
run: clean ollama
	go run ./$(GO_SOURCES)


## build vue frontend
dist:
	@mkdir -p $(VUE_DIST)
	@rm -rf $(VUE_DIST)/*
	@cd $(VUE_SOURCES) && npm install && npm run build
	@mv $(VUE_SOURCES)/dist/* $(VUE_DIST)/
	@rm -rf $(VUE_SOURCES)/dist/


## dev mode run in server mode
serve: clean dist ollama
	go run ./$(GO_SOURCES) serve


## dev mode populate rag
embed: clean ollama
	go run ./$(GO_SOURCES) embed pdfs/


## clean shit
clean:
	@go clean
	@rm -rf $(VUE_DIST)/*
	@rm -rf $(GO_BINARY_FOLDER)
	go mod download


# install dependencies
dependencies:
	go mod download
	cd ${VUE_SOURCES}; npm install


## pkg build
build: clean dist
	@rm -rf $(GO_BINARY_FOLDER)
	@go build -o $(GO_BINARY_FOLDER)/$(GO_BINARY_NAME) ./$(GO_SOURCES)


## docker 
docker-build:
	docker build -t ${DOCKER_TAG}:latest -f ./${DOCKER_FILE} .

docker-compose:
	docker compose --env-file=.env -f ${DOCKER_COMPOSE} up

docker-embed:
	docker exec -it ragbot devops-ragbot embed pdfs/

docker-milvus:
	docker compose -f ${MILVUS_COMPOSE} up


## ollama model
ollama:
	@echo "Downloading embeddinggemma:300m"
	ollama pull embeddinggemma:300m
	
	@echo "Downloading qwen3:14b"
	ollama pull qwen3:14b

	@echo "Downloading gemma3:4b"
	ollama pull gemma3:4b

	@echo "Installing custom model"
	ollama create ramaureirac/devops:latest -f ./model/devops.modelfile.md
	ollama create ramaureirac/generator:latest -f ./model/generator.modelfile.md

## vulnerabilities
vuln: clean
	- @echo govulncheck:
	- @${GO_VULN} ./src/...
	- @echo npm audit:
	- @cd ${VUE_SOURCES} ; npm audit