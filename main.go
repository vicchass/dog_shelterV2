package main

import (
	"fmt"
	"os"
	"slices"

	u "github.com/vicchass/dog_shelterV2/util"
)

// slice of all the Breed struct
var breed_slice []u.Breed

// declare all the the breed
var pitbull = u.Breed{Name: "pitbul", Croquette_day: 20, Size: "big", Has_hair: false}
var labrador = u.Breed{Name: "labrador", Croquette_day: 15, Size: "big", Has_hair: true}
var york = u.Breed{Name: "yorkshire", Croquette_day: 5, Size: "small", Has_hair: true}

// slice of all the dog struct
var dog_slice []u.Dog

// declare the dogs that are already in the shelter
var mike = u.Dog{The_name: "mike", The_age: 5, The_breed: pitbull, Is_here: true}
var bobby = u.Dog{The_name: "bobby", The_age: 10, The_breed: york, Is_here: true}

func main() {
	//append all breed to breed_slice
	breed_slice = append(breed_slice, pitbull, labrador, york)

	dog_slice = append(dog_slice, mike, bobby)

	// START LOGIC

	for {
		fmt.Println("what do you want to do \n a:add a dog \n b:print all dogs \n q:quit")
		var choice string
		fmt.Scan(&choice)
		switch choice {
		//add a new dog
		case "a":
			dog_slice = append(dog_slice, add_dog(dog_slice, breed_slice))
		//print all dogs
		case "b":
			fmt.Println("---------------")
			print_all_dogs(dog_slice)
			fmt.Println("---------------")
		case "q":
			os.Exit(3)
		}

	}

}

// END MAIN

// add a new dog to the dogs in shelter
func add_dog(all_dogs_shelter []u.Dog, the_breed []u.Breed) u.Dog {
	// create a name , check if it dont already  exist

	// get all existing name in a slice
	dogs_names := u.Return_names_dogs(dog_slice)
	// new name
	var new_name string
	for {
		fmt.Println("what is the name of dog?")
		fmt.Scan(&new_name)
		if slices.Contains(dogs_names, new_name) {
			fmt.Println(" name already exist ,choose new one ")
		} else {
			break
		}
	}

	//type age
	var new_age uint

	// get the input for the age
	for {
		fmt.Println("what is the age of your dog ?")
		if _, err := fmt.Scan(&new_age); err != nil {
			fmt.Println("error with your input , try again")
			continue
		} else {
			if new_age > 20 {
				fmt.Println("Your dog seems very old , type age correctly? y/n")
				var ans string
				fmt.Scan(&ans)
				if ans == "y" {
					break
				} else {
					continue
				}
			}
			break
		}
	}

	//choose breed

	breed_names := u.Return_breed_name(the_breed)
	var breed_input string
	var out_breed u.Breed

	for {
		//print all existing breed
		fmt.Println("what is the breed of your dog?")
		fmt.Println("the breeds are : ")
		for _, v := range breed_names {
			fmt.Println(v)
		}

		if _, err := fmt.Scan(&breed_input); err != nil {
			fmt.Println(" error with input try again")
			continue
		}

		out_breed = u.Return_breed(the_breed, breed_input)
		break

	}

	out_dog := u.Dog{The_name: new_name, The_age: new_age, The_breed: out_breed, Is_here: true}
	return out_dog
}

// print all the dog in the shelter
func print_all_dogs(all_dogs []u.Dog) {
	for _, v := range all_dogs {
		fmt.Println("we have", v.The_name, " who is a ", v.The_breed.Name)
	}

}
