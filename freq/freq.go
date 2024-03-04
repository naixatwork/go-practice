package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

// Q:What is the most common word in sherlock.txt

// "Who's on first?" -> [Who s on first]
var wordRegex = regexp.MustCompile(`[a-zA-Z]+`)

func mapDemo() {
	var stocks map[string]int
	sym := "TTWO"
	price := stocks
	fmt.Printf("%s -> ")
}

func wordFrequency(r io.Reader) (map[string]int, error) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		words := wordRegex.FindAllString(s.Text(), -1)
		if len(words) != 0 {
			fmt.Println(words)
			break
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}

	return nil, nil
}

func main() {
	file, err := os.Open("sherlock.txt")
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer file.Close()

	wordFrequency(file)

}
