package main

import (
	"github.com/dayterr/masters-go-practice/internal/client"
)

func main() {
	metricsClient := client.NewClient()

	metricsClient.Run()

}
