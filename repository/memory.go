package repository

import (
	"sync"
	"game-server-monitor/model"
)

var (
	Mu            sync.RWMutex
	Servers       = map[string]*model.ServerInfo{}
	Alerts        []model.AlertRecord
	OperationLogs []model.OperationLog
)
