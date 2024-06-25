package ctrlc

import (
	"context"
	"log"
	"os"
	"os/signal"
)

func Context() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		signal.Stop(c)
		log.Printf("CTRL+C pressed, context canceled...")
		cancel()
	}()
	return ctx
}
