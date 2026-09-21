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

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	ans, tempAns := 0, 0
	for range n {
		line, _ = reader.ReadString('\n')
		current, _ := strconv.Atoi(strings.TrimSpace(line))
		if current == 1 {
			tempAns++
		} else {
			ans = max(ans, tempAns)
			tempAns = 0
		}
	}

	ans = max(ans, tempAns)

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
