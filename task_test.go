package heartbeats

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestProcessingTask(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	letras := make(chan rune)
	go func() {
		defer close(letras)
		for i := 'a'; i <= 'g'; i++ {
			letras <- i
		}
	}()

	heartbeats, words := ProcessingTask(ctx, letras, time.Second)

	for {
		select {
		case <-ctx.Done():
			return
		case <-heartbeats:
			fmt.Printf("Application Up! \n")

		case letra, err := <-words:
			if !err {
				return
			}
			if _, notfound := dicionario[rune(letra[0])]; !notfound {
				t.Errorf("Letra %s não encontrada", letra)
			}
		}
	}
}
