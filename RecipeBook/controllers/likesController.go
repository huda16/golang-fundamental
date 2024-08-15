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

// GettingAllLikes godoc
// @Summary Get details
// @Description Get details of all likes
// @Tags Likes
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @Success 200 {object} models.Like "OK"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Router /likes [get]
func GetLike(c *gin.Context) {
	db := database.GetDB()
	likes := []models.Likes{}
	likeId := c.Param("likeId")

	if likeId != "" {
		result := db.Where("id = ?", likeId).Find(&likes)
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

		c.JSON(http.StatusOK, likes[0])
		return
	}

	err := db.Find(&likes).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, likes)
}

// AddingLike godoc
// @Summary Post details for a given Id
// @Description Post details of like corresponding to the input Id
// @Tags Likes
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @param helpers.LikeInput body helpers.LikeInput true "create like"
// @Success 201	{object} models.Like "Created"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Failure	400	{object} helpers.APIError "Bad Request"
// @Failure	500	{object} helpers.APIError "Server Error"
// @Router /likes [post]
func CreateLike(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Like := models.Likes{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Like)
	} else {
		c.ShouldBind(&Like)
	}

	Like.UserID = userID

	err := db.Debug().Create(&Like).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Like)
}

// DeletingLike godoc
// @Summary Delete like identified by the given id
// @Description Delete the order corresponding to the input id
// @Tags Likes
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @Param likeId path int true "ID of the like to be deleted"
// @Success 200 {object} helpers.DeleteResponse "OK"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Failure	400	{object} helpers.APIError "Bad Request"
// @Failure	500	{object} helpers.APIError "Server Error"
// @Router /likes/{likeId} [delete]
func DeleteLike(c *gin.Context) {
	db := database.GetDB()
	likeId, _ := strconv.Atoi(c.Param("likeId"))
	Like := models.Likes{}

	err := db.Where("id = ?", likeId).Delete(&Like).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Like with id %v has been deleted", likeId),
	})
}

// GettingLike godoc
// @Summary Get details for a given Id
// @Description Get details like corresponding to the input Id
// @Tags Likes
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @Param likeId path int true "ID of the like"
// @Success 200 {object} models.Like "OK"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Failure	400	{object} helpers.APIError "Bad Request"
// @Failure	500	{object} helpers.APIError "Server Error"
// @Router /likes/{likeId} [get]
func HelloLike(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}