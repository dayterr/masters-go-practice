package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dayterr/masters-go-practice/internal/client"
)

func main() {
	metricsClient := client.NewClient()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cancel()

	metricsClient.Run(ctx)
	time.Sleep(10 * time.Second)

}
