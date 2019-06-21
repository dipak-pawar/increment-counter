package rate_counter

import "sync/atomic"

type Counter int64

func (c *Counter) increment(val int64) {
	atomic.AddInt64((*int64)(c), val)
}

func (c *Counter) Value() int64 {
	return atomic.LoadInt64((*int64)(c))
}

func (c *Counter) Reset() {
	atomic.StoreInt64((*int64)(c), 0)
}


