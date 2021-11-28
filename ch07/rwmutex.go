package main

import (
	"fmt"
	"sync"
	"time"
)

var Password *secret
var wg sync.WaitGroup

type secret struct {
	RWM      sync.RWMutex
	password string
}

func Change(pass string) {
	fmt.Println("Change() function")
	Password.RWM.Lock()
	fmt.Println("Change() locked")
	time.Sleep(time.Second * 4)
	Password.password = pass
	Password.RWM.Unlock()
	fmt.Println("Change() unlocked")
}

func show() {
	defer wg.Done()
	Password.RWM.RLock()
	fmt.Println("show() locked")
	time.Sleep(time.Second * 2)
	fmt.Println("Pass value:", Password.password)
	defer Password.RWM.RUnlock()
}

func main() {
	Password = &secret{password: "12345"}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go show()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		Change("54321")
	}()

	wg.Wait()

	// Direct access to Password.password
	fmt.Println("Current password value:", Password.password)
}
