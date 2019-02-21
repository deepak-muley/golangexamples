package kvstore

import (
	"fmt"
	"sync"
	"testing"
)

func Test_set_get_positive(t *testing.T) {
	kvs := NewKVStore()
	kvs.Set(10, "test 1 data")
	val, err := kvs.Get(10)
	if err != nil {
		t.Fatalf("Failed to get key: %s", err.Error())
	}
	fmt.Println(val)
}

func Test_set_get_negative(t *testing.T) {
	kvs := NewKVStore()
	kvs.Set(10, "test 1 data")
	val, err := kvs.Get(20)
	if err == nil {
		t.Fatalf("Test failed as it did not error out on wrong key")
	}
	t.Log(val)
}

func Test_set_get_goroutine(t *testing.T) {
	kvs := NewKVStore()
	key := 10
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		kvs.Set(10, "test 1 data")
	}()
	wg.Wait()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			val, err := kvs.Get(key)
			if err != nil {
				t.Fatalf("Failed to get key: %s", err.Error())
			}
			t.Log("Got the key:value as ", key, val)
		}()
	}

	wg.Wait()
}
