// https://coderun.yandex.ru/selections/first-2023-backend/problems/median-out-of-three
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
	parts := strings.Fields(strings.TrimSpace(line))

	a, _ := strconv.Atoi(parts[0])
	b, _ := strconv.Atoi(parts[1])
	c, _ := strconv.Atoi(parts[2])

	ans := a + b + c - max(max(a, b), c) - min(min(a, b), c)

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
