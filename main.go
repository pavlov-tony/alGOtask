package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"strings"
	"sync"

	"github.com/arbovm/levenshtein"
)

type vocabulary struct {
	lock sync.RWMutex
	v    []string
}

func findDiff(dict *vocabulary, word string, answer chan int) {
	minDiff := math.MaxInt32
	for i := 0; i < 178691; i++ {
		dict.lock.RLock()
		dword := dict.v[i]
		dict.lock.RUnlock()
		diff := levenshtein.Distance(word, dword)
		if diff < minDiff {
			minDiff = diff
		}
		if diff == 0 {
			break
		}
	}
	answer <- minDiff
}

func main() {
	v := &vocabulary{
		v: make([]string, 178691),
	}

	file, err := os.Open("vocabulary.txt")
	if err != nil {
		fmt.Println(err)
	}
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		v.lock.Lock()
		v.v[i] = strings.ToLower(scanner.Text())
		v.lock.Unlock()
		i++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	file.Close()

	sentence := "orem ipsum dolor sit amet consectetur adipiscing elit nteger imperdiet elit et libero commodo et convallis est ultrices raesent faucibus ligula ullamcorper urna pellentesque faucibus liquam ultrices purus sit amet tellus malesuada malesuada hasellus varius faucibus nisl congue placerat mi suscipit vitae ivamus eu lorem mauris a elementum erat nteger a nisl sollicitudin mauris facilisis vehicula quis non erat tiam sit amet porta justo usce eget nisl ipsum am a ante neque egestas rhoncus urna orbi lectus lorem vehicula quis commodo sed scelerisque non diam enean enim quam sollicitudin vel dignissim et feugiat in risus orbi gravida urna in neque sollicitudin elementum nteger ut tortor lacus sed aliquam ipsum usce convallis purus at lobortis accumsan magna odio blandit orci sit amet semper ligula tortor sit amet nisi ellentesque luctus nisi ut placerat dictum massa libero suscipit mi id ullamcorper purus arcu at nunc t ut arcu orci"
	words := strings.Split(sentence, " ")

	answer := make(chan int)

	for _, word := range words {
		go findDiff(v, word, answer)
	}

	resultDiff := 0

	for i := 0; i < len(words); i++ {
		resultDiff += <-answer
	}

	fmt.Println(resultDiff)
}