# Shortener Go App

A Go web app that shortens URLs.

## Prerequisites

- Go 1.24 or later
- Docker or Docker Compose

## Getting Started

Clone the repository:

```bash
git clone https://github.com/josearpaiaq/shortener.git
```

Create a `.env` file in the root directory (or copy the values from the `.env.example` file) filling in the required values.

```bash
cp .env.example .env
```

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

Build and run the image:

```bash
docker compose up --buuld -d
```

Or run it with the `docker run` command:

```bash
docker run -p 8080:8080 shortener
```

Down the image:

```bash
docker compose down -v
```

## License

This project is not under any license, but it is open-source and free to use. Feel free to use it for your own projects.
