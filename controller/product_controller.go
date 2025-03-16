package controller

import (
	"go-api/model"
	usecase "go-api/useCase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUseCase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messageiro": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Product data is not valid",
		})
		return
	}

	insertProduct, err := p.productUseCase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, insertProduct)
}

func (p *productController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		respose := model.Response{
			Message: "Product id is null",
		}
		ctx.JSON(http.StatusBadRequest, respose)
		return
	}
	id_product, err := strconv.Atoi(id)
	if err != nil {
		respose := model.Response{
			Message: "Product id is not valid",
		}
		ctx.JSON(http.StatusBadRequest, respose)
		return
	}

	product, err := p.productUseCase.GetProductById(id_product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if product == nil {
		respose := model.Response{
			Message: "Product not found",
		}
		ctx.JSON(http.StatusNotFound, respose)
		return
	}

	ctx.JSON(http.StatusOK, product)
}
