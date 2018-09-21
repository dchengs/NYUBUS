package datatypes

//Station is a custom type that contains its corresponding route, station name, and time table
type Station struct {
	Route       string   `json:"route" bson:"route"`
	StationName string   `json:"stationame" bson:"stationname"`
	TimeTable   []string `json:"timetable" bson:"timetable"`
}

//Bus is a custom type that contains route and a list of stations
type Bus struct {
	Route       string    `json:"_id" bson:"_id"`
	StationList []Station `json:"stationlist" bson:"stationlist"`
}
