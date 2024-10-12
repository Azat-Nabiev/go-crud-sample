package handlers

import (
	"awesomeProject/internal/models"
	"awesomeProject/internal/services"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

type UserHandler struct {
	userService *services.UserService
	logger      *zap.SugaredLogger
}

func NewUserHandler(userService *services.UserService, logger *zap.SugaredLogger) *UserHandler {
	return &UserHandler{
		userService: userService,
		logger:      logger,
	}
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Get a list of all users
// @ID get-all-users
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /user/all [get]
func (h *UserHandler) GetAllUsers(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	users, _ := h.userService.GetAll(ctx)

	c.JSON(http.StatusOK, users)
}

// GetAllBooks godoc
// @Summary Get all users by their books
// @Description Get a list of all users and their books
// @ID get-all
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /user/all/books [get]
func (h *UserHandler) GetAll(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 1000000*time.Second)
	defer cancel()

	users, _ := h.userService.GetAllUserBooks(ctx)

	c.JSON(http.StatusOK, users)
}

// GetUser godoc
// @Summary Show a user
// @Description get string by ID
// @ID get-user
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Header 200 {string} Token "qwerty"
// @Failure 400 {string} string "We need ID!!"
// @Failure 404 {string} string "Can not find ID"
// @Router /user/{id} [get]
func (h *UserHandler) GetUserById(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err != nil {
		h.logger.Errorw("Error during parsing path param",
			"error", err.Error(),
			"userId", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	userName, _ := h.userService.GetById(ctx, userID)
	c.JSON(http.StatusOK, userName)
}

// AddUser godoc
// @Summary Add a user
// @Description Add a user to the database
// @ID add-user
// @Accept  json
// @Produce  json
// @Param user body models.User true "User"
// @Success 201 {object} models.User
// @Header 201 {string} Location "Location of the created user"
// @Failure 400 {string} string "Invalid input"
// @Router /user [post]
func (h *UserHandler) AddUser(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10000000*time.Second)
	defer cancel()

	h.userService.Add(ctx, &newUser)

	c.JSON(http.StatusCreated, newUser)
}

// GetUser godoc
// @Summary Delete a user
// @Description delete user by ID
// @ID delete-user
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {string} string "We need ID!!"
// @Failure 404 {string} string "Can not find ID"
// @Router /user/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10000000*time.Second)
	defer cancel()

	user, _ := h.userService.DeleteByID(ctx, userID)
	c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update user details by ID
// @ID update-user
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param user body models.User true "User data"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string "Invalid user ID or data"
// @Router /user/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
		return
	}

	updatedUser, err := h.userService.Update(ctx, userID, &newUser)

	c.JSON(http.StatusCreated, updatedUser)
}
