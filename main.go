package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func main() {
	var body = readTextFile("text.txt")
	fmt.Println(MostCommonWords(body, 10))
}
func MostCommonWords(body string, quantity int) map[string]int {
	var topTenFrequency = make(map[string]int, 0)
	var tempMap = make(map[string]int, 0)
	// remove all characters that are not letters from text
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(body, " ")
	processedString = strings.Trim(processedString, " ")
	processedString = strings.ToLower(processedString)
	if len(processedString) == 0 {
		return topTenFrequency
	}
	// create an array of words
	wordsArr := strings.Split(processedString, " ")
	// create a map of words and their frequencies
	for _, word := range wordsArr {
		if _, ok := topTenFrequency[word]; ok {
			topTenFrequency[word] += 1
			tempMap[word] +=1
		} else {
			topTenFrequency[word] = 1
			tempMap[word] = 1
		}
	}
	// get the top ten
	for len(topTenFrequency) > quantity {
		for word := range tempMap {
			if len(topTenFrequency) == quantity {
				return topTenFrequency
			}
			tempMap[word] -= 1
			if tempMap[word] == 0 {
				delete(topTenFrequency, word)
				delete(tempMap, word)
			}
		}
	}
	return topTenFrequency
}

func readTextFile(url string) string {
	content, err := ioutil.ReadFile(url)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}