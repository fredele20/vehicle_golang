package main

import (
	"log"
	"net/http"
	"time"
	"vehicle_golang/config"
	"vehicle_golang/db"
	api "vehicle_golang/handlers"
	"vehicle_golang/middleware"
	"vehicle_golang/services/auth"
	"vehicle_golang/services/jwt"
	"vehicle_golang/services/user"
	"vehicle_golang/utils"

	"github.com/globalsign/mgo"
	"github.com/gorilla/mux"
)

func main() {

	// Initialize the config
	c := config.Newconfig()

	// Make connection with db and get instance
	dbSession := db.GetInstance(c)

	dbSession.SetSafe(&mgo.Safe{})

	// UserService
	userService := user.New(dbSession, c)

	// AuthService
	authService := auth.New(dbSession, c)

	// Router
	router := mux.NewRouter()

	// UserRouter
	api.UserRouter(userService, router)

	// AuthRouter
	api.AuthRouter(authService, c, router)

	// JWT services
	jwtService := jwt.JwtToken{C: c}

	// Added all middleware over all request to authenticate
	router.Use(middleware.Cors, jwtService.ProtectedEndPoint)

	// Server configuration
	srv := &http.Server{
		Handler:      utils.Headers(router),
		Addr:         c.Address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Application running at ", c.Address)

	// Server application at specified port
	log.Fatal(srv.ListenAndServe())
}
