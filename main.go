package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/balamh/project1/netxd_customer"
	"google.golang.org/grpc"
)
func main(){
	connection,err:=grpc.Dial("localhost:5001",grpc.WithInsecure())
	if err!=nil{
		log.Fatalf("failed to connect:%v",err)
	}
	defer connection.Close()
	client :=pb.NewCustomerServiceClient(connection)
	response,err :=client.CreateCustomer(context.Background(),&pb.Customer{
		CustomerId: 1234,
		FirstName: "bala",
		LastName: "murugan",
		BankId: 123674,
		Balance: 25000,
		IsActive: true,
	})
	if err!=nil{
		log.Fatalf("failed to call:%v",err)
	}
	fmt.Printf("Response:%s\n",response)
}