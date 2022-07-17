package client

import (
	"goudptest/internal/domain"
	"net"
	"sync"
)

type Client struct {
	connection *net.UDPConn
	world      *domain.World
	playerName string

	onMsg func(string)

	mutex        sync.RWMutex
	disconnected bool
}

func NewClient(playerName string) *Client {
	return &Client{
		playerName: playerName,
		world:      domain.NewWorld(255),
	}
}

func (client *Client) Connect(addr string) error {
	s, err := net.ResolveUDPAddr("udp4", addr)

	if err != nil {
		return err
	}

	connection, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		return err
	}

	client.connection = connection

	go client.listen()

	return client.join()
}

func (client *Client) Disconnect() error {
	client.mutex.Lock()
	client.disconnected = true
	client.mutex.Unlock()
	return client.connection.Close()
}

func (client *Client) OnMsg(callback func(string)) {
	client.mutex.Lock()
	defer client.mutex.Unlock()

	client.onMsg = callback
}
