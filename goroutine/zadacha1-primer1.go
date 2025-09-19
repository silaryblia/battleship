package goroutine

import (
	"fmt"
	"sync"
)

//1. **Запуск горутины и ожидание её завершения**
//
//    > Задача: Напишите функцию, которая запускает горутину,
//   выполняющую fmt.Println("Hello from goroutine!"),
//  и использует sync.WaitGroup для ожидания её завершения.
//
//    Какие способы есть ещё кроме waitGroup, чтобы дождаться выполнения горутины?
//   Приведи хотя бы 2 примера.
//    >

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Hello from goroutine!")
	}()
	wg.Wait()
}
