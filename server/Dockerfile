#
FROM golang AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go .
RUN CGO_ENABLED=0 go build -o server .

COPY . .

#
FROM scratch

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8090
CMD ["/app/server"]