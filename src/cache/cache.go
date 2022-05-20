package cache

import (
	"fmt"
	"time"
)

func Run() {
	fibo := []int{150, 42, 40, 41, 42, 38, 50, 20, 60, 100, 122}
	startRun := time.Now()
	cache := NewCache(func(cache *Memory, n int) (interface{}, error) {
		return Fibonacci(cache, n), nil
	})

	for i := 0; i < len(fibo); i++ {
		n := fibo[i]
		start := time.Now()
		value, err := GetFibonacci(cache, n)

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Printf("%d, %s, %d\n", n, time.Since(start), value)
	}

	fmt.Printf("Process completed in %s\n", time.Since(startRun))
}

func NewCache(f Function) *Memory {
	return &Memory{
		f:     f,
		cache: make(map[int]FunctionResult),
	}
}
