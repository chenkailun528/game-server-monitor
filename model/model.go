package model

import "time"

type ServerInfo struct {
	ServerID   string    `json:"server_id"`
	Name       string    `json:"name"`
	IP         string    `json:"ip"`
	Status     string    `json:"status"`
	Online     int       `json:"online"`
	CPU        float64   `json:"cpu"`
	Memory     float64   `json:"memory"`
	LastReport time.Time `json:"last_report"`
}

type MetricReport struct {
	ServerID string  `json:"server_id" binding:"required"`
	Online   int     `json:"online"`
	CPU      float64 `json:"cpu"`
	Memory   float64 `json:"memory"`
}

type AlertRecord struct {
	ServerID string    `json:"server_id"`
	Type     string    `json:"type"`
	Message  string    `json:"message"`
	Level    string    `json:"level"`
	Time     time.Time `json:"time"`
}

type OperationLog struct {
	ServerID  string    `json:"server_id"`
	Operation string    `json:"operation"`
	Result    string    `json:"result"`
	Time      time.Time `json:"time"`
}
