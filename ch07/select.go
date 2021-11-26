package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

func gen(min, max int, creatNumber chan int, end chan bool) {
	time.Sleep(time.Second)
	for {
		select {
		case creatNumber <- rand.Intn(max-min) + min:
		case <-end:
			fmt.Println("Ended!")
			// return
		case <-time.After(time.Second * 4):
			fmt.Println("time.After()!")
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	creatNumber := make(chan int)
	end := make(chan bool)

	if len(os.Args) != 2 {
		fmt.Println("Please give me an integer")
		return
	}

	n, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Going to create %d random numbers.\n", n)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		gen(0, 2*n, creatNumber, end)
		wg.Done()
	}()

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", <-creatNumber)
	}

	end <- true
	wg.Wait()
	fmt.Println("Exiting...")
}
