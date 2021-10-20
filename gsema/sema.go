package gsema

import "sync"

type Semaphore struct {
	c  chan bool
	wg *sync.WaitGroup
}

func NewSemaphore(maxSize int) *Semaphore {
	return &Semaphore{
		c:  make(chan bool, maxSize),
		wg: new(sync.WaitGroup),
	}
}

func (s *Semaphore) Add(delta int) {
	s.wg.Add(delta)
	for i := 0; i < delta; i++ {
		s.c <- true
	}
}

func (s *Semaphore) Done() {
	<-s.c
	s.wg.Done()
}

func (s *Semaphore) Wait() {
	s.wg.Wait()
}
