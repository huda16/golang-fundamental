package helpers

type RegisterInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Bio      string `json:"bio"`
}

type RegisterResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

type LoginInput struct {
	Email    string `json:"email" valid:"required"`
	Password string `json:"password" valid:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type CommentInput struct {
	Message  string `json:"message" valid:"required"`
	RecipeID uint   `json:"recipe_id" valid:"required"`
}

type FollowInput struct {
	FollowedID uint `json:"followed_id" valid:"required"`
}

type LikeInput struct {
	RecipeID uint `json:"recipe_id" valid:"required"`
}

type RecipeCommentInput struct {
	Message string `json:"message" valid:"required"`
}

type RecipeInput struct {
	Title       string `json:"title" valid:"required"`
	Description string `json:"description" valid:"required"`
	Ingredients string `json:"ingredients" valid:"required"`
	Steps       string `json:"steos" valid:"required"`
	PictureUrl  string `json:"picture_url" valid:"required"`
	Category    string `json:"category"`
	Tags        string `json:"tags"`
}

type DeleteResponse struct {
	Message string `json:"message"`
}

type APIError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
