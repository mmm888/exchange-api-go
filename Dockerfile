From golang:1.9

MAINTAINER mmm888

RUN apt-get update; apt-get -y install vim

COPY cmd/goanda/goanda ${GOPATH}/bin

CMD goanda streaming
