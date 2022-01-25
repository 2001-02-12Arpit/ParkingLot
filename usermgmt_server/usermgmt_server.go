package main

import (
	"context"
	pb "example.com/go-usermgmt-grpc/usermgmt"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = ":50051"

func NewParkingLotServer() *ParkingLotServer {
	return &ParkingLotServer{
		CarsList: &pb.CarList{},
	}
}

type ParkingLotServer struct {
	pb.UnimplementedParkingLotServer
	CarsList *pb.CarList
}

func (server *ParkingLotServer) ParkYourCar(ctx context.Context, in *pb.NewCar) (*pb.SlotNumber, error) {
	Cars := server.CarsList.Cars
	tempCar := &pb.NewCar{Id: in.GetId(), Name: in.GetName()}
	tempSlot := pb.SlotNumber{}
	if len(Cars) == 0 {
		fmt.Println("oii")
		Cars = append(Cars, tempCar)
		server.CarsList.Cars = Cars
		tempSlot.Slot = 1
		return &tempSlot, nil
	} else {
		for i := 0; i < len(Cars); i++ {
			temp := *Cars[i]
			if temp.Name == "" && temp.Id == 0 {
				temp.Name = in.GetName()
				temp.Id = in.GetId()
				Cars[i] = &temp
				tempSlot.Slot = int32(i + 1)
				server.CarsList.Cars = Cars
				return &tempSlot, nil
			}
		}
		Cars = append(Cars, tempCar)
		server.CarsList.Cars = Cars
		tempSlot.Slot = int32(len(Cars))
		return &tempSlot, nil
	}
}
func (server *ParkingLotServer) UnparkYourCar(ctx context.Context, in *pb.Id) (*pb.CarList, error) {
	Cars := server.CarsList.Cars
	for i := 0; i < len(Cars); i++ {
		temp := *Cars[i]
		if temp.GetId() == in.GetId() {
			temp.Name = ""
			temp.Id = 0
			Cars[i] = &temp
			break
		}
	}
	server.CarsList.Cars = Cars
	return server.CarsList, nil

}
func (server *ParkingLotServer) GetParkedCars(ctx context.Context, in *pb.GetAllCars) (*pb.CarList, error) {
	return server.CarsList, nil
}

func (server *ParkingLotServer) Run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterParkingLotServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	return s.Serve(lis)
}

func main() {

	var parkingLotServer *ParkingLotServer = NewParkingLotServer()
	if err := parkingLotServer.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
