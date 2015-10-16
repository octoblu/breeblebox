FROM golang
MAINTAINER Octoblu, Inc. <docker@octoblu.com>
EXPOSE 80

ADD https://raw.githubusercontent.com/pote/gpm/v1.3.2/bin/gpm /go/bin/
RUN chmod +x /go/bin/gpm

WORKDIR /go/src/github.com/octoblu/breeblebox/
COPY Godeps /go/src/github.com/octoblu/breeblebox/
RUN gpm install

COPY . /go/src/github.com/octoblu/breeblebox/

VOLUME /export/

RUN go build -o breeblebox
