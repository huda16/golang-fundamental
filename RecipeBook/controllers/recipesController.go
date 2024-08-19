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

// AddingRecipeComment godoc
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
func CreateRecipeComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comments{}
	userID := uint(userData["id"].(float64))
	recipeId, _ := strconv.Atoi(c.Param("recipeId"))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.RecipeID = uint(recipeId)

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
func CreateRecipeLike(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	recipeId, _ := strconv.Atoi(c.Param("recipeId"))

	Like := models.Likes{}
	userID := uint(userData["id"].(float64))

	Like.UserID = userID
	Like.RecipeID = uint(recipeId)

	result := db.Where("recipe_id = ?", recipeId).Where("user_id = ?", userID).First(&Like)
	count := result.RowsAffected

	if count > 0 {
		err := db.Where("recipe_id = ?", recipeId).Where("user_id = ?", userID).Delete(&Like).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Like with recipe_id %v and user_id %v has been deleted", recipeId, userID),
		})
		return
	}

	errCreate := db.Debug().Create(&Like).Error

	if errCreate != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": errCreate.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Like)
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
func CreateRecipeFollow(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	Recipe := models.Recipes{}
	Follow := models.Follows{}
	userID := uint(userData["id"].(float64))
	recipeId, _ := strconv.Atoi(c.Param("recipeId"))

	recipe := db.Where("id = ?", recipeId).First(&Recipe)
	count := recipe.RowsAffected

	if count < 1 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Data Not Found",
			"message": "Recipe doesn't exist",
		})
		return
	}

	Follow.FollowedID = Recipe.UserID
	Follow.FollowerID = userID

	result := db.Where("followed_id = ?", Recipe.UserID).Where("follower_id = ?", userID).First(&Follow)
	countResult := result.RowsAffected

	if countResult > 0 {
		err := db.Where("followed_id = ?", Recipe.UserID).Where("follower_id = ?", userID).Delete(&Follow).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Follow with followed_id %v and follower_id %v has been deleted", Recipe.UserID, userID),
		})
		return
	}

	errCreate := db.Debug().Create(&Follow).Error

	if errCreate != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": errCreate.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Follow)
}
