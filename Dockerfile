FROM golang:1.24.4

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .

RUN ls -la /app
RUN go env && go version
RUN go build -v -x -o /app/app .

EXPOSE 8080

CMD ["/app/app"]