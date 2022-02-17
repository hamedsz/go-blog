package home

import "github.com/gin-gonic/gin"

func (controller *HomeController) Show(c *gin.Context)  {
	c.String(200 , "hell world")
}
