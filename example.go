package main

import (
	"bytes"
	"os"
	"runtime/debug"
	"time"

	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/pavlov-tony/alGOtask/search"
	"github.com/pavlov-tony/alGOtask/trie"
)

func main() {
	// little hack to switch off Garbage Collector and win in performance
	debug.SetGCPercent(-1)
	start := time.Now()
	vocabulary := trie.Init()
	vocabularyText, err := ioutil.ReadFile("vocabulary.txt")
	if err != nil {
		log.Fatal(err)
	}
	vocabularyWords := bytes.Split(vocabularyText, []byte{10})
	for _, word := range vocabularyWords {
		vocabulary.Insert(word)
	}
	if len(os.Args) < 2 {
		log.Fatal(errors.New("get command line args: the system cannot find the command line arg for input data file name"))
	}
	inputFileName := os.Args[1]
	text, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	textWords := bytes.Split(text, []byte{32})
	ch := make(chan int)
	for _, word := range textWords {
		go search.Distance(vocabulary, bytes.ToUpper(word), ch)
	}
	result := 0
	for i := 0; i < len(textWords); i++ {
		result += <-ch
	}
	fmt.Println(result)
	elapsed := time.Since(start)
	log.Printf("Search distance took %s", elapsed)
}
