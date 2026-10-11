package config

import (
	"os"
	"path/filepath"
	"testing"
)

// writeConfig はテスト用の設定ファイルを一時ディレクトリに作成するヘルパー
func writeConfig(t *testing.T, content string) string {
	t.Helper() // 失敗時の行番号を呼び出し元にする

	// テスト終了時に自動で削除される一時ディレクトリ
	path := filepath.Join(t.TempDir(), "app.conf")
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("write config: %v", err)
	}
	return path
}

func TestLoad(t *testing.T) {
	path := writeConfig(t, "env=dev\nport=8080\n")

	// テストが終わる直前にキャンセルされる context
	cfg, err := Load(t.Context(), path)
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if cfg.Env != "dev" || cfg.Port != "8080" {
		t.Errorf("Load() = %+v, want Env=dev Port=8080", cfg)
	}
}

func TestLoadEnvOverride(t *testing.T) {
	// テスト終了時に元の値へ戻る環境変数
	t.Setenv("APP_ENV", "prod")
	path := writeConfig(t, "env=dev\nport=8080\n")

	cfg, err := Load(t.Context(), path)
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if cfg.Env != "prod" {
		t.Errorf("Env = %q, want %q", cfg.Env, "prod")
	}
}

func TestLoadRelativePath(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "app.conf"), []byte("port=9090\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	// テスト終了時に元のディレクトリへ戻る
	t.Chdir(dir)

	cfg, err := Load(t.Context(), "app.conf")
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	t.Cleanup(func() {
		t.Log("cleanup: テスト終了後に実行される")
	})
	if cfg.Port != "9090" {
		t.Errorf("Port = %q, want %q", cfg.Port, "9090")
	}
}
