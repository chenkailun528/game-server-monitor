# Game Server Monitor Web

Vue 3 + Element Plus 前端页面。

## 修复内容

- 增加 vite.config.js
- 正确启用 @vitejs/plugin-vue
- 强制监听 127.0.0.1:5173

## 启动

```bash
npm install
npm run dev
```

访问：

```text
http://127.0.0.1:5173
```

需要另开终端启动 Go 后端：

```bash
cd game-server-monitor
go run main.go
```
