package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func stringMethods() {

	var longStr string = "Dharma is person he is a normal human until he has hungry"

	//clone methos
	cloneData := strings.Clone(longStr)
	fmt.Println(unsafe.StringData(longStr) == unsafe.StringData(cloneData)) // both will point different memory location

	//contains
	checkName := strings.Contains(longStr, "Dharma")
	fmt.Println(checkName)

	//contain any
	checkAnyName := strings.ContainsAny(longStr, "zy")
	fmt.Println(checkAnyName)

	fmt.Println(strings.ToUpper(cloneData))
	fmt.Println(strings.ReplaceAll(longStr, "he", "HE"))
	fmt.Println(strings.Join([]string{"Dharma", "durai"}, ","))
	fmt.Println(strings.Split(strings.Join([]string{"Dharma", "durai"}, ","), ","))
}
