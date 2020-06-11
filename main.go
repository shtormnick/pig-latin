package main

import (
	"fmt"
	"strings"
)

func Splitter(s string) []string {
	words := strings.Fields(s)
	return words
}

func ToPig(words []string) string {
	pig := "ay"

	split := map[int]string{
		1: ",",
		2: ".",
		3: "-",
		4: "?",
		5: "!",
		6: " ",
		7: ":",
		8: "...",
	}

	for i, _ := range words {
		w1 := words[i]
		s := w1[len(w1)-1:]
		for k := range split {
			rslt := s == split[k]
			if rslt {
				words[i] = w1[1:len(w1)-1] + w1[:1] + pig + s
			} else {
				words[i] = w1[1:] + w1[:1] + pig
			}
		}

	}
	pigwords := strings.Join(words, " ")
	return pigwords
}

func main() {
	word := []string(Splitter("hello world"))
	fmt.Printf(ToPig(word))
}
