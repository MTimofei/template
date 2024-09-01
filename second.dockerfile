FROM golang:1.22 AS builder

WORKDIR /builder
COPY . .
RUN go mod download

# RUN go build -o ./bin/app ./cmd/second/main.go

# FROM alpine AS runner

expose 10000

CMD ["go", "run", "cmd/second/main.go"]
# COPY --from=builder /builder/bin /
# CMD ["/bin/app"]
