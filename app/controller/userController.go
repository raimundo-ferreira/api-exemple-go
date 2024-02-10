package controller

import (
	"api-exemple/app/data/request"
	"api-exemple/app/data/response"
	"api-exemple/app/model"
	"api-exemple/app/service"
	"api-exemple/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type UserController struct {
	userService service.UserSerive
}

func NewUserController(service service.UserSerive) *UserController {
	return &UserController{
		userService: service,
	}
}

// @Summary	Get all users
// @Description	return list of user
// @Accept json
// @Produce	json
// @Tags User
// @Success	200	{object} []model.User
// @Failure	401,500
// @Security ApiKeyAuth
// @Router /api/user [get]
func (controller *UserController) GetAll(ctx *gin.Context) {
	users, err := controller.userService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Msg{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// @Summary	Get user by id
// @Description	return user by id
// @Accept json
// @Produce	json
// @Tags User
// @Param id path string true "enter the user id"
// @Success	200	{object} model.User
// @Failure	401,404,500
// @Security ApiKeyAuth
// @Router /api/user/{id} [get]
func (controller *UserController) GetById(ctx *gin.Context) {
	user, err := controller.userService.FindById(ctx.Param("id"))
	if utils.SendNotFoundOrError(ctx, err) != nil {
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// @Summary	Create user
// @Description	return user created
// @Accept json
// @Produce	json
// @Tags User
// @Param User	body request.CreateUser true "Create user"
// @Success	200	{object} model.User
// @Failure	400,401,500
// @Security ApiKeyAuth
// @Router	/api/user [post]
func (controller *UserController) Create(ctx *gin.Context) {
	var arg request.CreateUser
	if utils.SendErroBindJson(ctx, &arg) != nil {
		return
	}

	user, err := controller.userService.FindEmailRegistered(arg.Email)
	if err != nil {
		if err != pgx.ErrNoRows {
			ctx.JSON(http.StatusInternalServerError, response.Msg{Message: err.Error()})
			return
		}
	}

	if user.Name != "" {
		ctx.JSON(http.StatusBadRequest, response.Msg{Message: "Email registered for another user!"})
		return
	}

	user, err = controller.userService.Create(arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Msg{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// @Summary	Update user
// @Description	return user updated
// @Accept json
// @Produce	json
// @Tags User
// @Param User	body model.User true "Update user"
// @Success	200	{object} model.User
// @Failure	400,401,404,500
// @Security ApiKeyAuth
// @Router	/api/user [put]
func (controller *UserController) Update(ctx *gin.Context) {
	var user model.User
	if utils.SendErroBindJson(ctx, &user) != nil {
		return
	}

	err := controller.userService.Update(user)
	if utils.SendNotFoundOrError(ctx, err) != nil {
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// @Summary	Delete user
// @Description	return message deleted
// @Accept json
// @Produce	json
// @Tags User
// @Param id path string true "enter the user id"
// @Success	200
// @Failure	401,404,500
// @Security ApiKeyAuth
// @Router /api/user/{id} [delete]
func (controller *UserController) Delete(ctx *gin.Context) {
	err := controller.userService.Delete(ctx.Param("id"))
	if utils.SendNotFoundOrError(ctx, err) != nil {
		return
	}

	ctx.JSON(http.StatusOK, response.Msg{Message: "deleted"})
}
