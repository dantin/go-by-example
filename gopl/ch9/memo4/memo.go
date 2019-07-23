package memo

import "sync"

// Package memo provides a concurrency-unsafe
// memoization of a function of type Fucn.

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

// A Memo caches the results of calling a Func.
type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// New returns a Memo that caches the results of calling a Func.
func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

// Get returns cached value of key.
func (memo *Memo) Get(key string) (interface{}, error) {
	// concurrency-safe
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// This is the first request for this key.
		// This goroutine becomes reponsible for computing
		// the value and broadcasting the ready condition.
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)

		// the closing of the ready channel happens-before
		// any other goroutine receives the broadcast event.
		close(e.ready) // broadcast ready condition
	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()

		<-e.ready // wait for ready condition
	}
	return e.res.value, e.res.err
}
