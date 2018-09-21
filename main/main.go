package main

import (
	"nyubus/common"
	"nyubus/mongodb"
)

func main() {
	common.ParseCSV()
	mongodb.Find("nyubus", "route", "Z\n")
}
