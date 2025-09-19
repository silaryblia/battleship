package goroutine

import (
	"fmt"
	"sync"
)

func main() {
	mutex := &sync.Mutex{}
	mutex.Lock()

	go func() {
		defer mutex.Unlock()
		fmt.Println("Hello from goroutine!")

	}()

	// Ждем, пока горутина выполнится и разблокирует мьютекс
	mutex.Lock() // Блокируемся здесь, ждем разблокировки
	mutex.Unlock()
}
