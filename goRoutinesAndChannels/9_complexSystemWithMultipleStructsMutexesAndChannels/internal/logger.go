package internal

import (
	"fmt"
	"sync"
)

type Logger struct {
	ID       int
	ResultCh chan int
	mu       sync.Mutex
}

func NewLogger(id int) *Logger {
	return &Logger{
		ID:       id,
		ResultCh: make(chan int),
	}
}

func (l *Logger) LogResult() {
	for result := range l.ResultCh {
		l.mu.Lock()

		fmt.Printf("Logger %d: Received result %d\n", l.ID, result)

		l.mu.Unlock()
	}
}
