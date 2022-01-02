package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	apserver "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/server"
)

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// server up
	s := apserver.New()
	s.Run()

	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	s.Shutdown(ctx)
}
