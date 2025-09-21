package goroutine

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	input := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	var processedCount atomic.Int64

	go StartBatchProcessor(ctx, input, &processedCount)

	go func() {
		for i := 0; i <= 20; i++ {
			input <- i
			time.Sleep(300 * time.Millisecond)
		}
		close(input)
		cancel()
	}()

	<-ctx.Done()
	fmt.Println("Main: processing stopped")
}

func StartBatchProcessor(ctx context.Context, input <-chan int, counter *atomic.Int64) {
	const maxBatchSize = 5
	const batchTimeout = 2 * time.Second

	batch := make([]int, 0, maxBatchSize)
	timer := time.NewTimer(batchTimeout)
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			if len(batch) > 0 {
				fmt.Println("Processed batch:", batch)
				counter.Add(int64(len(batch)))
			}
			fmt.Println("Context canceled. Exiting.")
			return

		case val, ok := <-input:
			if !ok {
				if len(batch) > 0 {
					fmt.Println("Processed batch:", batch)
					counter.Add(int64(len(batch)))
				}
				fmt.Println("Input channel closed. Exiting.")
				return
			}

			batch = append(batch, val)
			if len(batch) == maxBatchSize {
				fmt.Println("Processed batch:", batch)
				counter.Add(int64(len(batch)))

				batch = batch[:0]

				if !timer.Stop() {
					select {
					case <-timer.C:
					default:
					}
				}
				timer.Reset(batchTimeout)
			}

		case <-timer.C:
			if len(batch) > 0 {
				fmt.Println("Processed batch (timeout):", batch)
				counter.Add(int64(len(batch)))

				batch = batch[:0]
			}
			timer.Reset(batchTimeout)
		}
	}
}
