package middleware

import (
	"context"
	"net/http"
	"os"

	"github.com/marsDev10/helpdesk-backend/utils"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			utils.JSONResponse(w, http.StatusUnauthorized, utils.ErrorResponse("Token requerido", ""))
			return
		}

		claims, err := utils.ValidateJWT(tokenString, os.Getenv("JWT_SECRET_KEY"))
		if err != nil {
			utils.JSONResponse(w, http.StatusUnauthorized, utils.ErrorResponse("Token inválido", err.Error()))
			return
		}

		// Guardar en contexto
		ctx := context.WithValue(r.Context(), utils.UserContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
