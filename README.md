# shortener

A Go web app that shortens URLs.

## Usage locally

Install the dependencies:

```bash
go mod download
```

Run the app:

```bash
go run main.go
```

## Usage in Docker

Build the image:

```bash
docker build -t shortener .
```

Run the image:

```bash
docker run -p 8080:8080 shortener
```

