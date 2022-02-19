package swagger

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/hamedsz/go-blog/docs"
)

func Routes(route *gin.Engine)  {
	docs.SwaggerInfo_swagger.Title = "Go blog documentation"
	docs.SwaggerInfo_swagger.Description = "coded by hamed sahrakhiz"
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
