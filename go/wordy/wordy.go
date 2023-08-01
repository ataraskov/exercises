// Package wordy provides a solution to the Wordy problem.
package wordy

import (
	"strconv"
	"strings"
)

var (
	PLUS     = "plus"
	MINUS    = "minus"
	DIVID    = "divided"
	MULTIPLY = "multiplied"
	BY       = "by"
)

// Answer returns the result of a Wordy arithmetic question.
//
// TODO: Better approach would be to parse:
// * all the numbers into an array
// * all the operations in another array
// That would simplify logic and improve readability.
//
// But current implementation is a good enough for now.
func Answer(question string) (int, bool) {
	// Cleanup question first
	question = strings.TrimPrefix(question, "What is ")
	question = strings.TrimSuffix(question, "?")
	question = strings.ReplaceAll(question, MULTIPLY+" "+BY, MULTIPLY)
	question = strings.ReplaceAll(question, DIVID+" "+BY, DIVID)

	res := 0
	operand := true // wating for an operand or operation
	operation := "" // operation to be applied to the next operand

	ops := strings.Split(question, " ")
	for i := 0; i < len(ops); i++ {
		if operand {
			// parse number
			num, err := strconv.Atoi(ops[i])
			if err != nil {
				return 0, false
			}
			operand = false

			// apply operation if any
			switch operation {
			case MINUS:
				res -= num
			case PLUS:
				res += num
			case DIVID:
				res /= num
			case MULTIPLY:
				res *= num
			case "":
				res = num
			default:
				return 0, false
			}
			operation = ""

		} else {
			// take operation
			operation = ops[i]
			operand = true
		}
	}

	// missing operand
	if operand {
		return 0, false
	}

	return res, true
}
