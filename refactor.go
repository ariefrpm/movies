package main

import (
	"log"
	"strings"
)

func main()  {
	str := "aa(abcdef)bb"
	log.Println(findFirstStringInBracketRefactor(str))
	log.Println(findFirstStringInBracket(str))
}

func findFirstStringInBracketRefactor(str string) string {
	if len(str) == 0 {
		return ""
	}

	indexFirstBracketFound := strings.Index(str,"(")
	indexClosingBracketFound := strings.Index(str,")")

	if indexFirstBracketFound < 0 || indexClosingBracketFound < 0 {
		return ""
	}

	return str[indexFirstBracketFound+1:indexClosingBracketFound-1]
}


func findFirstStringInBracket(str string) string {
	if (len(str) > 0) {
		indexFirstBracketFound := strings.Index(str,"(")
		if indexFirstBracketFound >= 0 {
			runes := []rune(str)
			wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket,")")
			if indexClosingBracketFound >= 0 {
				runes := []rune(wordsAfterFirstBracket)
				return string(runes[1:indexClosingBracketFound-1])
			}else{
				return ""
			}
		}else{
			return ""
		}
	}else{
		return ""
	}
	return ""
}
