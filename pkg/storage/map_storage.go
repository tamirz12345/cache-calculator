package storage

import "sync"

type CalculatorCache struct {
	Mutex sync.Mutex
	Data  map[string]float64
}

func NewCalculatorCache() *CalculatorCache {
	return &CalculatorCache{
		Data: make(map[string]float64),
	}
}

func (c *CalculatorCache) Get(key string) (float64, bool) {
	c.Mutex.Lock()
	val, ok := c.Data[key]
	c.Mutex.Unlock()
	return val, ok
}

func (c *CalculatorCache) Set(key string, value float64) {
	c.Data[key] = value
}

func (c *CalculatorCache) GetSize() int {
	return len(c.Data)
}
