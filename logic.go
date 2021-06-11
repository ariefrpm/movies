package main

import (
	"log"
	"sort"
	"strings"
)

func main()  {
	list := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	anagram := groupAnagram(list)
	log.Printf("anagram %v\n", anagram)
}

func groupAnagram(list []string) [][]string  {
	//using map to store grouped anagram
	//O(n) space complexity
	anagram := make(map[string][]string)
	for _, w := range list {
		//sort word to be a map key
		//sorting is O(n log n) time complexity
		//todo improve time complexity
		sWord := strings.Split(w,"")
		sort.Strings(sWord)
		word := strings.Join(sWord, "")

		//put word into a map, groped by sorted word
		if _, ok := anagram[word]; ok {
			anagram[word] = append(anagram[word], w)
		} else {
			anagram[word] = []string{w}
		}
	}

	//convert map to 2d array, ignore sorted word
	result := make([][]string, len(anagram))
	i := 0
	for _, v := range anagram {
		result[i] = v
		i++
	}
	return result
}