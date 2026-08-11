package snippets

import "slices"

func convertToRunes(words []string) [][]rune {
	var wordRunes [][]rune
	for _, word := range words {
		r := []rune(word)
		slices.Sort(r)
		wordRunes = append(wordRunes, r)
	}
	return wordRunes
}

func GroupAnagrams(words []string) [][]string {
	var groups [][]string
	mapped := map[string][]string{} // writing to a nil map is forbidden , reading is fine
	sortedWordRunes := convertToRunes(words)
	for index, word := range sortedWordRunes {
		mapped[string(word)] = append(mapped[string(word)], words[index])
	}
	for _, val := range mapped {
		groups = append(groups, val)
	}
	return groups
}
