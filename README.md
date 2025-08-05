# 📈 auto-go-app

`auto-go-app` is a backend service written in Go.

## 🚀 How to Run the Application

### 1️⃣ Set Up Environment Variables

Before running the application, make sure you have a `.env` file in the root folder with the required configurations:

```env
# ===================================
# Common Settings
# ===================================
GRACEFUL_SHUTDOWN_PERIOD=30s

# ===================================
# Global Log Settings
# ===================================
LOG_LEVEL=info

# ===================================
# Metric Server Settings
# ===================================
METRIC_SERVER_HOST=0.0.0.0
METRIC_SERVER_PORT=9092
METRIC_SERVER_READ_TIMEOUT=30s
METRIC_SERVER_WRITE_TIMEOUT=30s

# ===================================
# HTTP Server Settings
# ===================================
HTTP_SERVER_HOST=0.0.0.0
HTTP_SERVER_PORT=9080
HTTP_SERVER_READ_TIMEOUT=30s
HTTP_SERVER_WRITE_TIMEOUT=30s
HTTP_ENABLE_CORS=false
HTTP_BODY_LIMIT=100K
HTTP_SKIP_REQUEST_ID=true
```

If the file does not exist, create it manually:

```sh
touch .env
```

### 2️⃣ Start the Application

Once the .env is ready, run the application:

```sh
go run main.go
```

Now your application should be up and running! 🚀
