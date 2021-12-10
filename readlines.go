package readlines

import (
	"bufio"
	"io"
)

func ReadLines(reader io.Reader) <-chan string {
	ret := make(chan string)
	go func() {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			ret <- scanner.Text()
		}
		close(ret)
	}()
	return ret
}
