package util

//declare the dog struct
type Dog struct {
	The_name  string
	The_age   int
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
