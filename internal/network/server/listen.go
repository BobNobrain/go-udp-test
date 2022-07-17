package server

import (
	"goudptest/internal/network"
	"goudptest/internal/util/binpack"
)

func (srv *Server) listen() {
	buffer := make([]byte, 128)

	for {
		srv.mutex.RLock()
		if srv.closed {
			return
		}
		srv.mutex.RUnlock()

		n, addr, netErr := srv.connection.ReadFromUDP(buffer)

		if netErr != nil {
			srv.mutex.RLock()
			srv.onMsg(netErr.Error())
			srv.mutex.RUnlock()
			continue
		}

		bp := binpack.NewBinpackerAround(buffer[:n])

		datagram, dgErr := network.UnmarshallDatagram(bp)

		if dgErr != nil {
			srv.mutex.RLock()
			srv.onMsg(dgErr.Error())
			srv.mutex.RUnlock()
			continue
		}

		clientMessage := createClientMessage(*datagram)

		srv.mutex.RLock()
		client := srv.clients[addr.String()]
		srv.mutex.RUnlock()
		if client == nil {
			client = &RemoteClient{
				addr: addr,
			}
		}

		applyErr := clientMessage.Apply(srv, client)

		if applyErr != nil {
			srv.mutex.RLock()
			srv.onMsg(applyErr.Error())
			srv.mutex.RUnlock()
			continue
		}
	}
}
