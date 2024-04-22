package util

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func ContextBackground() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sCh := make(chan os.Signal, 3)
		signal.Notify(sCh, syscall.SIGINT)
		signal.Notify(sCh, syscall.SIGTERM)
		signal.Notify(sCh, syscall.SIGQUIT)
		<-sCh
		cancel()
	}()
	return ctx
}
