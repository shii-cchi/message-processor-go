FROM golang:1.22.5-alpine
WORKDIR /message-processor-go
COPY . .
RUN go build -o server cmd/main.go
CMD ["./server"]