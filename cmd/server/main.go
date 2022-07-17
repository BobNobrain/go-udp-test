package main

import (
	"fmt"
	"goudptest/internal/domain"
	"goudptest/internal/network/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	srv := server.NewServer()
	world := domain.NewWorld(3)

	port := "8050"
	args := os.Args
	if len(args) > 1 {
		port = args[1]
	}

	srv.OnMsg(func(msg string) {
		fmt.Println(msg)
	})

	fmt.Println("Starting server on port", port)
	err := srv.Start(world, port)
	if err != nil {
		fmt.Println("Server halted:", err.Error())
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
	fmt.Println("Closing server...")
	err = srv.Stop()
	if err != nil {
		fmt.Println("Cannot stop server:", err.Error())
		os.Exit(2)
	}
}
