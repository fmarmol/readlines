package readlines

import (
	"bufio"
	"io"
	"os"
)

type Config struct {
	Number          uint
	ApplyOnEachLine func(s string) string
}

func ReadLinesFromPath(path string, config *Config) <-chan string {
	fd, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return ReadLines(fd, config)
}

func ReadLines(reader io.Reader, config *Config) <-chan string {
	ret := make(chan string)
	var count uint = 0
	go func() {
		defer close(ret)
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			line := scanner.Text()
			if config != nil && config.ApplyOnEachLine != nil {
				ret <- config.ApplyOnEachLine(line)
			} else {
				ret <- line
			}
			count++
			if config != nil && config.Number > 0 {
				if count >= config.Number {
					return
				}
			}
		}
	}()
	return ret
}
