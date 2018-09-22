package common

import (
	"encoding/csv"
	dt "nyubus/datatypes"
	mdb "nyubus/mongodb"
	t "nyubus/tools"
	"os"
)

//ParseCSV parses a given schedule of a route to objects and store them in database
func ParseCSV() {
	//create reader to read filename
	fileName := t.Input("filename")
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	//reading route
	route := t.Input("route")

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	// get the number of rows
	rowNum := len(lines)
	// get the number of station (first row)
	stationNum := len(lines[0])
	// list of station
	stationList := []dt.Station{}
	//loop through the file by column
	for currentColumn := 0; currentColumn < stationNum; currentColumn++ {
		stationName := ""
		stationTT := []string{}
		for currentRow := 0; currentRow < rowNum; currentRow++ {
			if currentRow == 0 {
				stationName = lines[currentRow][currentColumn]
				currentRow++
			}
			stationTT = append(stationTT, lines[currentRow][currentColumn])
		}
		newStation := dt.Station{Route: route, StationName: stationName, TimeTable: stationTT}
		stationList = append(stationList, newStation)

	}
	//create new bus and push to database
	newBus := dt.Bus{Route: route, StationList: stationList}
	mdb.Insert("nyubus", "route", newBus)

}
