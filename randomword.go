//  Basic random line picker
//

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	// Print random line from a file (containing a list of words)
	randomLine("words")

}

func randomLine(path string) {
	var allLines []string
	var randomInt int

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		allLines = append(allLines, scanner.Text())
	}
	f.Close()

	size := len(allLines)
	RandomDigit(&randomInt, size)

	fmt.Println(allLines[randomInt])
}

// From the numbers in the given pool, pseudo-randomly pick one and write it to &mynumber
func RandomDigit(mynumber *int, pool int) {
	rand.Seed(time.Now().UnixNano())
	*mynumber = rand.Intn(pool)
}
