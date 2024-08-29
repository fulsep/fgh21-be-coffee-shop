package routers

import (
	"RGT/konis/controllers"

	"github.com/gin-gonic/gin"
)

func ProfileRouters(rg *gin.RouterGroup) {
	rg.GET("", controllers.ListAllProfiles)
	rg.PATCH("/:id", controllers.UpdateProfile)
	rg.GET("/:id", controllers.FindProfileById)
}
