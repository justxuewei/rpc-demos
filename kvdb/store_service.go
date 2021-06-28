package kvdb

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type StoreService struct {
	m      map[string]string
	filter map[string]func(key string)
	mu     sync.Mutex
}

func NewStoreService() *StoreService {
	return &StoreService{
		m:      make(map[string]string),
		filter: make(map[string]func(key string)),
	}
}

func (s *StoreService) Get(key string, value *string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if v, ok := s.m[key]; ok {
		*value = v
		return nil
	}

	return fmt.Errorf("not found")
}

func (s *StoreService) Set(kv [2]string, reply *struct{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	key, value := kv[0], kv[1]

	if oldValue := s.m[key]; oldValue != value {
		for _, fn := range s.filter {
			fn(key)
		}
	}

	s.m[key] = value
	return nil
}

func (s *StoreService) Watch(timeoutSecond int, keyChange *string) error {
	id := fmt.Sprintf("watch-%s-%03d", time.Now(), rand.Int())
	ch := make(chan string, 10) // buffered

	s.mu.Lock()
	s.filter[id] = func(key string) {
		ch <- key
	}
	s.mu.Unlock()

	select {
	case <-time.After(time.Duration(timeoutSecond) * time.Second):
		return fmt.Errorf("timeout")
	case key := <-ch:
		*keyChange = key
		return nil
	}
}


