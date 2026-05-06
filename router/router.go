package router

import (
	"game-server-monitor/controller"

	"github.com/gin-gonic/gin"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		if origin == "http://127.0.0.1:5173" || origin == "http://localhost:5173" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		}

		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(corsMiddleware())

	api := r.Group("/api")
	{
		api.GET("/servers", controller.GetServers)
		api.GET("/servers/:id", controller.GetServerDetail)
		api.POST("/metrics/report", controller.ReportMetric)
		api.GET("/alerts", controller.GetAlerts)
		api.GET("/logs", controller.QueryLogs)
		api.POST("/servers/:id/restart", controller.RestartServer)
		api.POST("/servers/:id/maintenance", controller.MaintenanceServer)
		api.GET("/operations", controller.GetOperationLogs)
	}
	return r
}
