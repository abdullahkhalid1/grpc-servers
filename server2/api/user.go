package api

import (
	"context"
	"fmt"
	"github.com/abdullahkhalid1/grpc-servers/server1/api"
	"github.com/abdullahkhalid1/grpc-servers/server2/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(r *gin.Engine, carClient api.CarServiceClient) {
	// API users
	router := r.Group("/users")
	{
		router.POST("/create-user-with-car", func(c *gin.Context) {
			CreateUserWithCarHandler(c, carClient)
		})
		router.GET("/", func(c *gin.Context) {
			GetAllUsersHandler(c, carClient)
		})
	}
}

func CreateUserWithCarHandler(c *gin.Context, client api.CarServiceClient) {
	u := &models.User{}
	if err := c.ShouldBind(u); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("The JSON body is invalid: %v", err))
		return
	}
	err := models.CreateUserWithCar(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Something went wrong creating user: %v", err))
		return
	}
	for i, car := range u.Cars {
		cr, err := client.BuyCar(context.Background(), &api.Request{Userid: u.Id, Manufacturer: car.Manufacturer, Model: car.Model, Color: car.Color, Year: car.Year})
		if err != nil {
			c.JSON(http.StatusInternalServerError, fmt.Sprintf("Something went wrong creating user Cars: %v", err))
			return
		}
		u.Cars[i] = models.Car{Id: cr.Id, UserId: u.Id, Manufacturer: cr.Manufacturer, Model: cr.Model, Year: cr.Year, Color: cr.Color}
	}
	c.JSON(http.StatusOK, u)
}

func GetAllUsersHandler(c *gin.Context, client api.CarServiceClient) {
	users, err := models.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Something went wrong fetching users: %v", err))
		return
	}
	for i, u := range users {
		carsList, err := client.GetUserCars(context.Background(), &api.GetCarsRequest{Userid: u.Id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, fmt.Sprintf("Something went wrong fetching user Cars: %v", err))
			return
		}
		for _, car := range carsList.Cars {
			users[i].Cars = append(users[i].Cars, models.Car{Id: car.GetId(), Model: car.GetModel(), Manufacturer: car.GetManufacturer(), Year: car.GetYear(), Color: car.GetColor()})
		}
	}
	c.JSON(http.StatusOK, users)
}
