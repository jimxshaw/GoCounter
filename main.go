package main

import (
	"bufio" // Read text.
	"fmt"
	"io" // Provides io.Reader interface.
	"os" // Use OS resources.
)

func main() {
	fmt.Println(count(os.Stdin))
}

func count(r io.Reader) int {
	// This is used to read text from a Reader, such as files.
	scanner := bufio.NewScanner(r)

	// Define the split type, which is words. The default is by lines.
	scanner.Split(bufio.ScanWords)

	// Define a counter.
	wordCounter := 0

	// Increment the counter for every word scanned.
	for scanner.Scan() {
		wordCounter++
	}

	return wordCounter
}
