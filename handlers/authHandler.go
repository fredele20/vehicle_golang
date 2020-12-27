package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"vehicle_golang/config"
	"vehicle_golang/models"
	"vehicle_golang/services/auth"
	"vehicle_golang/services/jwt"
	"vehicle_golang/utils"

	"github.com/gorilla/mux"
)

// AuthHandler
type AuthHandler struct {
	au auth.AuthService
	c  *config.Configuration
}

func AuthRouter(au auth.AuthService, c *config.Configuration, router *mux.Router) {
	authHandler := &AuthHandler{au, c}
	// ------------------------------- AUTH APIs ------------------------------
	router.HandleFunc(BaseRoute+"/auth/register", authHandler.Create).Methods(http.MethodPost)
	router.HandleFunc(BaseRoute+"/auth/login", authHandler.Login).Methods(http.MethodPost)
}

func (a *AuthHandler) Create(w http.ResponseWriter, r *http.Request) {
	requestUser := new(models.User)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)
	result := make(map[string]interface{})
	if validateError := requestUser.Validate(); validateError != nil {
		fmt.Println(validateError)
		result = utils.NewHTTPCustomError(utils.BadRequest, validateError.Error(), http.StatusBadRequest)
		utils.Response(w, result)
		return
	}

	requestUser.Initialize()

	if a.au.IsUserAlreadyExists(r.Context(), requestUser.Email) {
		result = utils.NewHTTPError(utils.UserAlreadyExists, http.StatusBadRequest)
		utils.Response(w, result)
		return
	}
	err := a.au.Create(r.Context(), requestUser)
	if err != nil {
		result = utils.NewHTTPError(utils.EntityCreationError, http.StatusBadRequest)
		fmt.Println(err)
	} else {
		result["message"] = "Successfully Registered"
	}
	utils.Response(w, result)
}

func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	credentials := new(models.LoginDetails)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&credentials)

	user, err := a.au.Login(r.Context(), credentials)
	if err != nil || user == nil {
		log.Println(err)
		result := utils.NewHTTPError(utils.Unauthorized, http.StatusBadRequest)
		utils.Response(w, result)
		return
	}

	j := jwt.JwtToken{C: a.c}
	tokenMap, err := j.CreateToken(user.ID.Hex(), user.Role)
	if err != nil {
		log.Println(err)
		result := utils.NewHTTPError(utils.InternalError, 500)
		utils.Response(w, result)
		return
	}

	res := &loginRes{
		Token: tokenMap["token"],
		User:  user,
	}

	utils.Response(w, res)
}
