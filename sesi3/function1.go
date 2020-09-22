package main

import "fmt"

func max(param1, param2 int) int{
	if param1 > param2 {
		return param1
	}
	return param2
}
func main(){
	no1 := 7
	no2 := 2
	fmt.Print("max(%d, %d) = %d\n", no1, no2, max(no1,no2))

	fmt.Print("max(%d, %d) = %d\n", no2, no1, max(no2,no1))
	
}

