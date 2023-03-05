package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	wordsFrequency := map[string]int{}
	for _, word := range strings.Fields(text) {
		wordsFrequency[word]++
	}

	words := getWordsOnly(wordsFrequency)
	sort.Slice(words, func(i, j int) bool {
		if wordsFrequency[words[i]] == wordsFrequency[words[j]] {
			return words[i] < words[j]
		}
		return wordsFrequency[words[i]] > wordsFrequency[words[j]]
	})

	if len(words) < 10 {
		return words
	}

	return words[:10]
}

func getWordsOnly(wf map[string]int) []string {
	words := make([]string, 0, len(wf))
	for w := range wf {
		words = append(words, w)
	}
	return words
}
