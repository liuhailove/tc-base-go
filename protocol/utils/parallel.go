package utils

import (
	"go.uber.org/atomic"
	"runtime"
	"sync"
)

// ParallelExec 将使用 vals 的每个元素执行给定函数，如果 len(vals) >= parallelThreshold，
// 将按照给定的步长并行执行它们。所以fn必须是线程安全的。
func ParallelExec[T any](vals []T, parallelThreshold, step uint64, fn func(T)) {
	if uint64(len(vals)) < parallelThreshold {
		for _, v := range vals {
			fn(v)
		}
		return
	}

	// 并行 - 实现更高效的多核利用
	start := atomic.NewUint64(0)
	end := uint64(len(vals))

	var wg sync.WaitGroup
	numCPU := runtime.NumCPU()
	wg.Add(numCPU)
	for p := 0; p < numCPU; p++ {
		go func() {
			defer wg.Done()
			for {
				n := start.Add(step)
				if n >= end+step {
					return
				}

				for i := n - step; i < n && i < end; i++ {
					fn(vals[i])
				}
			}
		}()
	}
	wg.Wait()
}
