// Package grep provides a colution for exercise Grep on Exercism's Go Track.
package grep

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// Search returns maching lines for given files acoording to pattern and flags
func Search(pattern string, flags, files []string) []string {
	matches := []string{}
	lineNumberes := false
	fileNamesOnly := false
	invertMatch := false
	filesCount := len(files)

	for _, flag := range flags {
		switch flag {
		case "-n":
			lineNumberes = true
		case "-l":
			fileNamesOnly = true
		case "-i":
			pattern = `(?i)` + pattern
		case "-v":
			invertMatch = true
		case "-x":
			pattern = `^` + pattern + `$`
		}
	}

	r, err := regexp.Compile(pattern)
	if err != nil {
		return []string{}
	}

	for _, name := range files {
		file, err := os.Open(name)
		if err != nil {
			fmt.Printf("error opening file %s: %s\n", name, err)
		}
		scanner := bufio.NewScanner(file)
		lineCount := 1
		for scanner.Scan() {
			prefix := ""
			if filesCount > 1 {
				prefix = fmt.Sprintf("%s:", name)
			}

			match := r.MatchString(scanner.Text())
			if invertMatch {
				match = !match
			}

			if match {
				if lineNumberes {
					prefix = fmt.Sprintf("%s%d:", prefix, lineCount)
				}
				if fileNamesOnly {
					matches = append(matches, name)
					break
				}
				matches = append(matches, fmt.Sprintf("%s%s", prefix, scanner.Text()))
			}
			lineCount++
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("error scanning line by line: %s\n", err)
		}
	}

	return matches
}
