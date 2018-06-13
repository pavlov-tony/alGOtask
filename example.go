package main

import (
	"bufio"
	"os"
	"strings"
	"time"

	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/pavlov-tony/alGOtask/search"
	"github.com/pavlov-tony/alGOtask/trie"
)

func main() {
	start := time.Now()
	vocabulary := trie.InitTrie()
	vocabularyFile, err := os.Open("vocabulary.txt")
	defer vocabularyFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(vocabularyFile)
	for scanner.Scan() {
		vocabulary.Insert(strings.ToLower(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	if len(os.Args) < 2 {
		log.Fatal(errors.New("get command line args: The system cannot find the command line arg for input data file name"))
	}
	inputFileName := os.Args[1]
	textFile, err := os.Open(inputFileName)
	defer textFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	text, err := ioutil.ReadAll(textFile)
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Split(string(text), " ")
	ch := make(chan int)
	for _, word := range words {
		go search.SearchDistance(vocabulary, strings.TrimSpace(word), ch)
	}
	result := 0
	for i := 0; i < len(words); i++ {
		result += <-ch
	}
	fmt.Println(result)
	elapsed := time.Since(start)
	log.Printf("Search distance took %s", elapsed)
}
