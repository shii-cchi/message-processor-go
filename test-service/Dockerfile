FROM golang:1.22.5-alpine
RUN apk update && apk upgrade && apk add pkgconf git bash build-base sudo
WORKDIR /test-service
COPY . .
RUN go mod download
RUN go build -tags musl --ldflags "-extldflags -static" -o test_service .
CMD ["./test_service"]