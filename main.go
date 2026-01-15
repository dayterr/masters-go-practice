package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dayterr/masters-go-practice/internal/client"
)

func main() {
	metricsClient := client.NewClient()

	fmt.Println("client created")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cancel()

	metricsClient.Run(ctx)

}
