package cache

type Memory struct {
	f     Function
	cache map[int]FunctionResult
}

func (m *Memory) Get(key int) (interface{}, error) {
	result, exists := m.cache[key]

	if !exists {
		result.value, result.err = m.f(m, key)
		m.cache[key] = result
	}

	return result.value, result.err
}
