package main

import (
	"fmt"
)

func main() {
	dataMap := map[string]int{"jan": 10, "feb": 25, "mar": 11}
	status := checkMap(dataMap, func(data map[string]int) string {
		if data["feb"] == 25 {
			return "feb"
		}
		return "error"
	})
	fmt.Println(status)
}

func checkMap(data map[string]int, cond func(map[string]int) string) string {
	if cond(data) != "error" {
		return cond(data)
	}

	return "error"
}