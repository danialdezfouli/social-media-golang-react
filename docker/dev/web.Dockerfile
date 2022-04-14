# syntax=docker/dockerfile:1

FROM golang:1.18.1-alpine AS development

WORKDIR /app

COPY go.sum ./
COPY go.mod ./
RUN go mod download

COPY . .

RUN go run cmd/migrate.go

RUN go install github.com/cosmtrek/air@latest

#CMD reflex -g '*.go' go run cmd/main.go --start-service
CMD air