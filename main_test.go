package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// TestTimeHandler 测试 timeHandler 函数
func TestTimeHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
		shouldContain  string
	}{
		{
			name:           "GET request to root path",
			method:         http.MethodGet,
			path:           "/",
			expectedStatus: http.StatusOK,
			shouldContain:  "Current Time:",
		},
		{
			name:           "POST request to root path",
			method:         http.MethodPost,
			path:           "/",
			expectedStatus: http.StatusMethodNotAllowed,
			shouldContain:  "Method not allowed",
		},
		{
			name:           "GET request to invalid path",
			method:         http.MethodGet,
			path:           "/invalid",
			expectedStatus: http.StatusNotFound,
			shouldContain:  "404",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			w := httptest.NewRecorder()

			timeHandler(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			body := w.Body.String()
			if !strings.Contains(body, tt.shouldContain) {
				t.Errorf("Expected response to contain '%s', got '%s'", tt.shouldContain, body)
			}
		})
	}
}

// TestGetCurrentTime 测试 GetCurrentTime 函数
func TestGetCurrentTime(t *testing.T) {
	result := GetCurrentTime()

	// 检查返回值不为空
	if result == "" {
														t.Error("GetCurrentTime returned empty string")
		}

	// 尝试解析返回值，确保是有效的 RFC1123 格式
	_, err := time.Parse(time.RFC1123, result)
	if err != nil {
		t.Errorf("GetCurrentTime returned invalid RFC1123 format: %v", err)
	}

	t.Logf("GetCurrentTime returned: %s", result)
}

// TestTimeFormatRFC1123 测试时间格式
func TestTimeFormatRFC1123(t *testing.T) {
	// 创建一个固定时间用于测试
	fixedTime := time.Date(2026, 4, 13, 10, 30, 45, 0, time.UTC)
	formatted := fixedTime.Format(time.RFC1123)

	// RFC1123 格式应该类似: "Mon, 13 Apr 2026 10:30:45 UTC"
	expectedParts := []string{"Apr", "2026", "10:30:45", "UTC"}
	for _, part := range expectedParts {
		if !strings.Contains(formatted, part) {
			t.Errorf("Expected formatted time to contain '%s', got '%s'", part, formatted)
		}
	}

	t.Logf("Formatted time: %s", formatted)
}