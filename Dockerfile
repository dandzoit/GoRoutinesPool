FROM golang:latest
RUN mkdir /usr/local/go/src/app
ADD . /usr/local/go/src/app
WORKDIR /usr/local/go/src/app/mylib
CMD go test
