package infra

import (
	"runtime"
)

type MemoryStatus struct {
	// Alloc is bytes of allocated heap objects.
	Alloc uint64 `json:"alloc"`
	// TotalAlloc is cumulative bytes allocated for heap objects.
	TotalAlloc uint64 `json:"total_alloc"`
	// HeapAlloc is bytes of allocated heap objects.
	HeapAlloc uint64 `json:"heap_alloc"`
	// Sys is the total bytes of memory obtained from the OS.
	Sys uint64 `json:"sys"`
	// NumGC is the number of completed GC cycles.
	NumGC uint32 `json:"num_gc"`
}

func GetMemoryStatus() *MemoryStatus {
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)
	return &MemoryStatus{
		Alloc:      m1.Alloc,
		TotalAlloc: m1.TotalAlloc,
		HeapAlloc:  m1.HeapAlloc,
		Sys:        m1.Sys,
		NumGC:      m1.NumGC,
	}
}
