package kvstore

import (
	"fmt"
	"sync"
)

type KVStore struct {
	kvStore map[int]string
	l       sync.RWMutex
}

func (kvs *KVStore) Get(key int) (string, error) {
	kvs.l.RLock()
	defer kvs.l.RUnlock()
	if val, ok := kvs.kvStore[key]; ok {
		return val, nil
	}
	return "", fmt.Errorf("key not found: %v", key)
}

func (kvs *KVStore) Set(key int, val string) {
	kvs.l.Lock()
	defer kvs.l.Unlock()
	kvs.kvStore[key] = val
}

func NewKVStore() *KVStore {
	kvs := new(KVStore)
	kvs.l = sync.RWMutex{}
	kvs.kvStore = make(map[int]string)
	return kvs
}
