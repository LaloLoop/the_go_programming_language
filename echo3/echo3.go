// Echo prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var s, sep string

	start := time.Now()

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

	start = time.Now()

	fmt.Println(strings.Join(os.Args[1:], " "))

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}
