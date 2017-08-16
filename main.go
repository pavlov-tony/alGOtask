package main

import (
	"bufio"
	"os"
	"strings"

	//"time"

	"errors"
	"fmt"
	"github.com/pavlov-tony/alGOtask/trie"
	"io/ioutil"
	"log"
)

func main() {
	//start := time.Now()
	vocabulary := trie.InitTrie()

	file, err := os.Open("vocabulary.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		vocabulary.Insert(strings.ToLower(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	file.Close()

	if len(os.Args) < 2 {
		log.Fatal(errors.New("get command line args: The system cannot find the command line arg for input data file name."))
	}

	inputFileName := os.Args[1]

	file, err = os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	text, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	words := strings.Split(string(text), " ")

	ch := make(chan int)

	for _, word := range words {
		go vocabulary.SearchDistance(strings.TrimSpace(word), ch)
	}

	result := 0

	for i := 0; i < len(words); i++ {
		result += <-ch
	}

	fmt.Println(result)
	//elapsed := time.Since(start)
	//log.Printf("Search distance took %s", elapsed)
}