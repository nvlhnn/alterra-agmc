FROM golang:1.18-alpine as builder
LABEL maintainer="Naufal Hanan"
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o /alterra-day-7
EXPOSE 8080
CMD ["/alterra-day-7"]


