FROM golang:1.18-buster AS builder

MAINTAINER Maintainer
ENV GIN_MODE=release
ENV PORT=8080

ADD . /RiOS
WORKDIR /RiOS

#RUN echo -e "https://nl.alpinelinux.org/alpine/v3.12/main\nhttps://nl.alpinelinux.org/alpine/v3.12/community" > /etc/apk/repositories

# install git and dependencies for the project.
#RUN apk update
#RUN apk add --no-cache git
RUN go get ./...

RUN go mod download
#RUN CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -a -o /main .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /main .

#FROM scratch
FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    curl && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /main ./
ENTRYPOINT ["./main"]

EXPOSE $PORT