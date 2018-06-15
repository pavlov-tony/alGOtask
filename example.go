package main

import (
	"bufio"
	"os"
	"runtime/debug"
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
	// little hack to switch off Garbage Collector and win in performance
	debug.SetGCPercent(-1)
	start := time.Now()
	vocabulary := trie.Init()
	vocabularyFile, err := os.Open("vocabulary.txt")
	defer vocabularyFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(vocabularyFile)
	for scanner.Scan() {
		vocabulary.Insert(scanner.Bytes())
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	if len(os.Args) < 2 {
		log.Fatal(errors.New("get command line args: the system cannot find the command line arg for input data file name"))
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
		go search.Distance(vocabulary, []byte(strings.ToUpper(word)), ch)
	}
	result := 0
	for i := 0; i < len(words); i++ {
		result += <-ch
	}
	fmt.Println(result)
	elapsed := time.Since(start)
	log.Printf("Search distance took %s", elapsed)
}
