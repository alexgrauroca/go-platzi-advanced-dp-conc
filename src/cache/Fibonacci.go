package cache

func Fibonacci(cache *Memory, n int) int {
	if n <= 1 {
		return n
	}

	fb1, _ := GetFibonacci(cache, n-1)
	fb2, _ := GetFibonacci(cache, n-2)

	return fb1.(int) + fb2.(int)
}

func GetFibonacci(cache *Memory, n int) (interface{}, error) {
	value, err := cache.Get(n)

	if err != nil {
		return nil, err
	}

	return value, nil
}
