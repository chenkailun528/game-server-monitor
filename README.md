# Game Server Monitor

A lightweight game server monitoring backend built with Go + Gin.

## Features

- Server list
- Metric report API
- CPU / memory alert
- Heartbeat check
- Log search
- Mock restart / maintenance mode
- Operation logs

## Project structure

```text
game-server-monitor/
├── main.go
├── router/
├── controller/
├── service/
├── model/
├── repository/
└── mock/
```

## Run

```bash
go mod tidy
go run main.go
```

## Test

```bash
curl http://127.0.0.1:8080/api/servers
curl http://127.0.0.1:8080/api/alerts
curl "http://127.0.0.1:8080/api/logs?keyword=error"
curl -X POST http://127.0.0.1:8080/api/servers/s1001/restart
```
