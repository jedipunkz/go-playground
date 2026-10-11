package config

import (
	"context"
	"os"
	"strings"
)

type Config struct {
	Env  string
	Port string
}

// Load は path のファイルから設定を読み込む
// 環境変数 APP_ENV が設定されていれば Env を上書きする
func Load(ctx context.Context, path string) (*Config, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	for line := range strings.Lines(string(b)) {
		k, v, ok := strings.Cut(strings.TrimSpace(line), "=")
		if !ok {
			continue
		}
		switch k {
		case "env":
			cfg.Env = v
		case "port":
			cfg.Port = v
		}
	}
	if env := os.Getenv("APP_ENV"); env != "" {
		cfg.Env = env
	}
	return cfg, nil
}
