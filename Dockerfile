FROM golang:latest

LABEL maintainer="Md. Fahim Shahrier Rasel <fahimshahrier2@gmail.com>"

RUN go get github.com/oxequa/realize

WORKDIR /opt/app

COPY go.* /opt/app/

RUN go mod download
RUN go mod vendor

COPY . .