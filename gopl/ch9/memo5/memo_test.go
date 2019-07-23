package memo_test

import (
	"testing"

	memo "github.com/dantin/go-by-example/gopl/ch9/memo5"
	"github.com/dantin/go-by-example/gopl/ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func TestSequential(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}
