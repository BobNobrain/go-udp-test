package client

import (
	"goudptest/internal/network"
	"goudptest/internal/util/binpack"
)

func (client *Client) listen() {
	buffer := make([]byte, 128)

	for {
		n, _, err := client.connection.ReadFromUDP(buffer)

		client.mutex.RLock()
		if client.disconnected {
			return
		}
		client.mutex.RUnlock()

		if err != nil {
			client.mutex.RLock()
			client.onMsg(err.Error())
			client.mutex.RUnlock()
			continue
		}

		bp := binpack.NewBinpackerAround(buffer[:n])

		datagram, dgErr := network.UnmarshallDatagram(bp)

		if dgErr != nil {
			client.mutex.RLock()
			client.onMsg(dgErr.Error())
			client.mutex.RUnlock()
			continue
		}

		serverMessage := createServerMessage(*datagram)
		applyErr := serverMessage.Apply(client)

		if applyErr != nil {
			client.mutex.RLock()
			client.onMsg(applyErr.Error())
			client.mutex.RUnlock()
			continue
		}
	}
}
