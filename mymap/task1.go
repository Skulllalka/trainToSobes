package mymap

import (
	"fmt"
	"sync"
)

type ConcurrentMap struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		data: make(map[string]string),
	}
}

func (m *ConcurrentMap) GetOrCreate(key, value string) string {
	m.mu.RLock()
	retValue, ok := m.data[key]
	m.mu.RUnlock()

	if ok {
		return retValue
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] =value
	return m.data[key]
}

func Task1() {
	cm := NewConcurrentMap()

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		val := cm.GetOrCreate("key1", "value1")
		fmt.Println("GR 1 got ", val)
	}()

	go func() {
		defer wg.Done()
		val := cm.GetOrCreate("key1", "value2")
		fmt.Println("GR 2 got ", val)
	}()

	wg.Wait()
}
