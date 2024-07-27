FROM golang:1.22.5-alpine
RUN apk update && apk upgrade && apk add pkgconf git bash build-base sudo
WORKDIR /message-processor-go
COPY . .
RUN go mod download
RUN go build -tags musl --ldflags "-extldflags -static" -o server cmd/main.go
CMD ["./server"]
