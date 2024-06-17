FROM golang:1.21 as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -v -o app ./cmd/main.go

FROM gcr.io/distroless/base-debian11 as final

COPY --from=builder /app/ .

EXPOSE 42069

ENTRYPOINT ["./app"]