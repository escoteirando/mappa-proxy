package fiberfilestorage

import (
	"os"
	"path"
	"testing"
	"time"
)

func TestNew(t *testing.T) {

	t.Run("Default", func(t *testing.T) {
		store := New(&Config{
			BasePath:   path.Join(os.TempDir(), "fiberfilestorage"),
			GCInterval: 2 * time.Second,
		})
		const key = "test"
		const value = "TESTING DATA"
		store.Set(key, []byte(value), time.Second)
		var content []byte
		var err error
		if content, err = store.Get(key); err != nil {
			t.Error(err)
		}
		if string(content) != value {
			t.Errorf("Expected %s, got %s", value, string(content))
		}
		time.Sleep(4 * time.Second)
		if _, err := store.Get(key); err == nil {
			t.Error("Should be expired")
		}
	})

}
