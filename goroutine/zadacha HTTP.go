package goroutine

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func FetchURLs(urls []string) map[string]string {
	results := make(map[string]string)
	var wg sync.WaitGroup
	var mu sync.Mutex

	// Создаем http.Client с таймаутом в 5 секунд
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// Для каждого URL запускаем горутину
	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()

			// Создаем контекст с тайм-аутом на 5 секунд для каждого запроса
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			// Выполняем HTTP запрос с контекстом
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
			if err != nil {
				mu.Lock()
				results[u] = "error"
				mu.Unlock()
				return
			}

			resp, err := client.Do(req)
			if err != nil {
				mu.Lock()
				results[u] = "error"
				mu.Unlock()
				return
			}
			defer resp.Body.Close()

			// Ограничиваем тело ответа 100 байтами
			body, err := io.ReadAll(io.LimitReader(resp.Body, 100))
			if err != nil {
				mu.Lock()
				results[u] = "error"
				mu.Unlock()
				return
			}

			// Записываем результат в map, синхронизируя доступ через мьютекс
			mu.Lock()
			results[u] = fmt.Sprintf("Status: %s, Body: %s", resp.Status, string(body))
			mu.Unlock()

		}(url)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()
	return results
}

func main() {
	urls := []string{
		"https://golang.org/",
		"http://www.google.com/",
		"https://example.com",
		"https://httpbin.org/get",
		"https://nonexistent.url.abc", // вызовет ошибку
	}

	results := FetchURLs(urls)
	for u, r := range results {
		fmt.Printf("URL: %s\nResult: %s\n\n", u, r)
	}
}
