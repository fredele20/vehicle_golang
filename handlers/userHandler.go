package handlers

import (
	"net/http"
	"vehicle_golang/services/user"
	"vehicle_golang/utils"

	"github.com/gorilla/mux"
)

// UserHandler - handles user requests
type UserHandler struct {
	u user.UserService
}

// UserRouter godoc
func UserRouter(u user.UserService, router *mux.Router) {
	userHandler := &UserHandler{u}

	// User APIs
	router.HandleFunc(BaseRoute+"/users/me", userHandler.Get).Methods(http.MethodGet)
}

// Get godoc
// @Summary Get Profile
// @Description Get user profile info
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User
// @Success 200 {object} errorRes
// @Security ApiKeyAuth
// @Router /users/me [get]
func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	user, err := h.u.Get(r.Context(), utils.GetLoggedInUserID(r))

	if err != nil {
		utils.Response(w, utils.NewHTTPError(utils.InternalError, 500))
	} else {
		utils.Response(w, user)
	}
}
