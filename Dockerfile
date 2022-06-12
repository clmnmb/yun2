FROM golang:1.17 AS builder
RUN  mkdir -p  /home/work/data/www/webgo
WORKDIR /home/work/data/www/webgo
COPY . .
RUN  go build -o /test
FROM centos:centos7
WORKDIR /home/work/data/www/webgo
COPY --from=builder test .
CMD ./test