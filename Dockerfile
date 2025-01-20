FROM golang:1.22.10

WORKDIR /app

COPY . . 

RUN go mod download

RUN go build -o main .

EXPOSE 3000

CMD ["./main"]