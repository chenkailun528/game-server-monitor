package main

import (
	"fmt"
	"game-server-monitor/mock"
	"game-server-monitor/router"
	"game-server-monitor/service"
)

func main() {
	service.InitData()
	service.EnsureDemoLog()
	service.StartHeartbeatChecker()
	mock.StartMockAgent()

	r := router.SetupRouter()
	fmt.Println("Game Server Monitor started: http://127.0.0.1:8080")
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
