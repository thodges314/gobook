// dup1 prints the text of each line that appears more than once
// in the standard input, prceeded by its count

//ctrl-d to interrupt
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// note: ignoring potential errors from Input.Err()
	for line, n:= range counts {
		if n> 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}