package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const otherWord = "*"

func makeWords() []string {
	return []string{
		otherWord,
		otherWord,
		otherWord,
		otherWord,
		otherWord + "app",
		otherWord + "site",
		otherWord + "time",
		"get" + otherWord,
		"go" + otherWord,
		"lets" + otherWord,
	}
}

func makeWordsFromFile() ([]string, error) {
	var transforms = []string{}

	f, err := os.Open("word.text")
	if err != nil {
		return nil, errors.New("File Not Found")
	}

	s := bufio.NewScanner(f)
	for s.Scan() {
		transforms = append(transforms, otherWord, otherWord+s.Text(), s.Text()+otherWord)
	}
	if err := s.Err(); err != nil {
		return nil, errors.New("Failed to read file")
	}

	defer f.Close()
	return transforms, nil
}

func main() {
	transforms, err := makeWordsFromFile()
	if err != nil {
		transforms = makeWords()
	}

	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := transforms[rand.Intn(len(transforms))]

		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
	}
	if err := s.Err(); err != nil {
		fmt.Printf("err %s", err)
	}
}
