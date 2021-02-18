package trylockmutex

import (
	"math/rand"
	"testing"
	"time"
)

// "go.testFlags": [ "-v" ]
func TestTryLock(t *testing.T) {
	var mu Mutex
	go func() {
		mu.Lock()
		t.Log("goroutine get the mutex lock")
		time.Sleep(time.Duration((2 + rand.Intn(5))) * time.Second)
		t.Log("goroutine release the mutex lock")
		mu.Unlock()
	}()

	time.Sleep(time.Second)
	ok := mu.TryLock()
	for !ok {
		t.Log("Can't get the mutex lock, sleep for 1 second")
		time.Sleep(time.Second)
		ok = mu.TryLock()
	}

	t.Log("Get the mutex lock!")
	mu.Unlock()
	return
}
