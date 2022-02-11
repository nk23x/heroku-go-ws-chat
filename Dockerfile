FROM golang:1.17.2-alpine3.14 as builder
COPY go.mod go.sum /go/src/github.com/SajjadManafi/ws-chat/
WORKDIR /go/src/github.com/SajjadManafi/ws-chat/
RUN go mod download
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/web
EXPOSE 8080 8080
CMD ["./web"]