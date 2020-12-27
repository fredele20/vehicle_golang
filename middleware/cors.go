package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

func Cors(handler http.Handler) http.Handler {
	handleCors := cors.Default().Handler
	return handleCors(handler)
}
