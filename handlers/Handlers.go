package handlers

import (
	"GoProject/db"
	"GoProject/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetTasks(ctx *gin.Context) {
	fmt.Println("Get request!")

	tasks := db.DB.GetTasks()
	ctx.IndentedJSON(http.StatusOK, tasks)
}

func CreateTask(ctx *gin.Context) {
	fmt.Println("Post request!")

	var newTask model.Task
	if err := ctx.BindJSON(&newTask); err != nil {
		return
	}

	db.DB.CreateTAsk(newTask, ctx)
}

func DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, nil)
	}
	db.DB.DeleteTask(idInt, ctx)
}
