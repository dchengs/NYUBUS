package mongodb

import (
	"fmt"
	dt "nyubus/datatypes"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Find takes database, collection, and route to search for associated bus object and return it if found
func Find(database, collection, route string) (dt.Bus, error) {

	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB(database).C(collection)
	//result is of type bus
	result := dt.Bus{}
	err = c.Find(bson.M{"_id": route}).One(&result)
	if err != nil {
		return result, err
	}
	fmt.Println(result)
	return result, nil
}
