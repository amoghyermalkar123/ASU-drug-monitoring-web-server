package repository


import (
	"log"

	. "models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type SingularServer struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "patient"
)

// Establish a connection to database
func (s *SingularServer) Connect() { // this pointer to our structure describes that we use the same original copy of our server for each and every function call
    //
	session, err := mgo.Dial(s.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(s.Database)
}

// Insert a patient into database
func (s *SingularServer) Insert(patiente Asu) error {
	err := db.C(COLLECTION).Insert(&patiente)
	return err
}

// Find list of patients
func (s *SingularServer) FindAll() ([]Asu, error) {
	var patientes []Asu
	err := db.C(COLLECTION).Find(bson.M{}).All(&patientes)
	return patientes, err //array of patients returned
}

// Find a patient by its id
func (s *SingularServer) FindById(id string) (Asu , error) {
	var patientes Asu
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&patientes)
	return patientes, err //one patient returned
}