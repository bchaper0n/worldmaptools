FROM golang:1.24.0

WORKDIR /app

COPY . /app

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
