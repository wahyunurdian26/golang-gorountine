package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	var x = 0
	var mutex sync.Mutex

	for i := 1; i < 1000; i++ {
		go func() {
			for j := 1; j < 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("counter =", x)
}
