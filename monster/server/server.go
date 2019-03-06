package main

import (
	"context"
	"log"
	"time"
	api "v1"
	//"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"

	"net"
	"google.golang.org/grpc"

	//"github.com/golang/protobuf/ptypes"
	//"github.com/golang/protobuf/ptypes/wrappers"
)


const (
	port = ":8080"

)

type server struct {
	//savedUsers []*api.UserRequest
}
	

func (s *server) CreateUser(ctx context.Context, in *api.UserRequest) (*api.UserResponse, error) {

		//MONGO DB CONNECTION
		log.Printf("connecting to MongoDB...")

		// Create MongoDB client
		client, err := mongo.NewClient("mongodb://127.0.0.1:27017")
		if err != nil {
			log.Fatalf("failed to create new MongoDB client: %#v", err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
	
		// Connect client
		if err = client.Connect(ctx); err != nil {
			log.Fatalf("failed to connect to MongoDB: %#v", err)
		}
	
		log.Printf("connected successfully")

			// Get collection from database
	//-------------DATABASE OBJECT----------
	coll := client.Database("whodb").Collection("whoforms")


	
	//s.savedCustomers = append(s.savedCustomers, in)
	log.Printf("insert data into DATABASE....")


	// Insert data into the collection
	res, err := coll.InsertOne(ctx, in)
	if err != nil {
		log.Fatalf("failed to insert", err)
	}
	id := res.InsertedID
	log.Printf("inserted new item with id=%v successfully", id)	

	return &api.UserResponse{Id: in.Id, Success: true}, nil
}

func(s *server) GetPatient (ctx context.context, in *api.PatientRequest) (*api.PatientResponse , error) {

		//MONGO DB CONNECTION
		log.Printf("connecting to MongoDB...")

		// Create MongoDB client
		client, err := mongo.NewClient("mongodb://127.0.0.1:27017")
		if err != nil {
			log.Fatalf("failed to create new MongoDB client: %#v", err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
	
		// Connect client
		if err = client.Connect(ctx); err != nil {
			log.Fatalf("failed to connect to MongoDB: %#v", err)
		}
	
		log.Printf("connected successfully")

			// Get collection from database
	//-------------DATABASE OBJECT----------
	coll := client.Database("patientsdb").Collection("patients")


	
	//s.savedCustomers = append(s.savedCustomers, in)
	log.Printf("getting data from DATABASE....")


}


func main() {


	//EXPOSING TO PORT 8080
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()
	api.RegisterUsersServer(s, &server{})
	s.Serve(lis)





//	CreateUser(ctx,in)


	
// CreateCustomer creates a new Customer



}
















































//protoc api/api.proto --go_out=plugins=grpc:server


//protoc api.proto --go_out=plugins=grpc:C:\Users\Amogh\go\src

/*
message Data {

    int64 int64Value = 2;
    string stringValue = 4; 

}

*/




/* ORIGINAL MAIN FUNCTION BACKUP



	log.Printf("connecting to MongoDB...")

	// Create MongoDB client
	client, err := mongo.NewClient("mongodb://127.0.0.1:27017")
	if err != nil {
		log.Fatalf("failed to create new MongoDB client: %#v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect client
	if err = client.Connect(ctx); err != nil {
		log.Fatalf("failed to connect to MongoDB: %#v", err)
	}

	log.Printf("connected successfully")

	// Get collection from database
	//-------------DATABASE OBJECT----------
	coll := client.Database("experiments").Collection("proto")



	// Fill in data structure
	//THE DATA IS THE MESSAGE DEFINED IN THE .PROTO FILE 
	//WE POPULATE OUR MESSAGE TO BE SENT OVER THE SERVER TO THE DATABASE
	in := &api.Data{
		
		Int64Value:  100,
		
		StringValue: "ISHAN",

	}

	log.Printf("insert data into DATABASE....")

	// Insert data into the collection
	res, err := coll.InsertOne(ctx, &in)
	if err != nil {
		log.Fatalf("insert data into collection", err)
	}
	id := res.InsertedID
	log.Printf("inserted new item with id=%v successfully", id)			//we print the unique ID of the object just inserted into the database

	// Create filter and output structure to read data from collection
	var out api.Data
	filter := bson.D{{Key: "_id", Value: id}}

	// Read data from collection
	err = coll.FindOne(ctx, filter).Decode(&out)
	if err != nil {
		log.Fatalf("failed to read data (id=%v) from collection <experiments.proto>: %#v", id, err)
	}

	log.Print("read successfully")



	*/