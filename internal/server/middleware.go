package server


import (
    "net/http"
    "strings"
     "github.com/agoudjiliss/auth-system/data"
    "github.com/agoudjiliss/auth-system/internal/config"
    "github.com/dgrijalva/jwt-go"
    "context"
)

// JWTMiddleware is a middleware that checks for a valid JWT in the request.
func JWTMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Get the token from the Authorization header
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Authorization header missing", http.StatusUnauthorized)
            return
        }

        // Split the token from the "Bearer " prefix
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")

        // Parse the token
        claims := &datatype.Claims{}
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            // Check if the signing method is correct
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, http.ErrNotSupported
            }
            return []byte(config.Config.Jwt.Jwtkey), nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Store the username in the context (optional)
        ctx := context.WithValue(r.Context(), "username", claims.Username)
        r = r.WithContext(ctx)

        // Call the next handler
        next.ServeHTTP(w, r)
    })
}
