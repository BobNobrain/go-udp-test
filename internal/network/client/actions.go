package client

import (
	"goudptest/internal/network"
	"goudptest/internal/util/geom"
)

func (client *Client) join() error {
	return client.send(&network.JoinCMsg{
		PlayerName: client.playerName,
	})
}

func (client *Client) Move(dv geom.Vec3) error {
	name := client.playerName

	newPos, err := client.world.MovePlayerBy(name, dv)
	if err != nil {
		return err
	}

	return client.send(&network.MoveCMsg{
		NewPosition: newPos,
	})
}

func (client *Client) Leave() error {
	err := client.send(&network.LeaveCMsg{})
	if err != nil {
		return err
	}

	return client.Disconnect()
}
