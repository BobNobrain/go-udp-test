package server

import (
	"goudptest/internal/network"
	"goudptest/internal/util/binpack"
)

func (srv *Server) send(client *RemoteClient, message network.Datagram) error {
	bp := binpack.NewBinpacker(128)
	network.MarshallDatagram(bp, message)

	_, err := srv.connection.WriteToUDP(bp.GetData(), client.addr)
	return err
}

func (srv *Server) broadcast(message network.Datagram) error {
	bp := binpack.NewBinpacker(128)
	network.MarshallDatagram(bp, message)
	data := bp.GetData()

	srv.mutex.RLock()
	defer srv.mutex.RUnlock()

	for _, c := range srv.clients {
		_, err := srv.connection.WriteToUDP(data, c.addr)
		if err != nil {
			return err
		}
	}

	return nil
}
