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

// GettingAllRecipes godoc
// @Summary Get details
// @Description Get details of all recipes
// @Tags Recipes
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @Success 200 {object} models.Recipe "OK"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Router /recipes [get]
func GetRecipe(c *gin.Context) {
	db := database.GetDB()
	recipes := []models.Recipes{}
	recipeId := c.Param("recipeId")

	if recipeId != "" {
		result := db.Where("id = ?", recipeId).Find(&recipes)
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

		c.JSON(http.StatusOK, recipes[0])
		return
	}

	err := db.Find(&recipes).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, recipes)
}

// AddingRecipe godoc
// @Summary Post details for a given Id
// @Description Post details of recipe corresponding to the input Id
// @Tags Recipes
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @param helpers.RecipeInput body helpers.RecipeInput true "create recipe"
// @Success 201	{object} models.Recipe "Created"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Failure	400	{object} helpers.APIError "Bad Request"
// @Failure	500	{object} helpers.APIError "Server Error"
// @Router /recipes [post]
func CreateRecipe(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Recipe := models.Recipes{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Recipe)
	} else {
		c.ShouldBind(&Recipe)
	}

	Recipe.UserID = userID

	err := db.Debug().Create(&Recipe).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Recipe)
}

// UpdatingRecipe godoc
// @Summary Update recipe identified by the given Id
// @Description Update the recipe corresponding to the input Id
// @Tags Recipes
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @Param recipeId path int true "ID of the recipe to be updated"
// @param helpers.RecipeInput body helpers.RecipeInput true "update recipe"
// @Success 200	{object} models.Recipe "OK"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Failure	400	{object} helpers.APIError "Bad Request"
// @Failure	500	{object} helpers.APIError "Server Error"
// @Router /recipes/{recipeId} [put]
func UpdateRecipe(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Recipe := models.Recipes{}

	recipeId, _ := strconv.Atoi(c.Param("recipeId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Recipe)
	} else {
		c.ShouldBind(&Recipe)
	}

	Recipe.UserID = userID
	Recipe.ID = uint(recipeId)

	err := db.Model(&Recipe).Where("id = ?", recipeId).Updates(models.Recipes{Title: Recipe.Title, Description: Recipe.Description, Ingredients: Recipe.Ingredients, Steps: Recipe.Steps, Category: Recipe.Category, Tags: Recipe.Tags, PictureUrl: Recipe.PictureUrl}).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Recipe)
}

// DeletingRecipe godoc
// @Summary Delete recipe identified by the given id
// @Description Delete the order corresponding to the input id
// @Tags Recipes
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @Param recipeId path int true "ID of the recipe to be deleted"
// @Success 200 {object} helpers.DeleteResponse "OK"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Failure	400	{object} helpers.APIError "Bad Request"
// @Failure	500	{object} helpers.APIError "Server Error"
// @Router /recipes/{recipeId} [delete]
func DeleteRecipe(c *gin.Context) {
	db := database.GetDB()
	recipeId, _ := strconv.Atoi(c.Param("recipeId"))
	Recipe := models.Recipes{}

	err := db.Where("id = ?", recipeId).Delete(&Recipe).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Recipe with id %v has been deleted", recipeId),
	})
}

// GettingRecipe godoc
// @Summary Get details for a given Id
// @Description Get details recipe corresponding to the input Id
// @Tags Recipes
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @Param recipeId path int true "ID of the recipe"
// @Success 200 {object} models.Recipe "OK"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Failure	400	{object} helpers.APIError "Bad Request"
// @Failure	500	{object} helpers.APIError "Server Error"
// @Router /recipes/{recipeId} [get]
func HelloRecipe(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}