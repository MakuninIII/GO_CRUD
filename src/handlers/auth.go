package handlers

import (
    "net/http"
	"time"
    "strings"
    "github.com/golang-jwt/jwt/v5"
	"encoding/json"
)

var jwtSecret = []byte("your_secret_key") // Замените на свой секрет

func JWTMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            http.Error(w, "Missing or invalid token", http.StatusUnauthorized)
            return
        }
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        _, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return jwtSecret, nil
        })
        if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r)
    })
}

func GetToken(w http.ResponseWriter, r *http.Request) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user": "test",
        "exp": time.Now().Add(time.Hour * 1).Unix(),
    })
    tokenString, _ := token.SignedString(jwtSecret)
    json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}