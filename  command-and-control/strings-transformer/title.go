package stringstransformer

import "strings"

func Title(s string) string {

	text := []string{"a", "an", "the", "and", "but", "or", "for", "nor", "on", "at", "to", "by", "in", "of", "up", "as", "is", "it "}

	word := strings.Fields(s)
	for i, w := range word {
		myflag := false
		for _, w1 := range text {
			if w == w1 {
				myflag = true
				break
			}

		}
		if myflag == false {
			word[i] = strings.ToUpper(string(word[i][0])) + strings.ToLower(word[i][1:])
		}
	}
	return strings.Join(word, " ")
}
