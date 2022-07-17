package client

import (
	"goudptest/internal/network"
	"goudptest/internal/util/binpack"
)

func (client *Client) send(message network.Datagram) error {
	bp := binpack.NewBinpacker(128)
	network.MarshallDatagram(bp, message)

	_, err := client.connection.Write(bp.GetData())
	return err
}
