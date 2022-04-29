FROM golang:1.18.1-alpine AS build

ENV HTTPS_PROXY="http://fodev.org:8118"

WORKDIR /app

COPY . ./

RUN go mod download && \
  # Build the app
  GOOS=linux GOARCH=amd64 go build -o main cmd/main.go && \
  # Make the final output executable
  chmod +x ./main

FROM alpine:latest

RUN apk --no-cache add bash

WORKDIR /app

COPY --from=build /app/main .

CMD ["./main"]

EXPOSE 3000