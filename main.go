package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/nzmprlr/highway/lane/restserver"

	"template/server/rest"
)

func main() {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	restServers := restserver.Get()
	// inject routes to rest server.
	rest.Routes(restServers[0])

	<-exit
	restserver.GracefulShutdown()
}
