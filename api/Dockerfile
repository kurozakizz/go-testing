FROM golang:1.9.2-alpine

ENV DB_HOST=mongodb

RUN export CGO_ENABLED=1
RUN export GOOS=darwin

WORKDIR /go/src/api

ENTRYPOINT [ "go" ]
CMD [ "run" ]
