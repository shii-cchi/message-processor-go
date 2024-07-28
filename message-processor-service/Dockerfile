FROM golang:1.22.5-alpine
RUN apk update && apk upgrade && apk add pkgconf git bash build-base sudo
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
ARG DB_URI
ENV DB_URI=$DB_URI
WORKDIR message-processor-service
COPY message-processor-service .
RUN go mod download
RUN go build -tags musl --ldflags "-extldflags -static" -o server cmd/main.go
CMD ["sh", "-c", "goose -dir internal/database/migrations postgres ${DB_URI} up && ./server"]
