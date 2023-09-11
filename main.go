package main

import (
	"fmt"

	u "github.com/vicchass/dog_shelterV2/util"
)

// slice of all the Breed struct
var breed_slice []u.Breed

// declare all the the breed
var pitbull = u.Breed{Name: "pitbul", Croquette_day: 20, Size: "big", Has_hair: false}
var labrador = u.Breed{Name: "labrador", Croquette_day: 15, Size: "big", Has_hair: true}
var york = u.Breed{Name: "yorkshire", Croquette_day: 5, Size: "small", Has_hair: true}

// slice of all the Breed struct
var dog_slice []u.Dog

// declare the dogs that are already in the shelter
var mike = u.Dog{The_name: "mike", The_age: 5, The_breed: pitbull, Is_here: true}
var bobby = u.Dog{The_name: "bobby", The_age: 10, The_breed: york, Is_here: true}

func main() {
	//append all breed to breed_slice
	breed_slice = append(breed_slice, pitbull, labrador, york)

	dog_slice = append(dog_slice, mike, bobby)

	// START LOGIC

}

// END MAIN

// add a new dog to the dogs in shelter
func add_dog(all_dogs_shelter []u.Dog, the_breed []u.Breed) {
	// create a name , check dont exist

	// get all existing name in a slice
	dogs_names := u.Return_names_dogs(dog_slice)
	// new name
	var new_name string
	for {
		fmt.Println("what is the name of dog?")
		fmt.Scan()

	}

}
