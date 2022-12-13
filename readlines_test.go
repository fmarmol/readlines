package readlines

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadLines(t *testing.T) {
	text :=
		`one
two
three`
	buff := bytes.NewBufferString(text)
	count := 0
	for range ReadLines(buff, nil) {
		count++
	}
	assert.Equal(t, count, 3)
}
func TestReadLinesNumber(t *testing.T) {
	text :=
		`one
two
three`
	buff := bytes.NewBufferString(text)
	count := 0
	for range ReadLines(buff, &Config{Number: 1}) {
		count++
	}
	assert.Equal(t, count, 1)
}

func TestReadLinesGreaterNumber(t *testing.T) {
	text :=
		`one
two
three`
	buff := bytes.NewBufferString(text)
	count := 0
	for range ReadLines(buff, &Config{Number: 14}) {
		count++
	}
	assert.Equal(t, count, 3)
}

func TestReadLinesZeroNumber(t *testing.T) {
	text :=
		`one
two
three`
	buff := bytes.NewBufferString(text)
	count := 0
	for range ReadLines(buff, &Config{Number: 0}) {
		count++
	}
	assert.Equal(t, count, 3)
}
func TestCapitalize(t *testing.T) {
	text :=
		`one
two
three`
	buff := bytes.NewBufferString(text)
	count := 0
	for line := range ReadLines(buff, &Config{ApplyOnEachLine: strings.ToUpper}) {

		count++
		switch count {
		case 1:
			assert.Equal(t, line, "ONE")
		case 2:
			assert.Equal(t, line, "TWO")
		case 3:
			assert.Equal(t, line, "THREE")
		}
	}
	if count != 3 {
		t.Fatalf("expected 3 lines got %v", count)
	}
}
