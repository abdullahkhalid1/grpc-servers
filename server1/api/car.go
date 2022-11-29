package api

import (
	"context"
	"github.com/abdullahkhalid1/grpc-servers/server1/models"
	"google.golang.org/grpc"
	"log"
)

type CarServer struct {
	UnimplementedCarServiceServer
}

func RegisterCarServer(s *grpc.Server) {
	RegisterCarServiceServer(s, &CarServer{})
}

/**
 * Implement Create Request
 */
func (s *CarServer) BuyCar(ctx context.Context, request *Request) (*Response, error) {
	log.Printf("Model: %v Year: %v Manufacturer: %v Color: %v", request.GetModel(), request.GetYear(), request.GetManufacturer(), request.GetColor())
	car := &models.Car{UserId: request.GetUserid(), Year: request.GetYear(), Model: request.GetModel(), Manufacturer: request.GetManufacturer(), Color: request.GetColor()}
	err := models.CreateUserCar(car)
	if err != nil {
		return &Response{}, err
	}
	return &Response{Id: car.Id, Year: request.GetYear(), Model: request.GetModel(), Manufacturer: request.GetManufacturer(), Color: request.GetColor()}, nil
}

/**
 * Implement Get Request
 */
func (s *CarServer) GetUserCars(ctx context.Context, request *GetCarsRequest) (*GetCarsResponse, error) {
	carsResponse := []*Response{}
	carsList, err := models.GetCarsByUserId(request.GetUserid())
	if err != nil {
		return &GetCarsResponse{}, err
	}
	for _, car := range carsList {
		carsResponse = append(carsResponse, &Response{Id: car.Id, Manufacturer: car.Manufacturer, Color: car.Color, Year: car.Year, Model: car.Model})
	}
	return &GetCarsResponse{Cars: carsResponse}, err
}

/**
 * Implement Delete Request
 */
func (s *CarServer) SellUserCar(ctx context.Context, request *GetCarsRequest) (*DeleteResponse, error) {
	err := models.DeleteCar(request.GetUserid())
	if err != nil {
		return &DeleteResponse{Msg: "Something went wrong selling user car: " + err.Error()}, err
	}
	return &DeleteResponse{Msg: "User Car Sold successfully"}, nil
}
