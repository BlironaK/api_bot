package main

import (
    "taskmanager/controllers"
    "github.com/gin-gonic/gin"
)



var tc controllers.TaskController

func main() {
    router := gin.Default()
    router.GET("/tasks", tc.ListTasks)
    router.POST("/task/update", tc.UpdateTaskController)
    router.POST("/task", tc.PostTask)
    router.DELETE("/task/delete", tc.DeleteTaskController)


    router.Run("localhost:8080")
}



