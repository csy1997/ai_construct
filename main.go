package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// timeHandler 处理根路径请求，返回当前时间
func timeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	currentTime := time.Now().Format(time.RFC1123)
	response := fmt.Sprintf("Current Time: %s\n", currentTime)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))

	log.Printf("Handled request: %s %s - Response: %s", r.Method, r.URL.Path, currentTime)
}

// GetCurrentTime 获取当前时间（用于测试）
func GetCurrentTime() string {
	return time.Now().Format(time.RFC1123)
}

func main() {
	http.HandleFunc("/", timeHandler)

	addr := ":8080"
	log.Printf("Starting HTTP server on http://localhost%s", addr)
	log.Printf("Listening for requests...")

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}