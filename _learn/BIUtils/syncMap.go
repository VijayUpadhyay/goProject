package main

import "sync"

type SyncMap struct {
	lock sync.RWMutex
	KV   map[string]interface{}
}

func (m *SyncMap) Store(k string, v interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.KV[k] = v
}

func (m *SyncMap) Load(k string) interface{} {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.KV[k]
}

func (m *SyncMap) Remove(k string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.KV, k)
}

func (m *SyncMap) Size() int {
	m.lock.Lock()
	defer m.lock.Unlock()
	return len(m.KV)
}
