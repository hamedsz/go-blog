package swagger

import (
	"github.com/gin-gonic/gin"
	"github.com/hamedsz/go-blog/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes(route *gin.Engine)  {
	docs.SwaggerInfo_swagger.Title = "Go blog documentation."
	docs.SwaggerInfo_swagger.Description = "coded by hamed sahrakhiz"
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
