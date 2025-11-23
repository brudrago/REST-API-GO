package controller

import (
	"net/http"
	"rest-api/model"
	"rest-api/usecase"

	"github.com/gin-gonic/gin"

	"strconv"
)

type ProductController struct {
	productUseCase *usecase.ProductUseCase
}

func NewProductController(usecase *usecase.ProductUseCase) *ProductController {
	return &ProductController{
		productUseCase: usecase,
	}
}

func (pc *ProductController) GetProducts(ctx *gin.Context) {
	products, err := pc.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (pc *ProductController) CreateProduct(ctx *gin.Context) {
	var newProduct model.Product

	if err := ctx.ShouldBindJSON(&newProduct); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	createdProduct, err := pc.productUseCase.CreateProduct(newProduct)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	ctx.JSON(http.StatusCreated, createdProduct)
}

func (pc *ProductController) GetProductByID(ctx *gin.Context) {
	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{Message: "Product ID is required"}
		ctx.JSON(http.StatusBadRequest, gin.H{response.Message: response})
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{Message: "Invalid product ID"}
		ctx.JSON(http.StatusBadRequest, gin.H{response.Message: response})
		return
	}

	product, err := pc.productUseCase.GetProductByID(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product"})
		return
	}

	if product == nil {
		response := model.Response{Message: "Product not found"}
		ctx.JSON(http.StatusNotFound, gin.H{response.Message: response})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (pc *ProductController) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{Message: "Product ID is required"}
		ctx.JSON(http.StatusBadRequest, gin.H{"message": response.Message})
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{Message: "Invalid product ID"}
		ctx.JSON(http.StatusBadRequest, gin.H{"message": response.Message})
		return
	}

	// Busca o produto atual (pra validar se existe)
	existingProduct, err := pc.productUseCase.GetProductByID(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product"})
		return
	}

	if existingProduct == nil {
		response := model.Response{Message: "Product not found"}
		ctx.JSON(http.StatusNotFound, gin.H{"message": response.Message})
		return
	}

	// Lê os novos dados do body
	var input model.Product
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	// Atualiza os campos que podem mudar
	existingProduct.Name = input.Name
	existingProduct.Price = input.Price

	updatedProduct, err := pc.productUseCase.UpdateProduct(productId, *existingProduct)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	ctx.JSON(http.StatusOK, updatedProduct)
}

func (pc *ProductController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{Message: "Product ID is required"}
		ctx.JSON(http.StatusBadRequest, gin.H{"message": response.Message})
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{Message: "Invalid product ID"}
		ctx.JSON(http.StatusBadRequest, gin.H{"message": response.Message})
		return
	}

	// Verifica se o produto existe
	existingProduct, err := pc.productUseCase.GetProductByID(productId)
	if err != nil {
		response := model.Response{Message: "Product not found"}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	if existingProduct == nil {
		response := model.Response{Message: "Product not found"}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	if err := pc.productUseCase.DeleteProduct(productId, *existingProduct); err != nil {
		response := model.Response{Message: "Failed to delete product"}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := model.Response{Message: "Product deleted successfully"}
	ctx.JSON(http.StatusOK, response)
}
