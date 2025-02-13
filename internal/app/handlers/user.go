package handlers

import (
	"99-user-service/internal/app/models"
	"99-user-service/internal/app/services"
	"99-user-service/internal/pkg/request"
	"99-user-service/internal/pkg/response"
	"99-user-service/internal/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	pageNumStr := c.DefaultQuery("page_num", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")

	offSet, pageSize, err := utils.ParsePaginationParams(pageNumStr, pageSizeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users, err := h.Service.GetAllUsers(offSet, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	var userResponses = make([]response.CreateUser, 0)
	for _, user := range users {
		userResponses = append(userResponses, response.CreateUser{
			ID:        user.ID,
			Name:      user.Name,
			CreatedAt: int(user.CreatedAt.UnixMicro()),
			UpdatedAt: int(user.UpdatedAt.UnixMicro()),
		})
	}

	c.JSON(http.StatusOK, gin.H{"result": true, "users": userResponses})
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req request.CreateUser

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "message": err.Error()})
		return
	}

	user := models.User{Name: req.Name}

	if err := h.Service.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user", "message": err.Error()})
		return
	}

	response := response.CreateUser{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: int(user.CreatedAt.UnixMicro()),
		UpdatedAt: int(user.UpdatedAt.UnixMicro()),
	}
	c.JSON(http.StatusCreated, gin.H{"result": true, "user": response})
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	userID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.Service.GetUserByID(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	response := response.CreateUser{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: int(user.CreatedAt.UnixMicro()),
		UpdatedAt: int(user.UpdatedAt.UnixMicro()),
	}

	c.JSON(http.StatusOK, gin.H{"result": true, "user": response})
}
