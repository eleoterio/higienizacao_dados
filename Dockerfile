FROM golang:1.8

RUN mkdir -p /go/src/app
RUN go get github.com/jmoiron/sqlx
WORKDIR /go/src/app

ADD . /go/src/app

RUN go get -v