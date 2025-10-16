FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o main .

ENV APP_HOST=0.0.0.0:8080

EXPOSE 8080

CMD ["./main"]