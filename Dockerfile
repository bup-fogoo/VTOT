FROM golang:latest

MAINTAINER foo
WORKDIR /VTOT
COPY . /VTOT
COPY go.mod ./
COPY go.sum ./
COPY * ./
RUN go env -w GO111MODULE=auto &&\
    go env -w GOPROXY=https://goproxy.cn,direct &&\
    go build -mod=mod main.go


EXPOSE 12333
CMD ["./main"]