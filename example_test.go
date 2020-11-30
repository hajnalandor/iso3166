package iso3166_test

import (
	"fmt"

	"github.com/hajnalandor/iso3166"
)

func exampleCountryCodeToName() {
	name, err := iso3166.CountryAlpha2ToName("US")
	if err != nil {
		panic(err)
	}
	fmt.Println(name)
	// Output:
	// United States
}

func exampleCountryNameToAlpha2() {
	name, err := iso3166.CountryNameToAlpha2("United States")
	if err != nil {
		panic(err)
	}
	fmt.Println(name)
	// Output:
	// US
}

func exampleSubDivisionNameToCode() {
	code, err := iso3166.SubdivisionNameToCode("GB", "Bath and North East Somerset")
	if err != nil {
		panic(err)
	}
	fmt.Println(code)
	// Output:
	// BAS
}
