package main

import (

	
	"log"
	api "v1"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	
)


const (
	address = "localhost:8080"
)



func createUser(client api.UsersClient,user *api.UserRequest){

	//calling our rpc method
	resp, err := client.CreateUser(context.Background(), user)
	if err != nil {//call failure
		log.Fatalf("Could not create user: %v", err)
	}
	if resp.Success {//call success
		log.Printf("A new user has been added with id: %d", resp.Id)
	}
	
}

func main() {

	//connecting our client to our server
		// Set up a connection to the gRPC server.
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()


		// Creates a new UsersClient
		//this is the object of our client side
		//we pass this in our rpc method to create the user along with the data
		//the server receives this and then updates the database with the data passed..
		//..through here
		client := api.NewUsersClient(conn)


			//ON BUTTON CLICK DATA IS SENT FROM CLIENT TO SERVER
			//This below data is protobuf referred
			//we cant send this from the borwser
			//from browser we hve to send data through json
			//TODO: create rest gateway
		user := &api.UserRequest {
			Id:    1111,
			Name:  "Ban this shitty drug",
		} 

		//data passed to the rpc method 
		createUser(client,user) //TODO  : extend further

		getUser()  //TODO : extend further
}
