package gengram

import (
	"fmt"
	"regexp"
	"strings"
)

func Ngram(phrases []string, min int, max int) []string {
	var ngrams []string
	var re = regexp.MustCompile(`/[^a-z0-9\s]/g`)
	for _, phrase := range phrases {
		phrase = strings.ToLower(phrase)
		phrase = re.ReplaceAllString(phrase, "")
		phrase = strings.TrimSpace(phrase)
		words := strings.Split(phrase, " ")
		for _, word := range words {
			if len(word) < min {
				ngrams = append(ngrams, word)
				continue
			}
			for i := min; i <= max; i++ {
				for j := 0; j < len(word)-i+1; j++ {
					cut := word[j : j+i]
					ngrams = append(ngrams, cut)
				}
			}
		}
	}
	return ngrams
}

func main() {
	x := []string{"hello world", "i am finding ngram", "this is a test"}
	fmt.Println(Ngram(x, 2, 9))
}
