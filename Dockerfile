#FIXME stop from polling proto and controller

FROM golang as build-env

WORKDIR /go/src/github.com/Daniel-W-Innes/shopify-add
ADD . /go/src/github.com/Daniel-W-Innes/shopify-add

RUN go mod download &&\
    go mod verify &&\
    go build -o /go/bin

FROM gcr.io/distroless/base
COPY --from=build-env /go/bin /
CMD ["/app"]