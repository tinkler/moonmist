package sub

import "sync"

type PubSub[T any] struct {
	subscribers map[chan<- T]struct{}
	mu          sync.Mutex
}

func NewPubSub[T any]() *PubSub[T] {
	return &PubSub[T]{
		subscribers: make(map[chan<- T]struct{}),
	}
}

func (ps *PubSub[T]) Subscribe() <-chan T {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ch := make(chan T)
	ps.subscribers[ch] = struct{}{}

	return ch
}

func (ps *PubSub[T]) Publish(msg T) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	for ch := range ps.subscribers {
		ch <- msg
	}
}
