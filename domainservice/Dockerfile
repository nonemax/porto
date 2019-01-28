FROM golang:1.11-alpine
ADD . /go/src/github.com/nonemax/porto/domainservice
EXPOSE 8080
WORKDIR /go/src/github.com/nonemax/porto/domainservice
RUN go install

ENTRYPOINT ["domainservice"]