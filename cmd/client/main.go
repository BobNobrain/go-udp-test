package main

import (
	"fmt"
	"goudptest/internal/network/client"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println("Usage: client <server_hostname> <player_name>")
		os.Exit(1)
	}

	serverHostname := args[1]
	playerName := args[2]

	c := client.NewClient(playerName)
	c.OnMsg(func(msg string) {
		fmt.Println(msg)
	})

	err := c.Connect(serverHostname)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
		return
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	<-sigChan
	fmt.Println("Disconnecting...")
	err = c.Leave()
	if err != nil {
		fmt.Println("Cannot leave properly:", err.Error())
		os.Exit(2)
	}
}
