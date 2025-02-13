package request

type CreateUser struct {
	Name string `form:"name" binding:"required"`
}
