package main

import (
	"fmt"
)
func cook(food string){
switch food {
case "rice":
	fmt.Println("Cooking rice...")
case "pasta":
	fmt.Println("Cooking pasta...")
default:
	fmt.Println("Unknown food item.")
}
}
func main() {
	orders:=[] string{"rice","pasta","bread"}
for _,order:=range(orders){
	if order=="pasta"{
		continue
	}
	cook(order)
}
	fmt.Println("Hello, World!")
}