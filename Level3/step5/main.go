package main

import (
	"fmt"
	"sort"
	"strings"
)

func getTopWords(wordMap map[string]int, n int) []string {
	topFiveWords := make([]string, 0, len(wordMap))

	for word := range wordMap {
		topFiveWords = append(topFiveWords, word)
	}

	sort.Slice(topFiveWords, func(i, j int) bool {
		return wordMap[topFiveWords[i]] > wordMap[topFiveWords[j]]
	})

	return topFiveWords[:n]
}

func AnalyzeText(text string) {
	words := strings.Split(text, " ")
	unicWords := make(map[string]int)
	for i, word := range words {
		for _, sign := range ".,!?" {
			word = strings.Trim(word, string(sign))
		}
		words[i] = word
		unicWords[strings.ToLower(word)]++
	}

	fmt.Println("Количество слов:", len(words))
	fmt.Println("Количество уникальных слов:", len(unicWords))

	topWords := getTopWords(unicWords, 5)

	fmt.Printf("Самое часто встречающееся слово: \"%s\" (встречается %d раз)\n",
		topWords[0], unicWords[topWords[0]])
	fmt.Println("Топ-5 самых часто встречающихся слов:")
	for _, word := range topWords {
		fmt.Printf("\"%s\": %d раз\n", word, unicWords[word])
	}

}
