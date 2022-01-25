package main

import (
	"context"
	pb "example.com/go-usermgmt-grpc/usermgmt"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln("Did not connect : %v", err)
	}
	defer conn.Close()
	c := pb.NewParkingLotClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	fmt.Println("Testing ParkYourCar functionality")
	res, err := c.ParkYourCar(ctx, &pb.NewCar{Name: "Aish", Id: 5})
	if err != nil {
		log.Fatalf("could not allot slot %v", err)
	}
	fmt.Printf("Slot %d\n", res.GetSlot())

	//res, err = c.ParkYourCar(ctx, &pb.NewCar{Name: "Arnav", Id: 2})
	//if err != nil {
	//	log.Fatalf("could not allot slot %v", err)
	//}
	//fmt.Printf("Slot %d\n", res.GetSlot())
	//
	//res, err = c.ParkYourCar(ctx, &pb.NewCar{Name: "Ajay", Id: 3})
	//if err != nil {
	//	log.Fatalf("could not allot slot %v", err)
	//}
	//fmt.Printf("Slot %d\n", res.GetSlot())
	//
	//res, err = c.ParkYourCar(ctx, &pb.NewCar{Name: "Adi", Id: 4})
	//if err != nil {
	//	log.Fatalf("could not allot slot %v", err)
	//}
	//fmt.Printf("Slot %d\n", res.GetSlot())
	//

	//fmt.Println("Testing Unpark your Car functionality")
	//
	//res, err := c.UnparkYourCar(ctx, &pb.Id{Id: 4})
	//if err != nil {
	//	log.Fatalf("Error occured while unparking your car %v", err)
	//}
	//log.Print("CarsList: ")
	//fmt.Printf("%v", res.GetCars())
	//fmt.Print("\n")
	//

	//fmt.Println("Testing Get All Users Functionality")
	//
	//res, err := c.GetParkedCars(ctx, &pb.GetAllCars{})
	//if err != nil {
	//	log.Fatalf("Error while fetching all the cars %v", err)
	//}
	//log.Print("\nCarsList : \n")
	//fmt.Printf("res.GetCars : %v\n", res.GetCars())

	//var t int
	//fmt.Scanf("%d", &t)
	////func (c *parkingLotClient) ParkYourCar(ctx context.Context, in *NewCar, opts ...grpc.CallOption) (*SlotNumber, error) {
	//var str string
	//
	//for i := 0; i < t; i++ {
	//	fmt.Scanf("%s", &str)
	//	if str == "ParkYourCar" {
	//		var name string
	//		var id int32
	//		fmt.Scanf("%s", &name)
	//		fmt.Scanf("%d", &id)
	//		res, err := c.ParkYourCar(ctx, &pb.NewCar{Name: name, Id: id})
	//		if err != nil {
	//			log.Fatalf("could not allot slot %v", err)
	//		}
	//		fmt.Printf("Slot %d\n", res.GetSlot())
	//	} else if str == "UnparkYourCar" {
	//		var id int32
	//		fmt.Scanf("%d", &id)
	//		res, err := c.UnparkYourCar(ctx, &pb.Id{Id: id})
	//		if err != nil {
	//			log.Fatalf("Error occured while unparking your car %v", err)
	//		}
	//		log.Print("\nCarsList : \n")
	//		fmt.Printf("res.GetCars : %v\n", res.GetCars())
	//	} else {
	//		res, err := c.GetParkedCars(ctx, &pb.GetAllCars{})
	//		if err != nil {
	//			log.Fatalf("Error while fetching all the cars %v", err)
	//		}
	//		log.Print("\nCarsList : \n")
	//		fmt.Printf("res.GetCars : %v\n", res.GetCars())
	//	}
	//}
}

//func (c *parkingLotClient) UnparkYourCar(ctx context.Context, in *Id, opts ...grpc.CallOption) (*CarList, error) {
////func (c *parkingLotClient) GetParkedCars(ctx context.Context, in *GetAllCars, opts ...grpc.CallOption) (*CarList, error) {
//1
//ParkYourCar
//Arpit
//1
