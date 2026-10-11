package cache

import (
	"context"
	"testing"
	"testing/synctest"
	"time"
)

func TestCacheTTL(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		c := New(5 * time.Minute)
		c.Set("user:1", "jedi")

		start := time.Now()

		// bubble 内では偽の時計が使われるので、実時間では一瞬で終わる
		time.Sleep(4 * time.Minute)
		if _, ok := c.Get("user:1"); !ok {
			t.Fatal("期限切れ前にエントリが消えている")
		}

		time.Sleep(2 * time.Minute)
		if _, ok := c.Get("user:1"); ok {
			t.Fatal("期限切れ後もエントリが残っている")
		}
		t.Logf("start = %v, elapsed = %v", start, time.Since(start))
	})
}

func TestCacheCleanup(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithCancel(t.Context())
		defer cancel()

		c := New(time.Minute)
		c.StartCleanup(ctx, 10*time.Second)
		c.Set("user:1", "jedi")

		// 時間を進め、cleanup goroutine が処理を終えて待機状態になるまで待つ
		synctest.Sleep(time.Minute + 10*time.Second)
		if n := c.Len(); n != 0 {
			t.Fatalf("Len() = %d, want 0", n)
		}
	})
}
