FROM golang:latest

MAINTAINER foo
WORKDIR /VTOT
WORKDIR $GOPATH/go
COPY . $GOPATH/go
COPY go.mod ./
COPY go.sum ./
COPY * ./
RUN go env -w GO111MODULE=auto
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build -mod=mod main.go


EXPOSE 12333
CMD ["./main"]