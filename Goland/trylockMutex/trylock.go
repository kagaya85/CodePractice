package trylockmutex

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

const (
	mutexLocked      = 1 << iota // 加锁标识位置
	mutexWoken                   // 唤醒标识位置
	mutexStarving                // 锁饥饿标识位置
	mutexWaiterShift = iota      // 标识waiter的起始bit位置
)

// Mutex 扩展了TryLock的Mutex
type Mutex struct {
	sync.Mutex
}

// TryLock 非阻塞加锁，加锁成功返回true，否则返回false
func (m *Mutex) TryLock() bool {

	// fast path
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked) {
		return true
	}

	old := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	if old&(mutexLocked|mutexStarving|mutexWoken) != 0 { // 如果mutex处于唤醒、加锁、饥饿状态，则直接返回false
		return false
	}

	new := old | mutexLocked
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), old, new) // 尝试竞争锁
}
