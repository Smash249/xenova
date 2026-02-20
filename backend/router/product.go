package router

import (
	"github.com/Smash249/xenova/backend/internal/api"
	"github.com/labstack/echo/v5"
)

func initProductRouter() {
	GroupRouterHubApp.RegisterRouterHub(
		func(public, private, admin *echo.Group) {
			public.GET("/product_series", api.ProductApi.GetProductSeries)
			admin.POST("/product_series", api.ProductApi.CreateProductSeries)
			admin.PUT("/product_series", api.ProductApi.UpdateProductSeries)
			admin.DELETE("/product_series", api.ProductApi.DeleteProductSeries)

			public.GET("/products", api.ProductApi.GetProductList)
			private.GET("/products/:product_id", api.ProductApi.GetProductDetail)
			admin.POST("/products", api.ProductApi.CreateProduct)
			admin.PUT("/products", api.ProductApi.UpdateProduct)
			admin.DELETE("/products", api.ProductApi.DeleteProduct)
		},
	)
}
