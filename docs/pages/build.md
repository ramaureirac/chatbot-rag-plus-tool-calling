# Building from source

The project is distributed as a Docker package, but you are completely free to compile, distribute, and deploy the code base on any operating system supported by Go, including (but not limited to) *Windows, macOS, Linux, BSD (FreeBSD/OpenBSD/NetBSD), Solaris, Illumos. 

Please, keep in mind this deployment method is unofficial and is provided without formal support or guarantee of maintenance.

## Building from source (using Make)

    make build

## Building from source for Windows

    GOOS=windows GOARCH=amd64 go build -o devops-ragbot.exe

## Building from source for Linux

    GOOS=linux GOARCH=amd64 go build -o devops-ragbot

## Run compiled binary

    ./bin/devops-ragbot                 # run in local mode
    ./bin/devops-ragbot serve           # run in server mode (http://localhost:8080)
    ./bin/devops-ragbot embed pdfs/     # populate rag where pdfs/ is path to files

## Notes

- You MUST install all external dependencies (Milvus, Ollama, Go, Node.js)
- Make sure your binary format is compatible with your CPU architecture.
- Remember creating your own `.env` file
