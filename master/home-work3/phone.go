package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	addressBook := make(map[string][]int)

	addressBook["Alex"] = []int{89996543210}
	addressBook["Bob"] = []int{89167243812}
	addressBook["Bob"] = append(addressBook["Bob"], 89155243627)
	addressBook["John"] = []int{7777777}

	fmt.Println(addressBook)

	for name, numbers := range addressBook {
		fmt.Println("Абонент:", name)
		for i, number := range numbers {
			fmt.Printf("\t %v) %v \n", i+1, number)
		}
	}
	b, err := json.Marshal(addressBook)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(b)

	fmt.Println(string(b))

	err = ioutil.WriteFile("phone.json", b, 0644)
	if err != nil {

		log.Fatal(err)
	}

}
