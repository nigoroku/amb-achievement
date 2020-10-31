FROM golang:1.14.10-alpine3.12

WORKDIR /go/src/github.com/nigoroku/amb-achievement
ADD . /go/src/github.com/nigoroku/amb-achievement

ENV GO111MODULE=on

RUN apk add --no-cache \
    alpine-sdk \
    git \
    && go get github.com/pilu/fresh

COPY go.mod go.sum ./
RUN go mod download

EXPOSE 8083

CMD ["fresh"]