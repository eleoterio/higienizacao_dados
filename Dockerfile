FROM golang:latest

WORKDIR /go/src/

ADD . /go/src/

RUN go get github.com/jmoiron/sqlx
RUN go get github.com/lib/pq

RUN mkdir -p github.com/eleoterio/neoway
RUN cp -R service github.com/eleoterio/neoway

RUN cd github.com/eleoterio/neoway