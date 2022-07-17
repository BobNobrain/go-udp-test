package server

import (
	"time"
)

func (srv *Server) startBroadcastingUpdates() {
	for {
		time.Sleep(50 * time.Millisecond)

		srv.mutex.RLock()
		if srv.closed {
			return
		}
		srv.mutex.RUnlock()

		srv.BroadcastUpdates()
	}
}
