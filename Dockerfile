# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

ENV GO111MODULE=on
ENV PORT=8090

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

EXPOSE 8080
ENTRYPOINT ["/app/go-vue-socket-chat"] 

#RUN go build -o /go-socket-chat

#EXPOSE 8080

#CMD [ "/go-socket-chat" ]