package goroutine

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

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
		return nil, fmt.Errorf("failed to sign JWT: %w", err)
	}
	return context.WithValue(ctx, jwtKey, signed), nil
}

// user_id from token to context
func ExtractUserIDFromContext(ctx context.Context) (int, error) {
	raw := ctx.Value(jwtKey)
	if raw == nil {
		return 0, fmt.Errorf("no JWT token in context")
	}

	tokenStr, ok := raw.(string)
	if !ok || tokenStr == "" {
		return 0, fmt.Errorf("invalid JWT token in context")
	}

	token, _, err := jwt.NewParser().ParseUnverified(tokenStr, jwt.MapClaims{})
	if err != nil {
		return 0, fmt.Errorf("failed to parse JWT: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("invalid claims type")
	}

	uid, ok := claims["user_id"].(float64)
	if !ok {
		return 0, fmt.Errorf("user_id missing or wrong type in claims")
	}

	return int(uid), nil
}

func main() {
	ctx := context.Background()

	//add token
	ctxWithJWT, err := AddJWTToContext(ctx, 123)
	if err != nil {
		fmt.Println("Error creating token:", err)
		return
	}

	// in main routine
	uid, err := ExtractUserIDFromContext(ctxWithJWT)
	if err != nil {
		fmt.Println("Error extracting user ID:", err)
		return
	}
	fmt.Println("Main goroutine extracted userID:", uid)

	// in other routine
	done := make(chan struct{})
	go func(c context.Context) {
		defer close(done)
		uid, err := ExtractUserIDFromContext(c)
		if err != nil {
			fmt.Println("Goroutine error extracting userID:", err)
			return
		}
		fmt.Println("Other goroutine extracted userID:", uid)
	}(ctxWithJWT)

	<-done

}
