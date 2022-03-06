package scheduler

import "goCrawler/engine"

type GoScheduler struct {
	workerChan chan engine.Request
}

func (s *GoScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *GoScheduler) WorkerReady(w chan engine.Request) {
}

func (s *GoScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *GoScheduler) Submit(r engine.Request) {
	// send requset down to worker chan
	go func() {
		s.workerChan <- r
	}()
}
