FROM golang:alpine AS builder

MAINTAINER Maintainer
ENV GIN_MODE=release
ENV PORT=9080

ADD . /RiOS
WORKDIR /RiOS

# install git and dependencies for the project.
RUN apk update && apk add --no-cache git
RUN go get ./...

RUN go mod download
#RUN CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -a -o /main .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /main .

FROM scratch

COPY --from=builder /main ./
ENTRYPOINT ["./main"]

EXPOSE $PORT