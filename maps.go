package main

import "fmt"

func main() {
	// Declare and initialize a map
	myMap := make(map[string]int)
	// Insert key-value pairs into the map
	myMap["apple"] = 1
	myMap["banana"] = 2
	myMap["orange"] = 3
	// Accessing values using keys
	appleValue := myMap["apple"]
	bananaValue := myMap["banana"]
	fmt.Println("Value of apple:", appleValue)
	fmt.Println("Value of banana:", bananaValue)
	// Updating values
	myMap["apple"] = 5
	fmt.Println("Updated value of apple:", myMap["apple"])
	// Deleting a key-value pair
	delete(myMap, "orange")
	fmt.Println("After deleting orange:", myMap)

	// Checking existence of a key
	value, exists := myMap["banana"]
	if exists {
		fmt.Println("Value of banana:", value)
	} else {
		fmt.Println("Banana not found in the map")
	}
	// Iterating over the map
	for key, value := range myMap {
		fmt.Println("Key:", key, "Value:", value)
	}
	// Length of the map
	fmt.Println("Length of the map:", len(myMap))
}
