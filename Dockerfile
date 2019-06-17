FROM golang:1.12-alpine AS builder
# Force the go compiler to use modules
ENV GO111MODULE=on 
# update git
RUN apk add --no-cache --update git
ARG project=ping-api
# set work directory of container
WORKDIR $GOPATH/src/$project/
# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .
# Download all the dependencies that are specified in 
# the go.mod and go.sum file.
RUN go mod download
# copy project to container
COPY . .

RUN rm -rf /go/pkg/mod/github.com/coreos/etcd@v3.3.10+incompatible/client/keys.generated.go
# build project
RUN CGO_ENABLED=0 GOARCH="amd64" GOOS="linux" \
	go build -o /bin/$project .


FROM alpine:latest
# copy binary from previous build to new container
COPY --from=builder /bin/$project /bin/$project
# set environment for consul
# public consul url should be here
ENV PING_API_CONSUL_URL="127.0.0.1:8500"
ENV PING_API_CONSUL_PATH=$project

EXPOSE 9001
ENTRYPOINT ["/bin/ping-api", "serve"]