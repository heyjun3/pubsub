FROM golang:1.23.3 AS builder

WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build -o /build/app ./cmd/server/main.go

FROM gcr.io/distroless/static-debian12

COPY --from=builder /build/app /
CMD ["/app"]
