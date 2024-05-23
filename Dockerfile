FROM golang:latest as BUILDER
LABEL maintainer="shishupei"

# build binary
RUN mkdir -p /go/src/github.com/opensourceways/message-collect
COPY . /go/src/github.com/opensourceways/message-collect
RUN cd /go/src/github.com/opensourceways/message-collect && CGO_ENABLED=1 go build -v -o ./message-collect main.go

# copy binary config and utils
FROM openeuler/openeuler:22.03

COPY --from=BUILDER /go/src/github.com/opensourceways/message-collect /opt/app
WORKDIR /opt/app/
EXPOSE 8081
ENTRYPOINT ["/opt/app/message-collect"]