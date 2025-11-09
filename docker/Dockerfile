# builder
FROM alpine AS builder

WORKDIR /opt

RUN apk add go npm nodejs make

COPY src/ src/
COPY go.mod go.sum Makefile .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 make build

# deploy app in lfs
FROM scratch

WORKDIR /srv

COPY --from=builder /opt/bin/devops-ragbot /bin/devops-ragbot

USER 1001:1001
ENTRYPOINT ["/bin/devops-ragbot"]