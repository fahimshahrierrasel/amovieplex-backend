FROM golang:latest

LABEL maintainer="Md. Fahim Shahrier Rasel <fahimshahrier2@gmail.com>"

WORKDIR /app
COPY . /app

RUN go mod download
RUN go mod vendor