package goroutine

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

//> Работа с JWT в контексте
//> **Условие**:
//> Напиши две функции:
//> 1. `AddJWTToContext(ctx context.Context, userID int) (context.Context, error)`
//> 2. `ExtractUserIDFromContext(ctx context.Context) (int, error)`
//>
//> `AddJWTToContext` должен:
//> - Создавать JWT-токен (используй `github.com/golang-jwt/jwt/v5`).
//> - Зашифровывать в него `userID`.
//> - Возвращать новый `context.Context`, в который записан JWT-токен.
//>
//> `ExtractUserIDFromContext` должен:
//> - Извлекать JWT-токен из контекста.
//> - Расшифровывать `userID`.
//> - Вывести его на экран
//>
//> **Дополнительное условие**:
//> Создай горутину, в которой будет использоваться `ExtractUserIDFromContext`.
//Покажи, что передача контекста работает и данные можно безопасно извлекать
//между горутинами.
//>

type ctxKey string

const jwtKey ctxKey = "jwt-token"

var hmacSecret = []byte("secret-key")

// add token
func AddJWTToContext(ctx context.Context, userID int) (context.Context, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(5 * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(hmacSecret)
	if err != nil {
		return nil, err
	}
	return context.WithValue(ctx, jwtKey, signed), nil
}

// user_id from token to context
func ExtractUserIDFromContext(ctx context.Context) (int, error) {
	tokerStr := ctx.Value(jwtKey).(string)
	token, _, _ := jwt.NewParser().ParseUnverified(tokerStr, jwt.MapClaims{})
	claims := token.Claims.(jwt.MapClaims)
	return int(claims["user_id"].(float64)), nil
}

func main() {
	ctx := context.Background()

	//add token
	ctxWithJWT, err := AddJWTToContext(ctx, 123)
	if err != nil {
		panic(err)
	}

	// in main routine
	uid, _ := ExtractUserIDFromContext(ctxWithJWT)
	fmt.Println("Main goroutine extracted userID:", uid)

	// in other routine
	done := make(chan struct{})
	go func(c context.Context) {
		defer close(done)
		uid, _ := ExtractUserIDFromContext(c)
		fmt.Println("Other goroutine extracted userID:", uid)
	}(ctxWithJWT)

	<-done

}
