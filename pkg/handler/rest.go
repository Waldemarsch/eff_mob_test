package handler

import (
	"eff_mob_test/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetUser(c *gin.Context) {
	params := c.Request.URL.Query()

	if params.Has("id") {
		id, err := strconv.Atoi(params.Get("id"))

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		user, err := h.services.GetSingleUser(id)

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, user)
		return
	}

	if _, exists := params["limit"]; !exists {
		newErrorResponse(c, http.StatusBadRequest, "limit parameter is missing")
		return
	}

	if _, exists := params["page"]; !exists {
		newErrorResponse(c, http.StatusBadRequest, "page parameter is missing")
		return
	}

	users, err := h.services.GetUser(params)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *Handler) CreateUser(c *gin.Context) {
	var user main_models.User

	decoder := json.NewDecoder(c.Request.Body)

	err := decoder.Decode(&user)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userJSON, err := json.Marshal(user)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.CreateUser(userJSON)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{Error: nil, Message: "User is added"})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	var UserToUpdate struct {
		ID int `json:"id"`
	}

	decoder := json.NewDecoder(c.Request.Body)

	err := decoder.Decode(&UserToUpdate)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.DeleteUser(UserToUpdate.ID)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{Error: nil, Message: "User is deleted"})
}

func (h *Handler) UpdateUser(c *gin.Context) {

	var UserToUpdate struct {
		ID          int    `json:"ID,omitempty"`
		Name        string `json:"name,omitempty"`
		Surname     string `json:"surname,omitempty"`
		Patronymic  string `json:"patronymic,omitempty"`
		Age         int    `json:"age,omitempty"`
		Gender      string `json:"gender,omitempty"`
		Nationality string `json:"nationality,omitempty"`
	}

	decoder := json.NewDecoder(c.Request.Body)

	err := decoder.Decode(&UserToUpdate)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id := UserToUpdate.ID

	UserToUpdate.ID = 0

	UserToUpdateJSON, err := json.Marshal(UserToUpdate)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.UpdateUser(id, UserToUpdateJSON)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{Error: nil, Message: "User is updated"})
}
