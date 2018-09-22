package mongodb

import (
	"fmt"
	dt "nyubus/datatypes"
	t "nyubus/tools"

	"gopkg.in/mgo.v2"
)

//Insert inserts bus object into database and return true and nil if insertion was successful
func Insert(database string, collection string, bus dt.Bus) (bool, error) {

	session, err := mgo.Dial("127.0.0.1")
	//check if connection is established
	if err != nil {
		return false, err
	}
	defer session.Close()
	//select db and collection
	c := session.DB(database).C(collection)
	err = c.Insert(bus)
	if mgo.IsDup(err) {
		//prompt for update
		prompt := true
		for prompt {
			response := t.Input("response")
			if response == "y" || response == "Y" {
				//do update
				prompt = false
				updated, err := Update(c, bus.Route, bus)
				if updated == false || err != nil {
					panic(err)
				}
			} else if response == "n" || response == "N" {
				fmt.Println("Insertion aborted")
				return false, nil
			} else {
				fmt.Println("Invalid response, please try again")
			}

		}
	}
	fmt.Println("Insert complete!")
	return true, nil
}
