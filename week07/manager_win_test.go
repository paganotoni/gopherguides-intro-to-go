//go:build !windows
// +build !windows

package week07

import (
	"context"
	"os/signal"
	"runtime"
	"syscall"
	"testing"
	"time"
)

func TestRunInterrupt(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("this test can only run on unix systems due to the signal")
	}

	t.Parallel()

	ctx, cnFn := signal.NotifyContext(context.Background(), syscall.SIGUSR2)
	defer cnFn()

	go func() {
		time.Sleep(time.Second)

		t.Log("sending test signal")

		// send the TEST_SIGNAL to the system
		syscall.Kill(syscall.Getpid(), syscall.SIGUSR2)
	}()

	products := []*Product{
		// 1_500 milliseconds = 1.5 seconds
		&Product{Quantity: 1_500},
	}

	cprod, err := Run(ctx, 1, products...)
	if err != nil {
		t.Errorf("got error %v, want nil", err)
	}

	if len(cprod) != 0 {
		t.Errorf("got %v, want 0", len(cprod))
	}
}
