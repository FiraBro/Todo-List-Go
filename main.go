package main

import (
	"errors"
	"fmt"
)
func prossesTracking( ) error {
	fmt.Println("Process Tracking")
	names := []string{"Alice", "Bob", "Charlie"}
	for index, value := range names {
		fmt.Printf("the valeu of index %d is %s\n", index, value)
	}
	return errors.New("Process Tracking Error")
}
func main() {
	prossesTracking()
	
}