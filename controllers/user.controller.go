package controllers

import (
	"net/http"

	"github.com/necmettindev/currency-conversion/models/user"
	"github.com/necmettindev/currency-conversion/services/authservice"
	"github.com/necmettindev/currency-conversion/services/userservice"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	PostRegister(*gin.Context)
	PostLogin(*gin.Context)
}

type userController struct {
	us userservice.UserService
	as authservice.AuthService
}

func NewUserController(
	us userservice.UserService,
	as authservice.AuthService) UserController {
	return &userController{
		us: us,
		as: as,
	}
}

// @BasePath /v1
// PostRegister godoc
// @Summary Register
// @Description Register
// @Tags User
// @Accept  json
// @Produce  json
// @Param data body user.UserRegisterInput true "Register Input"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /user/register [post]
func (ctl *userController) PostRegister(c *gin.Context) {
	var userInput user.UserRegisterInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		HTTPRes(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	u := InputToUser(userInput)

	if err := ctl.us.Create(&u); err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	err := ctl.login(c, &u)
	if err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
}

// @BasePath /v1
// PostLogin godoc
// @Summary Login
// @Description Login
// @Tags User
// @Accept  json
// @Produce  json
// @Param data body user.UserLoginInput true "Login Input"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /users/login [post]
func (ctl *userController) PostLogin(c *gin.Context) {
	var userInput user.UserLoginInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		HTTPRes(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	user, err := ctl.us.GetByUsername(userInput.Username)
	if err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	err = ctl.us.ComparePassword(userInput.Password, user.Password)
	if err != nil {
		HTTPRes(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err = ctl.login(c, user)
	if err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
}

func (ctl *userController) login(c *gin.Context, u *user.User) error {
	token, err := ctl.as.IssueToken(*u)
	if err != nil {
		return err
	}
	userOutput := MapToUserOutput(u)
	out := gin.H{"token": token, "user": userOutput}
	HTTPRes(c, http.StatusOK, "ok", out)
	return nil
}
