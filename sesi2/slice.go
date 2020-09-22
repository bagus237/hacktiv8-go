package main


func main(){
/*
	fruits := [4]string{"apel","Pisang","Dukuh","Durian"}
	fmt.Println(fruits)
	for i :=0; 1 < len(fruits) i++{
		fmt.Println(fruits[i])
	}
	for key val := range fruits {
		fmt.Println(fmt.Sprintf(%d:%s, key, val))
	}
	var animals [3]string
	animals[0]
	animals [2]
	fmt.Println{animals[1]}
*/
var arrayInt [2] [3]int
arrayInt[0][0] = 4
arrayInt[0][1] = 5 
arrayInt[0][2] = 6
arrayInt[1][0] = 2
arrayInt[1][1] = 1
arrayInt[1][2] = 3

for key, val := range arrayInt {
	fmt.Println(fmt.Sprintf("%d:%v", key, val))
	for k, v := range val {
		fmt.Println(fmt.Sprintf("%d:%d", k, v))
	}
}
newArray := [2] [2] [3] int{{{0,5,1}}, {{0,2,3}}, {{0,5,1}},{{0,2,3}}

}