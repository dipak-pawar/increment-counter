package rate_counter

import (
	"sync/atomic"
	"time"
)

type IncrementCounter struct {
	counter    Counter
	interval   time.Duration
	resolution int
	partials   []Counter
	current    int32
	running    int32
}

// resolution is nothing but number of time slots between given amount of time. (e.g. 1 minute has 60 resolutions)
func NewIncrementCounter(duration time.Duration, resolution int) *IncrementCounter {
	counter := &IncrementCounter{
		interval:   duration,
		running:    0,
		resolution: resolution,
		partials:   make([]Counter, resolution),
		current:    0,
	}

	return counter
}

func (r *IncrementCounter) run() {
	if ok := atomic.CompareAndSwapInt32(&r.running, 0, 1); !ok {
		return
	}

	go func() {
		ticker := time.NewTicker(time.Duration(float64(r.interval) / float64(r.resolution)))

		for range ticker.C {
			current := atomic.LoadInt32(&r.current)
			next := (int(current) + 1) % r.resolution
			r.counter.increment(-1 * r.partials[next].Value())
			r.partials[next].Reset()
			atomic.CompareAndSwapInt32(&r.current, current, int32(next))
		}
	}()
}

func (r *IncrementCounter) Incr(val int64) {
	r.counter.increment(val)
	r.partials[atomic.LoadInt32(&r.current)].increment(val)
	r.run()
}

func (r *IncrementCounter) Rate() int64 {
	return r.counter.Value()
}
