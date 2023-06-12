package golang_gorutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	x := 0
	mutex := sync.Mutex{} // menghindari race condition ketika melakukan sharing variable x
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock() // start lock
				x = x + 1
				mutex.Unlock() // finish unlock
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}
