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
				if !checkStringInSlice(word, ngrams) {
					ngrams = append(ngrams, word)
				}
				continue
			}
			for i := min; i <= max; i++ {
				for j := 0; j < len(word)-i+1; j++ {
					cut := word[j : j+i]
					if !checkStringInSlice(cut, ngrams) {
						ngrams = append(ngrams, cut)
					}
				}
			}
		}
	}
	return ngrams
}

// create a function to check specific string in slice
func checkStringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func main() {
	x := []string{"hello world", "world"}
	fmt.Println(Ngram(x, 2, 9))
}
