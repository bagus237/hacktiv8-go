package main

import "fmt"


//sample defer pada readfile 

func main(){
func ReadFile(filename string) {
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		fmt.Println("ada error", err)
		return
	}

	dbyte, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("error ke 2", err)
		return
	}

	fmt.Println(string(dbyte))
	return
}
}