FROM golang:latest

COPY ./ /app

WORKDIR /app

RUN go get ./...
RUN GOOS=linux GOARCH=amd64 go build -o main -buildvcs=false ./main.go

CMD ["./main"]