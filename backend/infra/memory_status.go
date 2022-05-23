package infra

import (
	"runtime"
)

type MemoryStatus struct {
	Alloc      uint64 `json:"alloc"`
	TotalAlloc uint64 `json:"total_alloc"`
	HeapAlloc  uint64 `json:"heap_alloc"`
}

func GetMemoryStatus() *MemoryStatus {
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)
	return &MemoryStatus{
		Alloc:      m1.Alloc,
		TotalAlloc: m1.TotalAlloc,
		HeapAlloc:  m1.HeapAlloc,
	}
}
