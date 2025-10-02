package goroutine

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func fanin(ctx context.Context, chans []chan int) chan int {
	out := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}
		for _, ch := range chans {
			wg.Add(1)
			go func(ch chan int) {
				defer wg.Done()
				for {
					select {
					case v, ok := <-ch:
						if !ok {
							return
						}
						select {
						case out <- v:
						case <-ctx.Done():
							return
						}
					case <-ctx.Done():
						return
					}

				}
			}(ch)
		}

		wg.Wait()
		close(out)
	}()

	return out
}

func fanout(in chan int, numChans int, f func(int) int) []chan int {
	chans := make([]chan int, numChans)

	for i := range numChans {
		chans[i] = pipeline(in, f)
	}

	return chans
}

func pipeline(in chan int, f func(int) int) chan int {
	out := make(chan int)

	go func() {
		for v := range in {
			out <- f(v)
		}
		close(out)
	}()

	return out
}

var numWorkers = 10

func generate() chan int {
	in := make(chan int)

	go func() {
		for i := range 100 {
			in <- i
		}
		close(in)
	}()

	return in
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	now := time.Now()
	for v := range fanin(ctx, fanout(generate(), numWorkers, square)) {
		fmt.Println(v)
	}

	timeFanin := time.Since(now)

	fmt.Println("time fanin:", timeFanin)

}

func square(a int) int {
	timeConsuming1()
	return a * a
}

// cpu intensive
func timeConsuming1() {
	counter := 0

	for range 100000 {
		counter++
	}
}
