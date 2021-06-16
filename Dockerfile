FROM golang:latest AS builder
WORKDIR /go/src/github.com/anoop-b/go-ifsc/
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/anoop-b/go-ifsc/app .
ENV GIN_MODE=release
EXPOSE 8080
CMD ["./app"]