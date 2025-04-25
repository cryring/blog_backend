package system

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitForSignal() os.Signal {
	ch := make(chan os.Signal, 1)
	defer close(ch)

	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	s := <-ch
	signal.Stop(ch)
	return s
}
