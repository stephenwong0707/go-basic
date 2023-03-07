package race

import (
	"sync"
	"sync/atomic"
)

const max = 500

func atomicIncrement(rountineSize int) {
	ch := make(chan bool)
	count := atomic.Int32{}
	for i := 0; i < rountineSize; i++ {
		go func() {
			for {
				c := count.Add(1)
				if c >= max {
					ch <- true
					return
				}
			}
		}()
	}
	<-ch
}

func mutexIncrement(rountineSize int) {
	ch := make(chan bool)
	count := 0
	var mu sync.Mutex
	for i := 0; i < rountineSize; i++ {
		go func() {
			for {
				mu.Lock()
				count++
				mu.Unlock()
				if count >= max {
					ch <- true
					return
				}
			}
		}()
	}
	<-ch
}
