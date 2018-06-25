package coordinator

import "sync"

var wg sync.WaitGroup

// Run Execute a function in a goroutine without having to muck with a
// WaitGroup. Call Wait() to wait.
func Run(f func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		f()
	}()
}

// Wait ok
func Wait() {
	wg.Wait()
}
