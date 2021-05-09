package consumerAndProvider

import "fmt"

type Task struct {
	begin int
	end int
	result chan <- int
}

func chain(in chan int) chan int  {
	out := make(chan int)
	go func() {
		for i := range in {
		out <-1+i

		}
		close(out)
	}()
	return out
}

func RunChain()  {
	in :=make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			in<-i
		}
		close(in)
	}()
	ints := chain(chain(chain(in)))
	for v := range ints {
		fmt.Println(v)
	}
}