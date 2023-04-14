package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/majorchork/tech-crib-africa/internal/models"
	"github.com/majorchork/tech-crib-africa/internal/services/utils"
	"github.com/majorchork/tech-crib-africa/internal/services/web"
	"net/http"
	"strconv"
)

func (h *Handler) GuestProfile(c *gin.Context) {
	guestEmail := c.Query("guest_email")
	if guestEmail == "" {
		web.JSON(c, "missing request param", http.StatusUnauthorized, nil, errors.New("invalid email"))
		return
	}
	guest, err := h.DB.GetGuestByEmail(c, guestEmail)
	if err != nil {
		web.JSON(c, "cannot get guest", http.StatusNotFound, nil, err)
		return
	}

	web.JSON(c, "successful", http.StatusOK, models.Guest{
		ID:          guest.ID,
		FullName:    guest.FullName,
		PhoneNumber: guest.PhoneNumber,
		Email:       guest.Email,
		CreatedAt:   guest.CreatedAt,
		Gender:      guest.Gender,
		Group:       guest.Group,
	}, nil)

}

func (h *Handler) GetGuests(c *gin.Context) {
	guests, err := h.DB.GetGuests(c)
	if err != nil {
		web.JSON(c, "cannot get transactions", http.StatusInternalServerError, nil, err)
		return
	}

	web.JSON(c, "successfully found", http.StatusOK, guests, nil)
}

func (h *Handler) GetGuestsByGroup(c *gin.Context) {
	group := c.Query("group")
	if group == "" {
		web.JSON(c, "missing request param", http.StatusUnauthorized, nil, errors.New("invalid group"))
		return
	}
	// covert group to int
	groupInt, err := strconv.Atoi(group)
	if err != nil {
		web.JSON(c, "invalid group", http.StatusBadRequest, nil, errors.New("invalid group"))
		return
	}

	guests, err := h.DB.GetGuestsByGroup(c, groupInt)
	if err != nil {
		web.JSON(c, "cannot get transactions", http.StatusInternalServerError, nil, err)
		return
	}

	web.JSON(c, "successfully found", http.StatusOK, guests, nil)
}

func (h *Handler) AssignGroupsAndSaveGuests(c *gin.Context) {
	_, err := h.GetUserFromContext(c)
	if err != nil {
		web.JSON(c, "invalid access token", http.StatusUnauthorized, nil, errors.New("invalid access_token"))
		return
	}

	request := models.GuestRequest{}
	err = c.ShouldBindJSON(&request)
	if err != nil {
		web.JSON(c, "bad request", http.StatusBadRequest, nil, err)
		return
	}
	groupedGuests, err := utils.AssignGroups(request.People, request.Group)
	if err != nil {
		web.JSON(c, "cannot get transactions", http.StatusInternalServerError, nil, err)
		return
	}

	// save guests
	err = h.DB.CreateGuests(c, groupedGuests)
	if err != nil {
		web.JSON(c, "cannot create guests", http.StatusInternalServerError, nil, err)
		return
	}

	web.JSON(c, "successfully found", http.StatusOK, groupedGuests, nil)
}
