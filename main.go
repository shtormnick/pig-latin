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
	for i, _ := range words {
		w1 := words[i]
		s := w1[len(w1)-1:]
		if s == "," || s == "." || s == "-"  || s == "?" || s == "!" || s == " " || s == ":" || s == "..." {
			words[i] = w1[1:len(w1)-1] + w1[:1] + pig + s
		} else {
			words[i] = w1[1:] + w1[:1] + pig
		}
		if w1[0] ==  
	}
	pigwords := strings.Join(words, " ")
	return pigwords
}

func main() {
	word := []string(Splitter("I"))
	fmt.Printf(ToPig(word))
}
