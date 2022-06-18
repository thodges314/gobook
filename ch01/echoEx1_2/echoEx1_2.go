//exercise 1.2 print indecies
package main

import (
	"fmt"
	"os"
)

func main() {
	// for i:=1; i<len(os.Args); i++ {
	// 	fmt.Println(i, ":",  os.Args[i])
	// }
	for idx, arg := range os.Args[1:] {
		fmt.Println(idx,":", arg)
	}
}