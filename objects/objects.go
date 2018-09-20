package objects

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

type Station struct {
	Route       string   `json:_id bson:_id`
	StationName string   `json:stationame bson:stationname`
	TimeTable   []string `json:timetable bson:timetable`
}

type Bus struct {
	Route       string    `json:route bson:route`
	StationList []Station `json:stationlist bson:stationlist`
}

func ParseCSV() {
	//create reader to read filename

	f, err := os.Open("Book1.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	//reading route
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Route: ")
	route, _ := reader.ReadString('\n')

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
	stationList := []Station{}
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
		newStation := Station{Route: route, StationName: stationName, TimeTable: stationTT}
		stationList = append(stationList, newStation)

	}
	//create new bus and push to database
	newBus := Bus{Route: route, StationList: stationList}
	//missing database functions for now
}
