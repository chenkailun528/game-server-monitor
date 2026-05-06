# Game Server Monitor Web

Vue + Element Plus 前端页面。

## 启动

```bash
npm install
npm run dev
```

默认访问：

```text
http://127.0.0.1:5173
```

注意：需要先启动 Go 后端：

```bash
cd game-server-monitor
go run main.go
```

如果浏览器提示跨域，需要在 Go 后端加 CORS 中间件。
