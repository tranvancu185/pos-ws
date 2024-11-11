package controller

import (
	"fmt"
	"path/filepath"
	"strconv"
	"tranvancu185/vey-pos-ws/global"
	"tranvancu185/vey-pos-ws/internal/model/rq"
	"tranvancu185/vey-pos-ws/internal/model/rs"
	"tranvancu185/vey-pos-ws/internal/service"
	"tranvancu185/vey-pos-ws/internal/uconst/messagecode"
	"tranvancu185/vey-pos-ws/pkg/response"
	"tranvancu185/vey-pos-ws/pkg/utils/ufile"
	urand "tranvancu185/vey-pos-ws/pkg/utils/urandom"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProductController struct {
	productService service.IProductService
}

func NewProductController(
	productService service.IProductService,
) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (mc *ProductController) GetListProduct(c *gin.Context) {
	var queryParams rq.GetListProductRequest
	if err := c.ShouldBindQuery(&queryParams); err != nil {
		c.Error(err)
		return
	}

	result, errRS := mc.productService.GetListProduct(&queryParams)
	if errRS != nil {
		c.Error(errRS)
		return
	}

	data := rs.GetListResponse{
		Page:     queryParams.Page,
		PageSize: queryParams.PageSize,
		Rows:     result,
	}

	if queryParams.Total != 0 {
		total, er := mc.productService.GetTotalProduct(&queryParams)
		if er != nil {
			c.Error(er)
			return
		}
		data.Total = total
	}

	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		Data:        data,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}

func (mc *ProductController) CreateProduct(c *gin.Context) {
	var params rq.CreateProductRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		c.Error(err)
		return
	}

	// Validate username, password
	validate := validator.New()
	if err := validate.Struct(params); err != nil {
		c.Error(err)
		return
	}

	id, errRS := mc.productService.CreateProduct(&params)
	if errRS != nil {
		c.Error(errRS)
		return
	}

	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		Data:        id,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}

func (mc *ProductController) UpdateProduct(c *gin.Context) {
	var params rq.UpdateProductRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		c.Error(err)
		return
	}

	// Validate username, password
	validate := validator.New()
	if err := validate.Struct(params); err != nil {
		c.Error(err)
		return
	}

	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	errRS := mc.productService.UpdateProduct(idInt, &params)
	if errRS != nil {
		c.Error(errRS)
		return
	}

	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		Data:        nil,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}

func (mc *ProductController) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	result, errRS := mc.productService.GetProductByID(idInt)
	if errRS != nil {
		c.Error(errRS)
		return
	}

	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		Data:        result,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}

func (mc *ProductController) DeleteProductByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	errRS := mc.productService.DeleteProductByID(idInt)
	if errRS != nil {
		c.Error(errRS)
		return
	}

	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}

func (mc *ProductController) SearchProduct(c *gin.Context) {
	var queryParams rq.SearchProductRequest
	if err := c.ShouldBindQuery(&queryParams); err != nil {
		c.Error(err)
		return
	}

	result, errRS := mc.productService.SearchProduct(&queryParams)
	if errRS != nil {
		c.Error(errRS)
		return
	}

	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		Data:        result,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}

func (mc *ProductController) UpdateProductImageByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	// Get product
	product, err := mc.productService.GetProductByID(idInt)
	if err != nil {
		c.Error(err)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.Error(err)
		return
	}

	// Tạo đường dẫn lưu trữ file
	filename := fmt.Sprintf("%d_%s_%s", idInt, urand.RandomTenDigit(), "product_image.png")
	path := filepath.Join(global.Config.Path.PathAvatar, filename)

	// Sử dụng một goroutine duy nhất để xử lý tất cả các tác vụ
	go func() {
		// Remove old avatar
		if product.ProductImage != "" {
			avatarPath := filepath.Join(global.Config.Path.PathAvatar, product.ProductImage)
			errRemove := ufile.RemoveFile(avatarPath)
			if errRemove != nil {
				global.SendLog("Error remove file", "error", errRemove)
			}
		}
	}()

	// Lưu file vào thư mục
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.Error(err) // Ghi nhận lỗi vào context
		return
	}

	errRS := mc.productService.UpdateProductImageByID(idInt, filename)
	if errRS != nil {
		c.Error(errRS)
		return
	}

	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}
