package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/majorchork/tech-crib-africa/internal/models"
	"github.com/majorchork/tech-crib-africa/internal/services/jwt"
	"github.com/majorchork/tech-crib-africa/internal/services/utils"
	"github.com/majorchork/tech-crib-africa/internal/services/web"
	"net/http"
)

func (h *Handler) SignUp(c *gin.Context) {

	request := models.UserRequest{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		web.JSON(c, "bad request", http.StatusBadRequest, nil, err)
		return
	}

	err = h.DB.CreateUser(c, request)
	if err != nil {
		web.JSON(c, "cannot create user", http.StatusInternalServerError, nil, err)
		return
	}

	web.JSON(c, "user created", http.StatusCreated, nil, nil)
}

func (h *Handler) Login(c *gin.Context) {

	request := models.LoginRequest{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		web.JSON(c, "bad request", http.StatusBadRequest, nil, err)
		return
	}

	user, err := h.DB.GetUserByEmail(c, request.Email)
	if err != nil {
		web.JSON(c, "cannot retrieve user", http.StatusInternalServerError, nil, err)
		return
	}

	if !(utils.ComputeHash(request.Password, user.Salt) == user.PasswordHash) {
		web.JSON(c, "invalid password", http.StatusUnauthorized, nil, err)
		return
	}

	token, err := jwt.GenerateToken(user.Email, jwt.AccessTokenValidity)
	if err != nil {
		web.JSON(c, "internal server error", http.StatusInternalServerError, nil, err)
		return
	}

	web.JSON(c, "login successful", http.StatusOK, models.LoginResponse{
		Token: token,
		User:  user,
	}, nil)
}

func (h *Handler) AdminProfile(c *gin.Context) {
	user, err := h.GetUserFromContext(c)
	if err != nil {
		web.JSON(c, "invalid access token", http.StatusUnauthorized, nil, errors.New("invalid access_token"))
		return
	}

	web.JSON(c, "successful", http.StatusOK, models.Admin{
		ID:          user.ID,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		CreatedAt:   user.CreatedAt,
	}, nil)

}
