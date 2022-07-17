package server

import (
	"goudptest/internal/domain"
	"net"
	"sync"
)

type Server struct {
	world      *domain.World
	port       string
	connection *net.UDPConn
	clients    map[string]*RemoteClient
	onMsg      func(string)

	mutex  sync.RWMutex
	closed bool
}

type RemoteClient struct {
	addr       *net.UDPAddr
	PlayerName string
}

func NewServer() *Server {
	return &Server{
		clients: make(map[string]*RemoteClient),
	}
}

func (srv *Server) Start(world *domain.World, port string) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	srv.world = world
	srv.port = port

	s, err := net.ResolveUDPAddr("udp4", ":"+port)
	if err != nil {
		return err
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		return err
	}

	srv.connection = connection

	go srv.startBroadcastingUpdates()
	go srv.listen()

	return nil
}
func (srv *Server) Stop() error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	srv.closed = true

	return srv.connection.Close()
}

func (srv *Server) OnMsg(callback func(string)) {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	srv.onMsg = callback
}
