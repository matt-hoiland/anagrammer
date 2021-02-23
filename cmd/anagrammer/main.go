package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	re "regexp"
	s "strings"
)

const allCapsPattern = `^[A-Z]+$`

func loadWordList() (words []string) {
	fileName := "wordlist.txt"

	fp, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open %s, err: %v", fileName, err)
	}

	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		word := s.ToUpper(s.Split(s.TrimSpace(line), " ")[0])
		if match, _ := re.MatchString(allCapsPattern, word); match {
			words = append(words, word)
		}
	}
	return
}

func validateWordList(words []string) bool {
	for i, word := range words {
		if match, _ := re.MatchString(allCapsPattern, word); !match {
			fmt.Println(i, word)
			return false
		}
	}
	return true
}

func main() {
	words := loadWordList()
	if validateWordList(words) {
		fmt.Println("Only ALLCAPS words here!")
	} else {
		fmt.Println("Something's fishy")
	}
	fmt.Printf("Total words: %d\n", len(words))
}
