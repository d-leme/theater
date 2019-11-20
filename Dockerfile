FROM golang:1.12-stretch AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=1

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build ./cmd/server/main.go

FROM scratch

COPY --chown=0:0 --from=builder /build /

ENTRYPOINT ["./main"]
