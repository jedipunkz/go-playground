package weather

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"testing/synctest"
	"time"
)

func TestClientTimeout(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		// 応答に 10 秒かかる外部 API を、インメモリの偽ネットワーク上に立てる
		srv := httptest.NewTestServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			select {
			case <-time.After(10 * time.Second):
				json.NewEncoder(w).Encode(Forecast{City: "tokyo", Weather: "sunny"})
			case <-r.Context().Done(): // クライアントが諦めたら終了する
			}
		}))

		hc := srv.Client()
		hc.Timeout = 3 * time.Second
		c := &Client{BaseURL: srv.URL, HTTPClient: hc}

		start := time.Now()
		_, err := c.Get(t.Context(), "tokyo")
		if err == nil {
			t.Fatal("Get() error = nil, want timeout")
		}
		t.Logf("err = %v, elapsed (fake clock) = %v", err, time.Since(start))
	})
}
