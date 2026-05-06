# Game Server Monitor

Go + Gin 后端，已加入 CORS 处理，支持 Vue 前端访问。

## 修复内容

- router/router.go 增加跨域中间件
- 支持前端地址：
  - http://127.0.0.1:5173
  - http://localhost:5173

## 启动

```bash
go mod tidy
go run main.go
```

## 测试

```bash
curl http://127.0.0.1:8080/api/servers
curl http://127.0.0.1:8080/api/alerts
curl "http://127.0.0.1:8080/api/logs?keyword=error"
```
