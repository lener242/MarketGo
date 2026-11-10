package handlers

import (
	"MarketGo/database"
	"MarketGo/models"
	"MarketGo/utils"
	"database/sql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	var body struct {
		Name     string `json:"name" binding:"required"`
		Phone    string `json:"phone" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		ErrorJSON(c, http.StatusBadRequest, "invalid body")
		return
	}

	body.Name = strings.TrimSpace(body.Name)
	body.Phone = strings.TrimSpace(body.Phone)

	if len(body.Name) < 2 {
		ErrorJSON(c, http.StatusBadRequest, "name must be at least 2 characters")
		return
	}
	if len(body.Phone) < 7 {
		ErrorJSON(c, http.StatusBadRequest, "invalid phone number")
		return
	}
	if len(body.Password) < 6 {
		ErrorJSON(c, http.StatusBadRequest, "password must be at least 6 characters")
		return
	}

	hash, err := utils.HashPassword(body.Password)
	if err != nil {
		ErrorJSON(c, http.StatusInternalServerError, "hash error")
		return
	}

	err = models.CreateUser(database.DB, body.Name, body.Phone, hash)
	if err != nil {
		ErrorJSON(c, http.StatusConflict, "user already exists or db error")
		return
	}

	SuccessJSON(c, gin.H{"message": "registered"})
}

func LoginHandler(c *gin.Context) {
	var body struct {
		Phone    string `json:"phone" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		ErrorJSON(c, http.StatusBadRequest, "invalid body")
		return
	}

	user, err := models.GetUserByPhone(database.DB, strings.TrimSpace(body.Phone))
	if err != nil {
		if err == sql.ErrNoRows {
			ErrorJSON(c, http.StatusUnauthorized, "invalid phone or password")
			return
		}
		ErrorJSON(c, http.StatusInternalServerError, "db error")
		return
	}

	if !utils.CheckPasswordHash(body.Password, user.Password) {
		ErrorJSON(c, http.StatusUnauthorized, "invalid phone or password")
		return
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		ErrorJSON(c, http.StatusInternalServerError, "token error")
		return
	}

	username := ""
	if user.Username.Valid {
		username = user.Username.String
	}

	SuccessJSON(c, gin.H{
		"message":  "logged in",
		"token":    token,
		"name":     user.Name,
		"phone":    user.Phone,
		"username": username,
	})
}
