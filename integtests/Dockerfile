FROM golang:1.11-alpine
ADD . /go/src/github.com/nonemax/porto/integtests
WORKDIR /go/src/github.com/nonemax/porto/integtests
RUN go install

ENTRYPOINT ["integtests"]