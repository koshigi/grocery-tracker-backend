package routers

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/koshigi/grocery-tracker-backend/controllers"
)

type ListRouteController struct {
	list_controller controllers.ListController
}

func NewListRouteController(listController controllers.ListController) ListRouteController {
	return ListRouteController{listController}
}

func (pc *ListRouteController) ListRoute(rg *gin.RouterGroup) {

	router := rg.Group("lists")
	router.GET("/:list_id", pc.list_controller.GetListById)
	//TODO: all these endpoints but for lists
	// router.POST("/", pc.postController.CreatePost)
	// router.GET("/", pc.postController.FindPosts)
	// router.PUT("/:postId", pc.postController.UpdatePost)
	// router.GET("/:postId", pc.postController.FindPostById)
	// router.DELETE("/:postId", pc.postController.DeletePost)
}
