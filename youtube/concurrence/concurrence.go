package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex

const goroutinesAmount = 100

var contador int

func main() {
	wg.Add(goroutinesAmount)

	for i := 0; i < goroutinesAmount; i++ {
		go incrementa()
		// fmt.Println(x)
	}

	wg.Wait()

	fmt.Println(contador)

}

func incrementa() {
	mutex.Lock()

	i := contador
	
	runtime.Gosched()
	
	i++

	contador = i

	mutex.Unlock()
	
	wg.Done()
}
