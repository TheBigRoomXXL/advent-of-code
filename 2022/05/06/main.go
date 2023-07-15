package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var result string
	part := flag.Int("part", 0, "part number (1 or 2)")
	flag.Parse()
	switch *part {
	case 1:
		result = part1(os.Stdin)
	case 2:
		result = part2(os.Stdin)
	default:
		panic("invalid part")
	}
	fmt.Println(result)
}

func part1(input io.Reader) string {
	result := checkForUniqueSeauence(input, 4)
	return fmt.Sprint(result)
}

func part2(input io.Reader) string {
	result := checkForUniqueSeauence(input, 14)
	return fmt.Sprint(result)
}

func checkForUniqueSeauence(input io.Reader, length int) int {
	br := bufio.NewReader(input)
	result := 0
	queue := []byte{}
	queueMap := map[byte]bool{}

	for {
		b, err := br.ReadByte()
		result++

		if err != nil {
			fmt.Println(err)
			break
		}
		if len(queue) < length {
			queue = append(queue, b)
		} else {
			queue = append(queue, b)
			queue = queue[1:]
			queueMap = map[byte]bool{queue[0]: true, queue[1]: true, queue[2]: true, queue[3]: true}
			for _, b := range queue {
				queueMap[b] = true
			}
			if len(queueMap) == length {
				break
			}
		}
	}
	return result
}
