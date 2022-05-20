package cache

type Function func(cache *Memory, key int) (interface{}, error)
