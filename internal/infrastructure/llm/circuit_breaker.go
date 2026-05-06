package llm

import (
	"sync"
	"time"
)

type State int

const (
	StateClosed State = iota
	StateOpen
	StateHalfOpen
)

type CircuitBreaker struct {
	mu sync.Mutex

	state State

	failures    int
	maxFailures int

	openedAt time.Time
	timeout  time.Duration
}

func NewCircuitBreaker(maxFailures int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:       StateClosed,
		maxFailures: maxFailures,
		timeout:     timeout,
	}
}

func (cb *CircuitBreaker) nextState() {
	if cb.state == StateOpen && time.Since(cb.openedAt) > cb.timeout {
		cb.state = StateHalfOpen
	}
}

func (cb *CircuitBreaker) Allow() bool {
	cb.nextState()

	switch cb.state {
	case StateOpen:
		return false
	case StateHalfOpen, StateClosed:
		return true
	}
	return false
}

func (cb *CircuitBreaker) Success() {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	cb.failures = 0
	cb.state = StateClosed
}
func (cb *CircuitBreaker) Fail() {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	cb.failures++

	switch cb.state {
	case StateClosed:
		if cb.failures >= cb.maxFailures {
			cb.state = StateOpen
			cb.openedAt = time.Now()
		}

	case StateHalfOpen:
		cb.state = StateOpen
		cb.openedAt = time.Now()
	}
}
