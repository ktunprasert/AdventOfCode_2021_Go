package helper

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadFile(filePath string) (lines []int64) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		lines = append(lines, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}
