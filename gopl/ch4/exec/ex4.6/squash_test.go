package squash

import "testing"

func TestSquash(t *testing.T) {
	s := []byte("\t测试   空格 \n咯")
	got := string(squash(s))
	want := " 测试 空格 咯"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
