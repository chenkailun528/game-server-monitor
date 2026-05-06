package service

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"game-server-monitor/model"
	"game-server-monitor/repository"
)

func InitData() {
	repository.Servers["s1001"] = &model.ServerInfo{ServerID: "s1001", Name: "server-1", IP: "127.0.0.1", Status: "running", Online: 1200, CPU: 30, Memory: 55, LastReport: time.Now()}
	repository.Servers["s1002"] = &model.ServerInfo{ServerID: "s1002", Name: "server-2", IP: "127.0.0.2", Status: "running", Online: 900, CPU: 45, Memory: 60, LastReport: time.Now()}
}

func GetServers() []*model.ServerInfo {
	repository.Mu.RLock()
	defer repository.Mu.RUnlock()
	list := make([]*model.ServerInfo, 0, len(repository.Servers))
	for _, s := range repository.Servers {
		list = append(list, s)
	}
	return list
}

func GetServer(serverID string) (*model.ServerInfo, bool) {
	repository.Mu.RLock()
	defer repository.Mu.RUnlock()
	server, ok := repository.Servers[serverID]
	return server, ok
}

func ReportMetric(req model.MetricReport) {
	repository.Mu.Lock()
	defer repository.Mu.Unlock()
	s, ok := repository.Servers[req.ServerID]
	if !ok {
		s = &model.ServerInfo{ServerID: req.ServerID, Name: req.ServerID, IP: "unknown"}
		repository.Servers[req.ServerID] = s
	}
	s.Online, s.CPU, s.Memory, s.LastReport = req.Online, req.CPU, req.Memory, time.Now()
	s.Status = calcStatus(req.CPU, req.Memory)
	checkAlertLocked(s)
}

func calcStatus(cpu, memory float64) string {
	if cpu >= 90 || memory >= 90 {
		return "down"
	}
	if cpu >= 80 || memory >= 85 {
		return "warning"
	}
	return "running"
}

func checkAlertLocked(s *model.ServerInfo) {
	if s.CPU >= 80 {
		repository.Alerts = append(repository.Alerts, model.AlertRecord{ServerID: s.ServerID, Type: "cpu", Message: fmt.Sprintf("CPU too high: %.2f%%", s.CPU), Level: "warning", Time: time.Now()})
	}
	if s.Memory >= 85 {
		repository.Alerts = append(repository.Alerts, model.AlertRecord{ServerID: s.ServerID, Type: "memory", Message: fmt.Sprintf("Memory too high: %.2f%%", s.Memory), Level: "warning", Time: time.Now()})
	}
}

func GetAlerts() []model.AlertRecord {
	repository.Mu.RLock()
	defer repository.Mu.RUnlock()
	return append([]model.AlertRecord(nil), repository.Alerts...)
}

func StartHeartbeatChecker() {
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for range ticker.C {
			repository.Mu.Lock()
			for _, s := range repository.Servers {
				if time.Since(s.LastReport) > 60*time.Second {
					s.Status = "down"
					repository.Alerts = append(repository.Alerts, model.AlertRecord{ServerID: s.ServerID, Type: "heartbeat", Message: "server heartbeat timeout", Level: "critical", Time: time.Now()})
				}
			}
			repository.Mu.Unlock()
		}
	}()
}

func SearchLogFile(path string, keyword string, limit int) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if keyword == "" || strings.Contains(strings.ToLower(line), strings.ToLower(keyword)) {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if len(lines) > limit {
		lines = lines[len(lines)-limit:]
	}
	return lines, nil
}

func EnsureDemoLog() {
	if _, err := os.Stat("./server.log"); err == nil {
		return
	}
	content := []string{
		"2026-05-02 10:00:01 INFO server start success",
		"2026-05-02 10:01:12 WARN player reconnect",
		"2026-05-02 10:02:33 ERROR mysql timeout",
		"2026-05-02 10:04:15 ERROR cross server rpc timeout",
	}
	_ = os.WriteFile("./server.log", []byte(strings.Join(content, "\n")), 0644)
}

func RestartServer(serverID string) {
	AddOperation(serverID, "restart", "success: mock restart executed")
}

func MaintenanceServer(serverID string) {
	repository.Mu.Lock()
	if s, ok := repository.Servers[serverID]; ok {
		s.Status = "maintenance"
	}
	repository.Mu.Unlock()
	AddOperation(serverID, "maintenance", "success: switch to maintenance mode")
}

func AddOperation(serverID, operation, result string) {
	repository.Mu.Lock()
	defer repository.Mu.Unlock()
	repository.OperationLogs = append(repository.OperationLogs, model.OperationLog{ServerID: serverID, Operation: operation, Result: result, Time: time.Now()})
}

func GetOperationLogs() []model.OperationLog {
	repository.Mu.RLock()
	defer repository.Mu.RUnlock()
	return append([]model.OperationLog(nil), repository.OperationLogs...)
}
