package mongodb

import (
	"fmt"
	"log"
	dt "nyubus/datatypes"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Insert inserts interface object into database
func Insert(database string, collection string, item interface{}) (bool, error) {

	session, err := mgo.Dial("127.0.0.1")
	//check if connection is established
	if err != nil {
		return false, err
	}
	defer session.Close()
	c := session.DB(database).C(collection)
	err = c.Insert(item)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	result := dt.Bus{}
	err = c.Find(bson.M{"route": "A"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("stationList", result.StationList)
	return true, nil
}
