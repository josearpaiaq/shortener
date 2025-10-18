# 🔗 Go URL Shortener

A simple and efficient URL Shortener built with **Go (net/http + gorilla/mux)** and **PostgreSQL**, designed to generate and manage short links easily.  
This project was developed as part of my learning process with Go and backend development best practices.

---

## 🧭 Overview

The **Go URL Shortener** allows users to create shorter versions of long URLs, making them easier to share and manage.  
When a shortened link is accessed, the application redirects the user to the original URL stored in the database.

The main goal of the project was to strengthen my backend development skills using Go, Docker, PostgreSQL, and deployment in a real environment.

---

## ⚙️ Technologies Used

- **Go (net/http + gorilla/mux)** — for routing and HTTP server.
- **PostgreSQL** — database for storing original and shortened URLs.
- **Docker & Docker Compose** — for local development and service orchestration.
- **Render.com** — for hosting the PostgreSQL instance and deploying the Go app.
- **HTML/CSS** — for the landing page.

---

## 🏗️ Architecture

The project follows a simple **client-server architecture**:

```
Browser (Client)
   ↓
HTTP Server (Go)
   ↓
PostgreSQL Database
```

Each short URL record includes:
- The **original URL**
- The **shortened path**
- The **number of visits (clicks)**

---

## 📁 Project Structure

```
├── main.go
├── routes/
│   └── main.go
├── handlers/
│   └── main.go
├── models/
│   └── url.go
├── templates/
│   └── index.html
├── .env
├── Dockerfile
├── docker-compose.yml
└── README.md
```

---

## 🐳 Running Locally with Docker

You can run the project locally using Docker and Docker Compose.

```bash
docker compose up --build
```

> ⚠️ **Note:**  
> I encountered several issues deploying the multi-container setup (Go + PostgreSQL) with Docker Compose.  
> For deployment, I took an alternative approach — running a **PostgreSQL instance on Render.com** and deploying the Go application separately as a single service connected to that database.  
> 
> Ideally, the project should run as a full multi-container setup (like with **VPS** or **Railway**), and I’m currently working toward that.

---

## 🧩 Environment Variables

The application expects the following variables in a `.env` file:

```
PORT=8080
DB_HOST=<host>
DB_PORT=<port>
DB_USER=<user>
DB_PASSWORD=<password>
DB_NAME=<database_name>
```

---

## 🚀 Deployment

The current deployment setup uses:

- **PostgreSQL Database**: hosted on Render.com  
- **Go Application**: deployed as a Render “Web Service”  
  - The service runs the compiled Go binary and connects to the PostgreSQL instance using environment variables.

---

## 🌐 API Endpoints

| Method | Endpoint                | Description                         |
|--------|--------------------------|-------------------------------------|
| `POST` | `/shorten`              | Creates a new shortened URL         |
| `GET`  | `/{short_url}`          | Redirects to the original URL       |
| `GET`  | `/urls`                 | (Optional) Lists all URLs           |

Example request:
```bash
POST /shorten
{
  "url": "https://www.example.com"
}
```

Example response:
```json
{
  "short_url": "http://localhost:8080/abc123"
}
```

---

## 🧠 Future Improvements

As detailed in [TODOs.md](./TODOs.md), planned improvements include:

### 🔥 Development
- Implement **hot reloading** for faster local development using [Air](https://github.com/cosmtrek/air) or [CompileDaemon](https://github.com/githubnemo/CompileDaemon).

### 🎨 Frontend
- Improve the **landing page** with **TailwindCSS**.
- Add:
  - Input form for URLs
  - Buttons to shorten and view results
  - Display of original/short URLs and click counts
  - Favicon and better layout

### 🧱 Database
- Add **migrations** using [GORM Migrations](https://gorm.io/docs/migration.html).

### 🧪 Testing
- Add unit tests and validations with [go-playground/validator](https://github.com/go-playground/validator).

### 🔐 Authentication
- Add user authentication:
  - Login / Logout
  - User-specific URLs
  - Display of click statistics per user

---

## 🧑‍💻 Author

**José Arpaia**  
Systems Engineer & Full Stack Developer  
- 💼 Focused on Go, JavaScript, and Cloud development  
- 🌎 Based in Latin America  

---

## 📜 License

This project is open-source and available under the **MIT License**.

