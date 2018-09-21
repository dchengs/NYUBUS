package mongodb

import (
	"bufio"
	"fmt"
	dt "nyubus/datatypes"
	"os"

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
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Would you like to update existing item? (y/n)")
			response, _ := reader.ReadString('\n')
			if response == "y" || response == "Y" {
				//do update
				prompt = false
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
