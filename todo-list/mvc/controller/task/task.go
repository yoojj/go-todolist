package task

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"todo-list/common"
	"todo-list/mvc/model"
)

func Routes(r *gin.Engine) {

	task := r.Group("/task")
	{
		// page
		task.GET("/", taskPage)
		// task.GET("/:param", taskPage)
		// task/today
		// task/important
		// task/star

		// ajax
		task.GET("/getTask.ajax", selectTask)
		task.GET("/getTaskList.ajax", selectTaskList)
		task.POST("/setTask.ajax", insertTask)
		task.PUT("/modifyTask.ajax", updateTask)
		task.DELETE("/deleteTask.ajax", deleteTask)
	}

}

func taskPage(c *gin.Context) {

	// user 임시
	c.HTML(http.StatusOK, "task.html", common.PageData{
		PageTitle:    "task",
		ContentTitle: "모든 할 일",
		User: common.User{
			Firstname: "유",
			Lastname:  "진주",
			LoginDate: "2023-11-22",
		},
	})

}

func selectTask(c *gin.Context) {

	if common.DB_STATUS == common.SUCCESS {

		var todos []model.Todo
		result := common.DB.Find(&todos)

		println(result.Error)
		println(result.RowsAffected)

		c.SecureJSON(http.StatusOK, common.JsonData{
			Status: common.Status{Code: http.StatusOK, Desc: ""},
			Data: common.Data{
				"todo": result,
			},
		})

	} else {
		// 다른 방법
	}

}

func selectTaskList(c *gin.Context) {
}

func insertTask(c *gin.Context) {

	var params model.Todo
	if err := c.BindJSON(&params); err != nil {
		// 오류처리
	}

	params.CreateUUID()

	result := common.DB.Create(&params)
	println(result.Error)
	println(result.RowsAffected)

	c.SecureJSON(http.StatusOK, common.JsonData{
		Status: common.Status{Code: http.StatusOK, Desc: ""},
		Data: common.Data{
			"parmas": params,
		},
	})

}

func updateTask(c *gin.Context) {
}

func deleteTask(c *gin.Context) {
}
