package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	j, _ := reader.ReadString('\n')
	j = strings.TrimSpace(j)

	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)

	special := make(map[rune]bool)

	for _, r := range j {
		special[r] = true
	}

	var n int

	for _, r := range s {
		_, ok := special[r]
		if ok {
			n++
		}
	}

	writer.WriteString(strconv.Itoa(n))
	writer.WriteByte('\n')
}
