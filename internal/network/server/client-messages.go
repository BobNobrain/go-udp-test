package server

import (
	"goudptest/internal/network"
	"goudptest/internal/util/geom"
)

type ClientMsg interface {
	Apply(srv *Server, client *RemoteClient) error
}

type NoopMsg struct{}

func (NoopMsg) Apply(_ *Server, _ *RemoteClient) error {
	return nil
}

type JoinMsg struct {
	PlayerName string
}

func (msg JoinMsg) Apply(srv *Server, client *RemoteClient) error {
	client.PlayerName = msg.PlayerName

	err := srv.world.RegisterPlayer(msg.PlayerName)
	if err != nil {
		return err
	}

	srv.mutex.Lock()
	srv.clients[client.addr.String()] = client
	srv.onMsg(client.PlayerName + " joined from " + client.addr.String())
	srv.mutex.Unlock()

	srv.BroadcastPlayerJoin(msg.PlayerName)
	srv.SendWorldState(client)
	return nil
}

type LeaveMsg struct{}

func (msg LeaveMsg) Apply(srv *Server, client *RemoteClient) error {
	err := srv.world.UnregisterPlayer(client.PlayerName)
	if err != nil {
		return err
	}

	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	delete(srv.clients, client.addr.String())
	srv.onMsg(client.PlayerName + " left")
	return nil
}

type MoveMsg struct {
	newPosition geom.Vec3
}

func (msg MoveMsg) Apply(srv *Server, client *RemoteClient) error {
	return srv.world.MovePlayerTo(client.PlayerName, msg.newPosition)
}

func createClientMessage(d network.Datagram) ClientMsg {
	switch v := d.(type) {
	case *network.NoopCMsg:
		return &NoopMsg{}

	case *network.JoinCMsg:
		return JoinMsg{
			PlayerName: v.PlayerName,
		}

	case *network.LeaveCMsg:
		return LeaveMsg{}

	case *network.MoveCMsg:
		return MoveMsg{
			newPosition: v.NewPosition,
		}
	}

	return nil
}
