// echo1 prints it's command line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string //variables s and sep of type string - initialises to ""
	// fmt.Println(os.Args[0])
	for i := 1; i < len(os.Args); i++ {//:= 'short variable declaration' - declared variables and automatically sets type
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}