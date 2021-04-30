// To execute Go code, please declare a func main() in a package "main"

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	count := 0
	resetTime := time.Now().Add(1 * time.Second)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		now := time.Now()
		if now.After(resetTime) {
			count = 0
			resetTime = time.Now().Add(1 * time.Second)
		}
		count += 1
		if count < 5 {
			fmt.Print(str)
		}
	}
}
