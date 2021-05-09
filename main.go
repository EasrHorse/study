package main

import (
	"fmt"
	"sync"
	"time"
)



func MyTicker(){
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	time.Sleep(time.Millisecond * 2100)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

func MyMutx() {
	var mutex sync.Mutex
	count := 0

	for r := 0; r < 50; r++ {
		go func() {
			mutex.Lock()
			count += 1
			mutex.Unlock()
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("the count is : ", count)

}

var b int
	func main()  {



		fmt.Println(b)
//consumerAndProvider.Run()
//	consumerAndProvider.Run2()
//MyMutx()
		//MyTicker()

		time.Sleep(2*time.Second)
	}





