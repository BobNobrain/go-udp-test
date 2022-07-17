package client

import (
	"goudptest/internal/domain"
	"goudptest/internal/network"
)

type ServerMsg interface {
	Apply(client *Client) error
}

type JoinMsg struct {
	playerName string
}

func (msg JoinMsg) Apply(client *Client) error {
	client.onMsg(msg.playerName + " joined")
	return client.world.RegisterPlayer(msg.playerName)
}

type LeaveMsg struct {
	playerName string
}

func (msg LeaveMsg) Apply(client *Client) error {
	client.onMsg(msg.playerName + " left")
	return client.world.UnregisterPlayer(msg.playerName)
}

type KickMsg struct {
	reason string
}

func (msg KickMsg) Apply(client *Client) error {
	client.mutex.RLock()
	client.onMsg("You have been kicked")
	client.onMsg(msg.reason)
	client.mutex.RUnlock()

	client.Disconnect()

	return nil
}

type UpdateMsg struct {
	updatedPlayers []*domain.Player
}

func (msg UpdateMsg) Apply(client *Client) error {
	client.onMsg("updates arrived")
	return client.world.ApplyPlayerUpdates(msg.updatedPlayers)
}

func createServerMessage(d network.Datagram) ServerMsg {
	switch v := d.(type) {
	case *network.JoinSMsg:
		return &JoinMsg{
			playerName: v.PlayerName,
		}

	case *network.LeaveSMsg:
		return &LeaveMsg{
			playerName: v.PlayerName,
		}

	case *network.KickSMsg:
		return KickMsg{
			reason: v.Reason,
		}

	case *network.UpdateSMsg:
		return UpdateMsg{
			updatedPlayers: v.UpdatedPlayers,
		}
	}

	return nil
}
