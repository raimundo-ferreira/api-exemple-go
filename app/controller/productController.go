package controller

import (
	"api-exemple/app/data/request"
	"api-exemple/app/data/response"
	"api-exemple/app/model"
	"api-exemple/app/service"
	"api-exemple/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService service.ProductSerive
}

func NewProductController(service service.ProductSerive) *ProductController {
	return &ProductController{
		productService: service,
	}
}

// @Summary	Get all products
// @Description	return list of product
// @Accept json
// @Produce	json
// @Tags Product
// @Success	200	{object} []model.Product
// @Failure	401,500
// @Security ApiKeyAuth
// @Router /api/product [get]
func (controller *ProductController) GetAll(ctx *gin.Context) {
	products, err := controller.productService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Msg{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

// @Summary	Get product by id
// @Description	return product by id
// @Accept json
// @Produce	json
// @Tags Product
// @Param id path string true "enter the product id"
// @Success	200	{object} model.Product
// @Failure	401,404,500
// @Security ApiKeyAuth
// @Router /api/product/{id} [get]
func (controller *ProductController) GetById(ctx *gin.Context) {
	product, err := controller.productService.FindById(ctx.Param("id"))
	if utils.SendNotFoundOrError(ctx, err) != nil {
		return
	}
	ctx.JSON(http.StatusOK, product)
}

// @Summary	Create product
// @Description	return product created
// @Accept json
// @Produce	json
// @Tags Product
// @Param Product	body request.CreateProduct true "Create product"
// @Success	200	{object} model.Product
// @Failure	400,401,500
// @Security ApiKeyAuth
// @Router	/api/product [post]
func (controller *ProductController) Create(ctx *gin.Context) {
	var arg request.CreateProduct
	if utils.SendErroBindJson(ctx, &arg) != nil {
		return
	}

	product, err := controller.productService.Create(arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Msg{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

// @Summary	Update product
// @Description	return product updated
// @Accept json
// @Produce	json
// @Tags Product
// @Param Product	body model.Product true "Update product"
// @Success	200	{object} model.Product
// @Failure	400,401,404,500
// @Security ApiKeyAuth
// @Router	/api/product [put]
func (controller *ProductController) Update(ctx *gin.Context) {
	var product model.Product
	if utils.SendErroBindJson(ctx, &product) != nil {
		return
	}

	err := controller.productService.Update(product)
	if utils.SendNotFoundOrError(ctx, err) != nil {
		return
	}

	ctx.JSON(http.StatusOK, product)
}

// @Summary	Delete product
// @Description	return message deleted
// @Accept json
// @Produce	json
// @Tags Product
// @Param id path string true "enter the product id"
// @Success	200
// @Failure	401,404,500
// @Security ApiKeyAuth
// @Router /api/product/{id} [delete]
func (controller *ProductController) Delete(ctx *gin.Context) {
	err := controller.productService.Delete(ctx.Param("id"))
	if utils.SendNotFoundOrError(ctx, err) != nil {
		return
	}

	ctx.JSON(http.StatusOK, response.Msg{Message: "deleted"})
}
