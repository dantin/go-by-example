package memo

// Package memo provides a concurrency-unsafe
// memoization of a function of type Fucn.

// A Memo caches the results of calling a Func.
type Memo struct {
	f     Func
	cache map[string]result
}

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// New returns a Memo that caches the results of calling a Func.
func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// Get returns cached value of key.
func (memo *Memo) Get(key string) (interface{}, error) {
	// NOTE: not concurrency-safe!
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}
