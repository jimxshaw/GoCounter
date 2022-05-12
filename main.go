package main

import (
	"bufio" // Read text.
	"flag"  // Create and manage command line flags.
	"fmt"
	"io" // Provides io.Reader interface.
	"os" // Use OS resources.
)

func main() {
	// Define a boolean flag to count lines rather than words.
	// The default value is false, meaning count words by default.
	lines := flag.Bool("l", false, "Count lines")

	bytes := flag.Bool("b", false, "Count bytes")

	// Parse the user provided flag.
	flag.Parse()

	fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(r io.Reader, countLines bool, countBytes bool) int {
	// This is used to read text from a Reader, such as files.
	scanner := bufio.NewScanner(r)

	if countBytes {
		scanner.Split(bufio.ScanBytes)
	} else if !countLines {
		// If the count lines flag is not set then we count by words.
		// Define the split type, which is words. The default is by lines.
		scanner.Split(bufio.ScanWords)
	}

	// Define a counter.
	counter := 0

	// Increment the counter for every word or line scanned.
	for scanner.Scan() {
		counter++
	}

	return counter
}
