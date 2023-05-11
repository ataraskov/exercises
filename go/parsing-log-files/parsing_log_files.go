package parsinglogfiles

import (
	"fmt"
	"regexp"
)

/* Predefine regular expressions to avoid multiple compilations */
var reValidLog = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
var reSplitLog = regexp.MustCompile(`<[~*=-]*>`)
var reQuotedPassword = regexp.MustCompile(`(?i)".*password.*"`)
var reEndOfLineText = regexp.MustCompile(`end-of-line\d+`)
var reUserName = regexp.MustCompile(`User\s+(\S+)`)

func IsValidLine(text string) bool {
	return reValidLog.MatchString(text)
}

func SplitLogLine(text string) []string {
	return reSplitLog.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	for _, line := range lines {
		if reQuotedPassword.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	return reEndOfLineText.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	var result []string
	for _, line := range lines {
		submatch := reUserName.FindStringSubmatch(line)
		if len(submatch) > 0 {
			line = fmt.Sprintf("[USR] %s %s", submatch[1], line)
		}
		result = append(result, line)
	}
	return result
}
