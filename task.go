package heartbeats

import (
	"context"
	"fmt"
	"time"
)

func ProcessingTask(
	ctx context.Context, letras chan rune, interval time.Duration,
) (<-chan struct{}, <-chan string) {

	heartbeats := make(chan struct{}, 1)
	names := make(chan string)

	go func() {
		defer close(heartbeats)
		defer close(names)

		beat := time.NewTicker(interval)
		defer beat.Stop()

		for letra := range letras {
			select {
			case <-ctx.Done():
				return
			case <-beat.C:
				select {
				case heartbeats <- struct{}{}:
				default:
				}
			case names <- dicionario[letra]:
				lether := dicionario[letra]
				fmt.Printf("Letra: %s \n", lether)

				time.Sleep(3 * time.Second) // Simula um tempo de espera para vermos o hearbeats
			}
		}
	}()

	return heartbeats, names
}
