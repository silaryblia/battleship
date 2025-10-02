package goroutine

import "fmt"

func main() {
	ch := make(chan int)

	for i := range 5 {
		go func() {
			ch <- i + 1
		}()
	}

	for range 5 {
		val := <-ch
		fmt.Println(val)
	}

}
