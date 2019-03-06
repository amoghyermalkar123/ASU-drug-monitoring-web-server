package models


import "gopkg.in/mgo.v2/bson"


type Asu struct {
	ID          bson.ObjectId     `bson:"_id" json:"id"`
	id			int				  `bson:"uid" json:"uid"`
	DrugName    string            `bson:"drug_name" json:"drug_name"`
	ADRdesc     string      	  `bson:"ADRdesc" json:"ADRdesc"`
	PatientName string            `bson:"pat_name" json:"pat_name"`
	type EventDetails struct	
	{
		DateofOnset int			  `bson:"dateoo" json:"dateoo"` 	 
		Outcome int			   `bson:"outcome" json:"outcome"`
		RecoveryDate int		    `bson:"recdate" json:"recdate"`
	} `bson:"ev" json:"ev"`	
	
	
}




/* BACKUP ORIGINAL STRUCTURE

type Asu struct {
	ID          bson.ObjectId     `bson:"_id" json:"id"`
	//id			int				  `bson:"uid" json:"uid"`
	DrugName    string            `bson:"drug_name" json:"drug_name"`
	ADRdesc     string      	  `bson:"ADRdesc" json:"ADRdesc"`
	PatientName string            `bson:"pat_name" json:"pat_name"`
}


*/