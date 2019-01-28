FROM golang:1.11-alpine
ADD . /go/src/github.com/nonemax/porto/clientapi
EXPOSE 8081
WORKDIR /go/src/github.com/nonemax/porto/clientapi
RUN go install

ENTRYPOINT ["clientapi"]