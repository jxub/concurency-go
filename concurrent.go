package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// WaitGroup will be used to wait until all goroutines finish
var wg sync.WaitGroup

func main() {
	time.Sleep(5)
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go beep(i)
	}
	wg.Wait()
}

func beep(i int) {
	defer wg.Done()
	time.Sleep(1 * time.Duration(rand.Intn(5)))
	fmt.Println("beep from " + strconv.Itoa(i))
}
