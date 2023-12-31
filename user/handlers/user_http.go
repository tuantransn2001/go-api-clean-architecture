package handlers

import (
	"food-comming-api/user/models"
	"food-comming-api/user/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type userHttpHandler struct {
	userUseCase usecases.UserUseCase
}

func NewUserHttpHandler(userUseCase usecases.UserUseCase) UserHandler {
	return &userHttpHandler{
		userUseCase: userUseCase,
	}
}

func (h *userHttpHandler) DetectUser(c echo.Context) error {
	reqBody := new(models.AddUserData)

	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return response(c, http.StatusBadRequest, "Bad request")
	}

	if err := h.userUseCase.UserDataProcessing(reqBody); err != nil {
		log.Errorf("Error processing user data: %v", err)
		return response(c, http.StatusInternalServerError, "Internal server error - Processing user data failed!")

	}
	return response(c, http.StatusOK, "Success 🚀")
}
