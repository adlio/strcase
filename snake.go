package strcase

import (
	"unicode"
)

// ToSnake converts the given string to "snake case" (lower case with
// underscores separating words.
//
func ToSnake(str string) string {
	in := []rune(str)

	startsNewWord := func(i int) bool {
		if i < 1 {
			// The beginning of the string isn't a "new" word
			return false
		}
		if !unicode.IsUpper(in[i]) {
			// New words always start with an uppercase letter
			return false
		}
		if unicode.IsLower(in[i-1]) {
			return true
		}
		return i < len(in)-2 && unicode.IsLower(in[i+1])
	}

	out := make([]rune, 0, len(in)+len(in)/2)
	for i, r := range in {
		if i > 0 && startsNewWord(i) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(r))
	}

	return string(out)
}
