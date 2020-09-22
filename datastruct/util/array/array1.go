package array

import "sync"

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

var (
	data   = make(map[string]string)
	locker sync.RWMutex
)

func WriteToMap(k, v string) {
	// locker.Lock()
	// defer locker.Unlock()
	data[k] = v
}

func ReadFromMap(k string) string {
	// locker.RLock()
	// defer locker.RUnlock()
	return data[k]
}
