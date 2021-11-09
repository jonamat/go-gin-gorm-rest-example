FROM golang:1.17.0-bullseye AS builder
WORKDIR /build

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY . .

# Create statically linked server binary
RUN go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o ./bin/server ./cmd/server

FROM scratch AS runner
WORKDIR /app

COPY --from=builder /build/bin/server ./bin/server
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# Default config
COPY ./.env ./.env

ENTRYPOINT ["/app/bin/server"]