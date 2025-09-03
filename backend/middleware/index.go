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

// RoleMiddleware recibe los roles permitidos para la ruta
func RoleMiddleware(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			claims, err := utils.GetUserFromContext(r)

			if err != nil {
				utils.JSONResponse(w, http.StatusUnauthorized, utils.ErrorResponse("No autorizado", err.Error()))
				return
			}

			// Verificar si el rol del usuario está permitido
			userRole := claims.Role
			for _, role := range allowedRoles {
				if string(userRole) == role {
					next.ServeHTTP(w, r)
					return
				}
			}

			// Si no tiene permiso
			http.Error(w, "No tienes permisos para acceder a este recurso", http.StatusForbidden)
		})
	}
}
