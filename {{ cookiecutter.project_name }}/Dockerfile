FROM golang:1.20 AS builder
WORKDIR /go/src/{{ cookiecutter.project_name }}
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOPROXY=https://goproxy.cn,direct go build -o {{ cookiecutter.project_name }} cmd/main.go cmd/wire_gen.go



FROM alpine:latest
CMD  ./{{ cookiecutter.project_name }}

EXPOSE 8702

