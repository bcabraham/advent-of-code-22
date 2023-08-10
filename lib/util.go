package lib

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func joinPath(trim bool, parts ...string) string {
	path := ""
	for _, p := range parts {
		if !strings.HasSuffix(p, string(os.PathSeparator)) {
			p = p + string(os.PathSeparator)
		}

		path += p
	}

	if trim {
		path = strings.TrimSuffix(path, string(os.PathSeparator))
	}

	return path
}

func getModulePath(module string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return joinPath(false, cwd, module), nil
}

// ReadLines reads a whole file into memory
// and returns a slice of its lines.
func ReadLines(module string, filename string) ([]string, error) {
	path, err := getModulePath(module)
	if err != nil {
		return nil, err
	}

	fullPath := joinPath(true, path, filename)
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

func StrToIntArray(s string) []int {
	data := strings.Split(s, ",")
	arr := []int{}

	for _, d := range data {
		i := StrToInt(d)
		arr = append(arr, i)
	}

	return arr
}
