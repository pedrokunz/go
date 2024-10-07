package internal

import (
	"fmt"
	"sync"
)

type Processor struct {
	ID       int
	DataCh   chan int
	ResultCh chan int
	mu       sync.Mutex
}

func NewProcessor(id int) *Processor {
	return &Processor{
		ID:       id,
		DataCh:   make(chan int),
		ResultCh: make(chan int),
	}
}

func (p *Processor) ProcessData() {
	for data := range p.DataCh {
		p.mu.Lock()

		result := data * 2

		fmt.Printf("Processor %d: Processed data %d\n", p.ID, result)

		p.ResultCh <- result

		p.mu.Unlock()
	}

	close(p.ResultCh)
}
