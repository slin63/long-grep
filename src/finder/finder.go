package finder

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func Find(filename string, expression string) string {
	var results string
	file, err := os.Open(filename)
	log.Println("Searching in: ", filename, expression)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lineNumber := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// check if line matches pattern
		text := scanner.Text()
		matched, _ := regexp.MatchString(expression, text)
		if matched {
			result := fmt.Sprintf("%s | %d: %s\n", filename, lineNumber, text)
			results += result
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return results
}
