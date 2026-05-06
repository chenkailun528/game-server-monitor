package router

import (
	"game-server-monitor/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
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
