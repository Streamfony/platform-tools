package errsignal

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"
)

var ErrStopped = errors.New("stopped")

func NewListener(ctx context.Context) func() error {
	return func() error {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(stop)

		select {
		case <-ctx.Done():
			return nil
		case <-stop:
			return ErrStopped
		}
	}
}

func IsStopped(err error) bool {
	return errors.Is(err, ErrStopped)
}
