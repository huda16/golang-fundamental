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

// GettingAllComments godoc
// @Summary Get details
// @Description Get details of all comments
// @Tags Comments
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @Success 200 {object} models.Comments "OK"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Router /comments [get]
func GetComment(c *gin.Context) {
	db := database.GetDB()
	comments := []models.Comments{}
	commentId := c.Param("commentId")

	if commentId != "" {
		result := db.Where("id = ?", commentId).Find(&comments)
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

		c.JSON(http.StatusOK, comments[0])
		return
	}

	err := db.Find(&comments).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, comments)
}

// AddingComment godoc
// @Summary Post details for a given Id
// @Description Post details of comment corresponding to the input Id
// @Tags Comments
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @param helpers.CommentInput body helpers.CommentInput true "create comment"
// @Success 201	{object} models.Comments "Created"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Failure	400	{object} helpers.APIError "Bad Request"
// @Failure	500	{object} helpers.APIError "Server Error"
// @Router /comments [post]
func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comments{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID

	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Comment)
}

// UpdatingComment godoc
// @Summary Update comment identified by the given Id
// @Description Update the comment corresponding to the input Id
// @Tags Comments
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @Param commentId path int true "ID of the comment to be updated"
// @param helpers.CommentInput body helpers.CommentInput true "update comment"
// @Success 200	{object} models.Comments "OK"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Failure	400	{object} helpers.APIError "Bad Request"
// @Failure	500	{object} helpers.APIError "Server Error"
// @Router /comments/{commentId} [put]
func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comments{}
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.ID = uint(commentId)

	err := db.Model(&Comment).Where("id = ?", commentId).Updates(models.Comments{RecipeID: Comment.RecipeID, Message: Comment.Message}).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)
}

// DeletingComment godoc
// @Summary Delete comment identified by the given id
// @Description Delete the order corresponding to the input id
// @Tags Comments
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @Param commentId path int true "ID of the comment to be deleted"
// @Success 200 {object} helpers.DeleteResponse "OK"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Failure	400	{object} helpers.APIError "Bad Request"
// @Failure	500	{object} helpers.APIError "Server Error"
// @Router /comments/{commentId} [delete]
func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	Comment := models.Comments{}

	err := db.Where("id = ?", commentId).Delete(&Comment).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Comment with id %v has been deleted", commentId),
	})
}

// GettingComment godoc
// @Summary Get details for a given Id
// @Description Get details comment corresponding to the input Id
// @Tags Comments
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @Param commentId path int true "ID of the comment"
// @Success 200 {object} models.Comments "OK"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Failure	400	{object} helpers.APIError "Bad Request"
// @Failure	500	{object} helpers.APIError "Server Error"
// @Router /comments/{commentId} [get]
func HelloComment(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}
