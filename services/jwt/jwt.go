package jwt

import (
	"log"
	"net/http"
	"strings"
	"vehicle_golang/config"
	"vehicle_golang/utils"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
)

type JwtToken struct {
	C *config.Configuration
}

// Token ,contains data that will enrypted in JWT token
// When jwt token will decrypt, token model will returns
// Need this model to authenticate and validate resources access by loggedIn user
type Token struct {
	ID   string `json:"id"`   // User Id
	Role string `json:"role"` // User role
	jwt.StandardClaims
}

// CreateToken takes the user id as parameter
// generate JWT token and return JWT token string
func (jt *JwtToken) CreateToken(id, role string) (map[string]string, error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &Token{
		ID:   id,
		Role: role,
	})
	// token -> string, only server knows this secret(foobar)
	tokenString, err := token.SignedString([]byte(jt.C.JwtSecret))
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	m["token"] = tokenString // set response data
	return m, nil
}

// ProtectedEndPoint, authenticate all request
// takes http handler as param and performs authentication by JWT token
// if everything works fine, it redirect the request to the actual handler,
// otherwise sends unauthenticated or unauthorized error response
func (jt *JwtToken) ProtectedEndPoint(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware", r.URL)
		if strings.Contains(r.URL.Path, "/auth/") || strings.Contains(r.URL.Path, "/swagger/") {
			h.ServeHTTP(w, r)
		} else {
			// JWT from request header
			tokenString := r.Header.Get("Authorization")
			// In another way, you can decode your struct, which needs to satisfy `jwt.StandardClaims`
			t := Token{}
			token, err := jwt.ParseWithClaims(tokenString, &t, func(token *jwt.Token) (interface{}, error) {
				return []byte(jt.C.JwtSecret), nil
			})
			if !token.Valid || err != nil {
				utils.Response(w, utils.NewHTTPError(utils.Unauthenticated, http.StatusUnauthorized))
			} else {
				// set userId in context so that we can access it over the request
				// in some requests, we need login user information
				context.Set(r, "userId", t.ID) // set logged in user id in context
				context.Set(r, "role", t.Role) // set logged in user role in context
				// Redirect call to original http handler
				h.ServeHTTP(w, r)
			}
		}
	})
}
