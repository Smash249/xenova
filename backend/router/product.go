package router

import (
	"github.com/Smash249/xenova/backend/internal/api"
	"github.com/labstack/echo/v5"
)

func initProductRouter() {
	GroupRouterHubApp.RegisterRouterHub(
		func(public, private *echo.Group) {
			public.GET("/product_series", api.ProductApi.GetProductSeries)
			private.POST("/product_series", api.ProductApi.CreateProductSeries)
			private.PUT("/product_series", api.ProductApi.UpdateProductSeries)
			private.DELETE("/product_series", api.ProductApi.DeleteProductSeries)

			public.GET("/products", api.ProductApi.GetProductList)
			public.GET("/products/:product_id", api.ProductApi.GetProductDetail)
			private.POST("/products", api.ProductApi.CreateProduct)
			private.PUT("/products", api.ProductApi.UpdateProduct)
			private.DELETE("/products", api.ProductApi.DeleteProduct)
		},
	)
}
