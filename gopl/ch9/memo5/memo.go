package memo

// Package memo provides a concurrency-unsafe
// memoization of a function of type Fucn.

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

// A request is a message requesting that the Func be applied to key.
type request struct {
	key      string
	response chan<- result // the client wants a single result
}

// A Memo caches the results of calling a Func.
type Memo struct {
	requests chan request
}

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

// A result is the result of calling a Func.
type result struct {
	value interface{}
	err   error
}

// New returns a memoization of f.  Client must subsequently call Close.
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	// start the monitor goroutine.
	go memo.server(f)
	return memo
}

// Get returns cached value of key.
func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	// sent request to the monitor goroutine, then immediately receives from it.
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

// Close release the requests channel.
func (memo *Memo) Close() {
	close(memo.requests)
}

func (memo *Memo) server(f Func) {
	// the cache variable is confined to the monitor goroutine.
	cache := make(map[string]*entry)

	// reads request in a loop until the request channel is closed by
	// the Close() method
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// This is the first request for this key.
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key) // call f(key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	// Evaluate the function.
	e.res.value, e.res.err = f(key)
	// Broadcast the ready condition.
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// Wait for the ready condition.
	<-e.ready
	// Send the result to the client.
	response <- e.res
}
