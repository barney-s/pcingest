FROM golang:1.15 as build-env

WORKDIR /go/src/pcingest
ADD . /go/src/pcingest

RUN go get -d -v ./...

RUN go build -o /go/bin/pcingest

FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/pcingest /
CMD ["/pcingest"]
