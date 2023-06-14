package lib

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const DATA_PATH = "/home/babraham/projects/personal/advent-of-code-22/data/"

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// ReadLines reads a whole file into memory
// and returns a slice of its lines.
func ReadLines(path string) ([]string, error) {
	fullPath := DATA_PATH + path
	file, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func StrToInt(str string) int {
	s := strings.Trim(str, " ")
	i, err := strconv.Atoi(s)
	HandleError(err)

	return i
}
