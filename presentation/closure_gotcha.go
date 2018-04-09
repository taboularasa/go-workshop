package main

import (
	"fmt"
	"time"
)

func main() {
	values := []string{"this", "is", "a", "test"}
	for _, val := range values {
		go func() {
			fmt.Println(val)
		}()
	}
	time.Sleep(1 * time.Second)
}
