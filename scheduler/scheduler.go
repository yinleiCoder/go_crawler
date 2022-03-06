package scheduler

import "goCrawler/engine"

type GoScheduler struct {
	workerChan chan engine.Request
}

func (s *GoScheduler) Submit(r engine.Request) {
	// send requset down to worker chan
	go func() {
		s.workerChan <- r
	}()
}

func (s *GoScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}
