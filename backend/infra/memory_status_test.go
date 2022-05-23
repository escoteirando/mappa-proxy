package infra

import (
	"testing"
)

func TestGetMemoryStatus(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		if got := GetMemoryStatus(); got.Alloc <= 0 {
			t.Errorf("Expected Alloc memory greater than 0: %v", got)
		}
	})

}
