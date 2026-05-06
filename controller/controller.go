package controller

import (
	"net/http"
	"game-server-monitor/model"
	"game-server-monitor/service"
	"github.com/gin-gonic/gin"
)

func GetServers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": service.GetServers()})
}

func GetServerDetail(c *gin.Context) {
	server, ok := service.GetServer(c.Param("id"))
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "msg": "server not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": server})
}

func ReportMetric(c *gin.Context) {
	var req model.MetricReport
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	service.ReportMetric(req)
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "report success"})
}

func GetAlerts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": service.GetAlerts()})
}

func QueryLogs(c *gin.Context) {
	result, err := service.SearchLogFile(c.DefaultQuery("file", "./server.log"), c.Query("keyword"), 200)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": result})
}

func RestartServer(c *gin.Context) {
	service.RestartServer(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "mock restart success"})
}

func MaintenanceServer(c *gin.Context) {
	service.MaintenanceServer(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "switch maintenance success"})
}

func GetOperationLogs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": service.GetOperationLogs()})
}
