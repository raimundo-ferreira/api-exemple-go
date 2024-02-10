package controller

import (
	"api-exemple/app/data/request"
	"api-exemple/app/data/response"
	"api-exemple/app/service"
	"api-exemple/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type AuthController struct {
	userService service.UserSerive
}

func NewAuthController(service service.UserSerive) *AuthController {
	return &AuthController{
		userService: service,
	}
}

// @Summary	Create token
// @Description	return token for access api
// @Accept json
// @Produce	json
// @Tags Authorization
// @Param Login	body request.Login true "Create token"
// @Success	200	{object} response.Token
// @Failure	400,500
// @Router	/api/token [post]
func (controller *AuthController) Token(ctx *gin.Context) {
	var input request.Login
	if utils.SendErroBindJson(ctx, &input) != nil {
		return
	}

	user, err := controller.userService.FindByEmailAndPassword(input)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.Msg{Message: "Email or password is incorrect"})
		return
	}

	token, err := utils.GenerateNewAccessToken(user)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.Msg{Message: "Failed generate token"})
		return
	}

	ctx.JSON(http.StatusOK, response.Token{Token: token, User: user})
}

// @Summary	Register user
// @Description	Create user for access api
// @Accept json
// @Produce	json
// @Tags Authorization
// @Param AccessKey query string true "Access key for register user" Format(GUID)
// @Param Register	body request.Register true "Register user"
// @Success	200
// @Failure	400,500
// @Router	/api/register [post]
func (controller *AuthController) Register(ctx *gin.Context) {
	accessKey := ctx.Request.URL.Query().Get("AccessKey")
	if accessKey != "42975313-085b-4c09-90c0-11fe6ba6de5d" {
		ctx.JSON(http.StatusBadRequest, response.Msg{Message: "Access key inválid!"})
		return
	}

	var input request.Register
	if utils.SendErroBindJson(ctx, &input) != nil {
		return
	}

	user, err := controller.userService.FindEmailRegistered(input.Username)
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

	var arg = request.CreateUser{
		Name:     input.Name,
		Email:    input.Username,
		Password: input.Password,
		Active:   true,
	}

	user, err = controller.userService.Create(arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Msg{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response.Msg{Message: "Successful registered user"})
}

// @Summary	Get state
// @Description	return state of authorization
// @Accept json
// @Produce	json
// @Tags Authorization
// @Success	200	{object} response.State
// @Failure	401,500
// @Security ApiKeyAuth
// @Router /api/state [get]
func (controller *AuthController) State(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response.State{State: "Authorized"})
}
