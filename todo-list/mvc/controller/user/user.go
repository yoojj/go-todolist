package user

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {

	user := r.Group("/user")
	{
		user.GET("/:param", selectUser)
	}

}

func selectUser(c *gin.Context) {

}
