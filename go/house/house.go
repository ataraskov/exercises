package house

import "strings"

var (
	start = "This is"
	end   = "the house that Jack built."
	head  = []string{
		"malt",
		"rat",
		"cat",
		"dog",
		"cow with the crumpled horn",
		"maiden all forlorn",
		"man all tattered and torn",
		"priest all shaven and shorn",
		"rooster that crowed in the morn",
		"farmer sowing his corn",
		"horse and the hound and the horn",
	}
	postmodifier = []string{
		"lay in",
		"ate",
		"killed",
		"worried",
		"tossed",
		"milked",
		"kissed",
		"married",
		"woke",
		"kept",
		"belonged to",
	}
)

func Verse(v int) string {
	res := []string{}
	switch v {
	case 1:
		res = append(res, start+" "+end)
	default:
		for i := v - 2; i >= 0; i-- {
			if i == v-2 {
				res = append(res, start+" the "+head[i])
			}
			if i == 0 {
				res = append(res, "that "+postmodifier[i]+" "+end)
			} else {
				res = append(res, "that "+postmodifier[i]+" the "+head[i-1])
			}
		}
	}
	return strings.Join(res, "\n")
}

func Song() string {
	res := []string{}
	for i := 0; i <= len(head); i++ {
		res = append(res, Verse(i+1))
	}
	return strings.Join(res, "\n\n")
}
