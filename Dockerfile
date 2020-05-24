FROM golang:1.13-alpine as builder
WORKDIR /var/build
COPY . /var/build/
RUN go build -o server cmd/main.go

FROM alpine
WORKDIR /opt/server
ENV CGO_ENABLED=0
COPY --from=builder /var/build/server /opt/server/server
CMD ["./server"]
