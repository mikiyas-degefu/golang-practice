package main


import (
	"fmt"
	"strings"
	"regexp"
)


func wordFrequency(text string) map[string]int {
	text = strings.ToLower(text)

	//remove punctuation
	reg, _ := regexp.Compile(`[^\w\s]`)
	text = reg.ReplaceAllString(text, "")

	//split text	
	words := strings.Fields(text)

	fmt.Println(words)

	//count
	freq := make(map[string]int)
	for _,word := range words{
		freq[word]++
	}

	return freq

}

func main (){
	text := "Go is fun! Go makes learning Go exciting. Fun times with Go!"
	frequencies := wordFrequency(text)
	fmt.Println(frequencies)
}