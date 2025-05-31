FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o bin/website .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bin/website .
CMD ["./website"]