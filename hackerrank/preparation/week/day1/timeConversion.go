package main

import (
	"fmt"
)

func main() {
	fmt.Println(timeConversion("12:01:00AM"))
}


func timeConversion(s string) string {
	var myMap = map[string]string{
		"01":"13",
		"02":"14",
		"03":"15",
		"04":"16",
		"05":"17",
		"06":"18",
		"07":"19",
		"08":"20",
		"09":"21",
		"10":"22",
		"11":"23",
		"12":"00",
	}
	var result string
	if string(s[0:2]) == "12" && string(s[8]) == "P" || string(s[0:2]) != "12" && string(s[8]) == "A" {
		result = s[0:8]
	} else {
		result = myMap[string(s[0:2])] + s[2:8]
	}
	return result
}