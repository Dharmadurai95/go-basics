package main

import (
	"fmt"
)

type addressMapping struct {
	street  string
	city    string
	state   string
	pincode string
}

func main() {
	makeEducationMap()
	/*
		In Go, a map is a built-in data structure that represents a collection of key-value pairs. It allows you to associate a value (the "value" part) with a unique key, similar to how objects work in JavaScript
		Key can be any type but all key should be same type . same condition for value
	*/

	// declare map in three types
	/*
		1. var age map[string]string
		2. age := make(map[string]string)
		3. age :=map[string]int {
			"dharma":25,
			"lohi":23,
			"thor":27
		}
	*/
	// address mapping
	subiAddr := addressMapping{
		street:  "west street",
		city:    "Kulithalai",
		state:   "TN",
		pincode: "621313",
	}

	addressMapping := make(map[string]addressMapping)
	addressMapping["subi"] = subiAddr
	fmt.Println(addressMapping["subi"].street, "Subi state")

	fmt.Println(addressMapping)

	groupName := map[string]string{
		"whatsAppStatus": "nothing",
		"telegram":       "yahoo",
		"facebook":       "fb",
		"instagram":      "google",
	}

	for _, value := range groupName {
		fmt.Println(value, "hellossss")
	}

	names := make(map[string]string)
	names["firstName"] = "Dharmadurai"
	names["lastName"] = "Muthusamy"

	printMapKeyVal(names)

	// fmt.Printf("%+v", names)
	// fmt.Printf("%+v", names["lastName"])

	//delete
	// delete(names, "firstName")
	// fmt.Printf("%+v", names)
}

func printMapKeyVal(names map[string]string) {

	for key, value := range names {
		fmt.Println("your", key, "is", value)
	}
}

// age map

func generateAgeMap() map[string]int {

	return map[string]int{
		"dharma":  29,
		"ironman": 23,
		"thor":    27,
	}

}

// education
type Education struct {
	UG          bool
	PG          bool
	companyName string
	salary      int
}

func makeEducationMap() {

	employees := make(map[string]Education)
	p1 := Education{
		UG:          true,
		PG:          true,
		companyName: "techkion",
		salary:      78,
	}
	p2 := Education{
		UG:          true,
		PG:          false,
		companyName: "a8",
		salary:      55,
	}
	employees["Sathish"] = p1
	employees["Dharma"] = p2

	printEducation(employees)

}

func printEducation(edu map[string]Education) {
	for key, value := range edu {
		fmt.Println(key, value)
	}
}
