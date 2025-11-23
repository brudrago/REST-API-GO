FROM golang:1.23

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main cmd/main.go

EXPOSE 8000

CMD ["./main"]  
