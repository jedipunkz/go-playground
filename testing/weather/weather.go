package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Handler は /hello?name=xxx に JSON で挨拶を返す
func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "hello " + name})
}

// Client は外部の天気 API を呼び出すクライアント
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

type Forecast struct {
	City    string `json:"city"`
	Weather string `json:"weather"`
}

func (c *Client) Get(ctx context.Context, city string) (*Forecast, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.BaseURL+"/forecast?city="+city, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}
	var f Forecast
	if err := json.NewDecoder(resp.Body).Decode(&f); err != nil {
		return nil, err
	}
	return &f, nil
}
