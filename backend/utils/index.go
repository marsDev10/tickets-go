package utils

import (
	"errors"
	"net/http"
)

// Función auxiliar para convertir punteros
func IntPtr(i int) *int {
	return &i
}

func StringPtr(s string) *string {
	return &s
}

type contextKey string

const UserContextKey = contextKey("user")

func GetUserFromContext(r *http.Request) (*Claims, error) {
	claims, ok := r.Context().Value(UserContextKey).(*Claims)
	if !ok || claims == nil {
		return nil, errors.New("no se encontraron claims en el contexto")
	}
	return claims, nil
}
