package goroutine

import (
	"fmt"
	"math/rand"
	"sync"
)

//4. **Потокобезопасный инкремент - Mutex.**
//
//    > Задача: Напишите программу, где 10 горутин инкрементируют один счётчик,
//   защищая его sync.Mutex.
//
//    1. Что если не обложить мютексом? Воспроизвести race condition.
//    >

func main() {
	wg := &sync.WaitGroup{}
	mutex := sync.Mutex{}

	number := rand.Intn(10)
	result := 0

	wg.Add(10)

	for range 10 {
		go func() {
			defer wg.Done()
			mutex.Lock()
			result += number
			mutex.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(result)
}
