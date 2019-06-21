package rate_counter

import (
	"testing"
	"time"
)

func TestRateCounter(t *testing.T) {
	interval := 5 * time.Second
	r := NewIncrementCounter(interval, 5)

	check := func(expected int64) {
		val := r.Rate()
		if val != expected {
			t.Error("Expected ", val, " to equal ", expected)
		}
	}

	check(0)
	r.Incr(1)
	check(1)
	time.Sleep(time.Second)
	r.Incr(2)
	check(3)
	time.Sleep(6 * time.Second)
	r.Incr(10)
	check(10)
	time.Sleep(5 * time.Second)
	check(0)
}

func TestRateCounterResetAndRestart(t *testing.T) {
	interval := 100 * time.Millisecond

	r := NewIncrementCounter(interval, 5)

	check := func(expected int64) {
		val := r.Rate()
		if val != expected {
			t.Error("Expected ", val, " to equal ", expected)
		}
	}

	check(0)
	r.Incr(1)
	check(1)
	time.Sleep(2 * interval)
	check(0)
	time.Sleep(2 * interval)
	r.Incr(2)
	check(2)
	time.Sleep(2 * interval)
	check(0)
	r.Incr(2)
	check(2)
}
