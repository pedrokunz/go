package internal

import (
	"fmt"
	"sync"
	"time"
)

type Sensor struct {
	ID     int
	DataCh chan int
	mu     sync.Mutex
}

func NewSensor(id int) *Sensor {
	return &Sensor{
		ID:     id,
		DataCh: make(chan int),
	}
}

func (s *Sensor) GenerateData() {
	for i := 1; i <= 5; i++ {
		s.mu.Lock()

		data := i * 10

		fmt.Printf("Sensor %d: Generated data %d\n", s.ID, data)

		s.DataCh <- data

		s.mu.Unlock()

		time.Sleep(1 * time.Second)
	}

	close(s.DataCh)
}
