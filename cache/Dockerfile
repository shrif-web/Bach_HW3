#
FROM golang AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o cache_server .

#
FROM scratch

WORKDIR /app

COPY --from=builder /app/cache_server .

EXPOSE 8090
CMD ["/app/cache_server"]