FROM golang:1.12-stretch AS builder

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Reference https://medium.com/@diogok/on-golang-static-binaries-cross-compiling-and-plugins-1aed33499671
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -a -tags netgo -ldflags '-w' ./cmd/server/main.go

ENTRYPOINT ["./main"]

FROM scratch

COPY --chown=0:0 --from=builder /build/main .

ENTRYPOINT ["./main"]
