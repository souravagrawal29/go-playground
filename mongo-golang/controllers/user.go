package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/souravagrawal29/go-playground/mongo-golang/models"
	"github.com/souravagrawal29/go-playground/mongo-golang/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)


type UserController struct {
	client *mongo.Client
}


func NewUserController(client *mongo.Client) *UserController {
	return &UserController{client}
}

func (uc UserController) getCollection() *mongo.Collection {
	return uc.client.Database("mongo-golang").Collection("users");
}

func (uc UserController) GetAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var users []models.User
	c, err := uc.getCollection().Find(context.TODO(), bson.D{})

	if err != nil {
		fmt.Println(err)
		utils.SetErrorResponse(w, err, http.StatusInternalServerError)
		return
	}
	if err = c.All(context.TODO(), &users); err != nil {
		fmt.Println(err)
		utils.SetErrorResponse(w, err, http.StatusInternalServerError)
		return
	}
	utils.SetSuccessResponse(w, http.StatusOK, users)
}


func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	id := params.ByName("id")
	oid, err := bson.ObjectIDFromHex(id) 

	if err != nil {
		fmt.Printf("Failed to convert id=%v to objectId, %v", id, err)
		utils.SetErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	user := models.User{}

	if err := uc.getCollection().FindOne(context.TODO(), bson.M{ "_id" : oid }).Decode(&user); err != nil {
		fmt.Println(err)
		utils.SetNotFoundErrorResponse(w, err)
		return
	}
	utils.SetSuccessResponse(w,http.StatusOK, user)
}


func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	user := models.User{}
	utils.ParseBody(r, &user)
	user.Id = bson.NewObjectID()
	user.CreatedAt = time.Now()

	if _, err := uc.getCollection().InsertOne(context.TODO(), user); err != nil {
		fmt.Printf("Failed to insert %v, %v\n", user, err)
		utils.SetErrorResponse(w, err, http.StatusInternalServerError)
		return
	}
	utils.SetSuccessResponse(w, http.StatusCreated, user)
}


func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	id := params.ByName("id")
	oid, err := bson.ObjectIDFromHex(id) 

	if err != nil {
		fmt.Printf("Failed to convert id=%v to objectId, %v", id, err)
		utils.SetErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	dr, err := uc.getCollection().DeleteOne(context.TODO(), bson.M{ "_id" : oid }); 
	
	if err != nil || dr.DeletedCount == 0 {
		fmt.Printf("Failed to delete user %v, %v\n", oid, err)
		utils.SetNotFoundErrorResponse(w, err)
		return

	}
	utils.SetSuccessResponse(w, http.StatusOK, fmt.Sprintf("Successfully deleted user=%v", oid))
}




