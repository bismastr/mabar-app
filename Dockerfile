ARG GO_VERSION=1.22.4
FROM golang:${GO_VERSION}-bookworm as builder


WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /mabar-app ./cmd


FROM debian:bookworm
RUN apt-get update && apt-get install -y ca-certificates
COPY --from=builder /run-app /usr/local/bin/
CMD ["mabar-app"]