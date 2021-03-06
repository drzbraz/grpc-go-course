package main

import (
	"context"
	"fmt"
	"log"

	"github.com/drzbraz/grpc-go-course/greet/greetpb"

	"google.golang.org/grpc"
)

func main(){

	fmt.Println("Hello I'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil{
		log.Fatal("could not connect: %v",err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	doUnary(c)


	
}

func doUnary(c greetpb.GreetServiceClient){ 
	fmt.Println("Starting to do a Unary RPC....")

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{FirstName: "John", LastName: "Marek"},
	}

	res,err := c.Greet(context.Background(),req)
	if err != nil{
		log.Fatal("could not %v",err)
	}

	log.Println("Response from Greet: %v", res.Result)
}