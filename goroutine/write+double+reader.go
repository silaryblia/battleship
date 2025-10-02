package goroutine

import (
	"fmt"
	"time"
)

// написать 3 функции:
//  writer - генерит числа от 1 до 10
//	doubler - умножает числа на 2, имитируя работу (500ms)
//	reader - читает и выводит на экран

func writer() <-chan int {
	ch := make(chan int)

	go func() {
		for v := range 10 {
			ch <- v + 1
		}
		close(ch)
	}()

	return ch
}

func double(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for v := range in {
			time.Sleep(500 * time.Millisecond)
			out <- v * 2
		}
		close(out)
	}()

	return out
}

func reader(out <-chan int) {
	for r := range out {
		fmt.Println("r = ", r)
	}

}

func main() {
	reader(double(writer()))
}
