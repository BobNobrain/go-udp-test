package server

import "goudptest/internal/network"

func (srv *Server) BroadcastPlayerJoin(name string) {
	srv.broadcast(&network.JoinSMsg{
		PlayerName: name,
	})
}

func (srv *Server) BroadcastPlayerLeave(name string) {
	srv.broadcast(&network.LeaveSMsg{
		PlayerName: name,
	})
}

func (srv *Server) KickPlayer(player *RemoteClient, reason string) {
	srv.send(player, &network.KickSMsg{
		Reason: reason,
	})

	srv.mutex.Lock()
	delete(srv.clients, player.addr.String())
	srv.mutex.Unlock()

	srv.BroadcastPlayerLeave(player.PlayerName)
}

func (srv *Server) BroadcastUpdates() error {
	srv.mutex.RLock()
	defer srv.mutex.RUnlock()

	updates := srv.world.GatherPlayerUpdates()

	if len(updates) == 0 {
		return nil
	}

	return srv.broadcast(&network.UpdateSMsg{
		UpdatedPlayers: updates,
	})
}

func (srv *Server) SendWorldState(client *RemoteClient) error {
	srv.mutex.RLock()
	defer srv.mutex.RUnlock()

	return srv.send(client, &network.UpdateSMsg{
		UpdatedPlayers: srv.world.GatherPlayerUpdates(),
	})
}
