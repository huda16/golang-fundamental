package controllers

import (
	"fmt"
	"net/http"
	"recipebook/database"
	"recipebook/helpers"
	"recipebook/models"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// GettingAllFollow godoc
// @Summary Get details
// @Description Get details of all follow
// @Tags Follow
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @Success 200 {object} models.Follows "OK"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Router /follows [get]
func GetFollow(c *gin.Context) {
	db := database.GetDB()
	socialMedia := []models.Follows{}
	followId := c.Param("followId")

	if followId != "" {
		result := db.Where("id = ?", followId).Find(&socialMedia)
		err := result.Error
		count := result.RowsAffected
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid parameter",
			})
			return
		}

		if count < 1 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "data doesn't exist",
			})
			return
		}

		c.JSON(http.StatusOK, socialMedia[0])
		return
	}

	err := db.Find(&socialMedia).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, socialMedia)
}

// AddingFollow godoc
// @Summary Post details for a given Id
// @Description Post details of follow corresponding to the input Id
// @Tags Follow
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @param helpers.FollowInput body helpers.FollowInput true "create follow"
// @Success 201	{object} models.Follows "Created"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Failure	400	{object} helpers.APIError "Bad Request"
// @Failure	500	{object} helpers.APIError "Server Error"
// @Router /follows [post]
func CreateFollow(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Follow := models.Follows{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Follow)
	} else {
		c.ShouldBind(&Follow)
	}

	Follow.FollowerID = userID

	err := db.Debug().Create(&Follow).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Follow)
}

// DeletingFollow godoc
// @Summary Delete follow identified by the given id
// @Description Delete the order corresponding to the input id
// @Tags Follow
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @Param followId path int true "ID of the follow to be deleted"
// @Success 200 {object} helpers.DeleteResponse "OK"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Failure	400	{object} helpers.APIError "Bad Request"
// @Failure	500	{object} helpers.APIError "Server Error"
// @Router /follows/{followId} [delete]
func DeleteFollow(c *gin.Context) {
	db := database.GetDB()
	followId, _ := strconv.Atoi(c.Param("followId"))
	Follow := models.Follows{}

	err := db.Where("id = ?", followId).Delete(&Follow).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Follow with id %v has been deleted", followId),
	})
}

// GettingFollow godoc
// @Summary Get details for a given Id
// @Description Get details follow corresponding to the input Id
// @Tags Follow
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @Param followId path int true "ID of the follow"
// @Success 200 {object} models.Follows "OK"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Failure	400	{object} helpers.APIError "Bad Request"
// @Failure	500	{object} helpers.APIError "Server Error"
// @Router /follows/{followId} [get]
func HelloFollow(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}
