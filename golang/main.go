package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
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

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func main() {
	lines, err := readLines("../message.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	sorted_lines := make([]string, len(lines)+1)

	max_key := 0
	for _, line := range lines {
		strs := strings.Fields(line)
		i, err := strconv.Atoi(strs[0])
		if err != nil {
			// ... handle error
			panic(err)
		}
		sorted_lines[i] = strs[1]
		if i > max_key {
			max_key = i
		}
	}

	key := 0
	i := 1

	for {
		if key >= max_key {
			break
		}
		key = (i + 1) * i / 2
		fmt.Printf("%s ", sorted_lines[key])
		i += 1
	}
}
