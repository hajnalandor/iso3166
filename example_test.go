package iso3166_test

import (
	"fmt"

	"github.com/hajnalandor/iso3166"
)

func ExampleCountryCodeToName() {
	name, err := iso3166.CountryCodeToName("US")
	if err != nil {
		panic(err)
	}
	fmt.Println(name)
	// Output:
	// United States
}

func ExampleCountryNameToAlpha2() {
	name, err := iso3166.CountryNameToAlpha2("United States")
	if err != nil {
		panic(err)
	}
	fmt.Println(name)
	// Output:
	// US
}


func ExampleSubDivisionNameToCode() {
	code, err := iso3166.SubDivisionNameToCode("GB", "Bath and North East Somerset")
	if err != nil {
		panic(err)
	}
	fmt.Println(code)
	// Output:
	// BAS
}


