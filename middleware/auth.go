package middleware

import (
    "net/http"
    "strings"
    "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your_secret_key")

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Forbidden", http.StatusForbidden)
            return
        }

        tokenString := strings.Split(authHeader, " ")[1]
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })

        if err != nil {
            http.Error(w, "Forbidden", http.StatusForbidden)
            return
        }

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            r.Header.Set("userID", claims["userID"].(string))
            r.Header.Set("role", claims["role"].(string))
            next.ServeHTTP(w, r)
        } else {
            http.Error(w, "Forbidden", http.StatusForbidden)
        }
    })
}
