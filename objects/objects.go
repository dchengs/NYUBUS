package objects

import (
    "encoding/csv"
    "fmt"
	"os"
    "strconv"
)

type Station struct{
	Route string
	StationName string
	TimeTable []string 
}

type Bus struct{
	Route string
	StationList []Station
}

func ParseCSV() {
    //create reader to read filename

    // Open CSV file
    f, err := os.Open("Book1.csv")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    // Read File into a Variable
    lines, err := csv.NewReader(f).ReadAll()
    if err != nil {
        panic(err)
    }

	// get the number of rows 
	rowNum := len(lines)
	// get the number of station (first row)
	stationNum := len(lines[0])
    
    //loop through the file by column
	for currentColumn:=0; currentColumn<stationNum; currentColumn++{
        stationName := ""
        stationTT := []string{}
		for currentRow:=0;currentRow < rowNum; currentRow++{
            if currentRow == 0 {
                stationName = lines[currentRow][currentColumn]
                fmt.Println("station name is: " + stationName)
                currentRow++
            }
            fmt.Println("appending time: " + lines[currentRow][currentColumn])
            stationTT = append(stationTT, lines[currentRow][currentColumn])
            fmt.Println("size of array rn: " + strconv.Itoa(len(stationTT)))
                
            }
        }
    }
    