package util

import "fmt"

//declare the dog struct
type Dog struct {
	The_name  string
	The_age   uint
	The_breed Breed
	Is_here   bool
}

type Breed struct {
	Name          string
	Croquette_day int
	Size          string
	Has_hair      bool
}

// return a slice of strings containing all the dogs name
func Return_names_dogs(slice_dogs []Dog) []string {
	var names_slices []string
	for _, v := range slice_dogs {
		names_slices = append(names_slices, v.The_name)
	}
	return names_slices
}

// return a slice of strings of all breed name
func Return_breed_name(slice_breed []Breed) []string {
	var breed_slice_name []string
	for _, v := range slice_breed {
		breed_slice_name = append(breed_slice_name, v.Name)
	}
	return breed_slice_name
}

//return a breed type based on its name
func Return_breed(slice_breed []Breed, the_name string) Breed {
	for _, v := range slice_breed {
		if v.Name == the_name {
			return v
		}
	}
	empty_breed := Breed{}
	fmt.Println("COULDNT FIND BREED")
	return empty_breed

}
