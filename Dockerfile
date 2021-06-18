##
## Build
##

FROM golang:1.16-buster AS build

WORKDIR /go/src/github.com/anoop-b/go-ifsc/
COPY .  .
RUN go mod download

RUN go build -tags=jsoniter -o /go-ifsc

##
## Deploy
##

FROM gcr.io/distroless/base

WORKDIR /

COPY --from=build /go-ifsc /go-ifsc

ENV GIN_MODE=release

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/go-ifsc"]