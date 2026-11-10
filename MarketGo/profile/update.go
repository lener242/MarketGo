package profile

import (
	"MarketGo/database"
	"MarketGo/handlers"
	"MarketGo/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type updateReq struct {
	Username string `json:"username,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

func UpdateProfile(c *gin.Context) {
	uid := c.GetInt("userID")

	var r updateReq
	if err := c.ShouldBindJSON(&r); err != nil {
		handlers.ErrorJSON(c, http.StatusBadRequest, "invalid body")
		return
	}

	r.Phone = strings.TrimSpace(r.Phone)
	r.Username = strings.TrimSpace(r.Username)

	if r.Phone == "" && r.Username == "" {
		handlers.ErrorJSON(c, http.StatusBadRequest, "nothing to update")
		return
	}

	if r.Phone != "" && len(r.Phone) < 7 {
		handlers.ErrorJSON(c, http.StatusBadRequest, "invalid phone number")
		return
	}

	if r.Phone != "" {
		existing, err := models.GetUserByPhone(database.DB, r.Phone)
		if err == nil && existing.ID != uid {
			handlers.ErrorJSON(c, http.StatusBadRequest, "phone already used")
			return
		}
	}

	if err := models.UpdateUser(database.DB, uid, r.Username, r.Phone); err != nil {
		handlers.ErrorJSON(c, http.StatusInternalServerError, "db error")
		return
	}
	handlers.SuccessJSON(c, gin.H{"message": "updated"})
}
