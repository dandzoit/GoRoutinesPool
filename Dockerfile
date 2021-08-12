FROM golang:latest
RUN mkdir /app
COPY . /app
WORKDIR /app/mylib
CMD go test
