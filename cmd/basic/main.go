package main

import (
	"errors"
	"fmt"
	"time"
)

//var *wg WaitGroup = &sync.WaitGroup

func main() {
	var dividend, divisor int = 10, 0
	factor, reminder, err := divide(dividend, divisor)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Factor = %v Reminder = %v\n", factor, reminder)

	var n int = 100000
	slc := []int{}
	slcSize := make([]int, 0, n)
	fmt.Printf("Time taken without preallocation = %v\n", timeLoop(slc, n))
	fmt.Printf("Time taken with preallocation = %v\n", timeLoop(slcSize, n))

	var ptr *int
	fmt.Println(ptr)

	var channel = make(chan int, 5)

	go produce(channel)
	for c := range channel {
		fmt.Println(c)
		time.Sleep(time.Second * 2)
	}

}

func produce(channel chan int) {
	defer close(channel)
	for i := 0; i < 5; i++ {
		channel <- i
	}
	fmt.Println("Done producing")
}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}

func divide(dividend int, divisor int) (int, int, error) {
	if divisor == 0 {
		var err = errors.New("divide by zero err")
		return 0, 0, err
	}
	return dividend / divisor, dividend % divisor, nil

}
