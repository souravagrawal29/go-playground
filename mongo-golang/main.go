package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/souravagrawal29/go-playground/mongo-golang/controllers"
	"github.com/souravagrawal29/go-playground/mongo-golang/models"
)


func main() {

	router := httprouter.New()
	userController := controllers.NewUserController(models.GetClient())

	router.GET("/user/:id", userController.GetUser)

	router.GET("/user", userController.GetAllUsers)
	
	router.POST("/user", userController.CreateUser)

	router.DELETE("/user/:id", userController.DeleteUser)

	err := http.ListenAndServe("localhost:8080", router)
	
	if err != nil {
		log.Fatal("Failed to start server" , err)
	}
}


