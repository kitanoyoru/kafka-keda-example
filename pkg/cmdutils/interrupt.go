package cmdutils

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func InterruptChan() <-chan interface{} {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	ret := make(chan interface{}, 1)
	go func() {
		s := <-c
		ret <- s
		close(ret)
	}()

	return ret
}

func NewInterruptContext() (context.Context, context.CancelFunc) {
	interruptChan := InterruptChan()

	return InterruptContextFromChan(interruptChan)
}

func InterruptContextFromChan(interruptChan <-chan interface{}) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-interruptChan
		cancel()
	}()

	return ctx, cancel
}
