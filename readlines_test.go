package readlines

import (
	"bytes"
	"testing"
)

func TestReadLines(t *testing.T) {
	text :=
		`one
two
three`
	buff := bytes.NewBufferString(text)
	count := 0
	for range ReadLines(buff) {
		count++
	}
	if count != 3 {
		t.Fatalf("expected 3 lines got %v", count)
	}
}
