package mock

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func StartMockAgent() {
	go func() {
		client := &http.Client{Timeout: 3 * time.Second}
		for {
			online := rand.Intn(3000)
			cpu := 20 + rand.Float64()*75
			memory := 30 + rand.Float64()*60
			body := fmt.Sprintf(`{"server_id":"s1001","online":%d,"cpu":%.2f,"memory":%.2f}`, online, cpu, memory)
			req, err := http.NewRequest("POST", "http://127.0.0.1:8080/api/metrics/report", strings.NewReader(body))
			if err == nil {
				req.Header.Set("Content-Type", "application/json")
				_, _ = client.Do(req)
			}
			time.Sleep(5 * time.Second)
		}
	}()
}
