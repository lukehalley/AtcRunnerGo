package main

import "fmt"

func main() {

	dict := make(map[string]int)

	dict["pair_one"] = 1
	dict["pair_two"] = 10
	dict["pair_three"] = 5

	for key, value := range dict {
		fmt.Println("Key" , key, "Value", value)
	}

}