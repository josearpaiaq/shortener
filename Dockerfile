FROM golang:1.24.4

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .

RUN go build -o /app/app .

EXPOSE 8080

CMD ["/app/app"]