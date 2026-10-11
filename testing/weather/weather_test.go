package weather

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		name       string
		target     string
		wantStatus int
	}{
		{"name あり", "/hello?name=jedi", http.StatusOK},
		{"name なし", "/hello", http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.target, nil)
			rec := httptest.NewRecorder() // レスポンスを記録する ResponseWriter

			Handler(rec, req)

			if rec.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
		})
	}
}

func TestClientGet(t *testing.T) {
	// 外部 API の代わりになるテスト用サーバ
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.URL.Query().Get("city"); got != "tokyo" {
			t.Errorf("city = %q, want %q", got, "tokyo")
		}
		json.NewEncoder(w).Encode(Forecast{City: "tokyo", Weather: "sunny"})
	}))
	defer srv.Close()

	c := &Client{BaseURL: srv.URL, HTTPClient: srv.Client()}
	f, err := c.Get(t.Context(), "tokyo")
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if f.Weather != "sunny" {
		t.Errorf("Weather = %q, want %q", f.Weather, "sunny")
	}
}
