package mongodb

import (
	dt "nyubus/datatypes"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Update takes the target route, database name, collection name, and the new bus object and update the existing bus object in database
func Update(c *mgo.Collection, route string, newBus dt.Bus) (bool, error) {
	err := c.UpdateId(bson.M{"_id": route}, newBus)
	if err != nil {
		return false, err
	}
	return true, nil
}
