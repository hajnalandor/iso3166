package iso3166_test

import (
	"fmt"

	"github.com/hajnalandor/iso3166"
)

func exampleCountryCodeToName() {
	country, err := iso3166.ParseCountry("US")
	if err != nil {
		panic(err)
	}
	fmt.Println(country.Name)
	// Output:
	// United States
}

func exampleCountryNameToAlpha2() {
	country, err := iso3166.ParseCountry("United States")
	if err != nil {
		panic(err)
	}
	fmt.Println(country.Alpha2)
	// Output:
	// US
}

func exampleSubDivisionNameToCode() {
	subdivision, err := iso3166.ParseSubdivision( "Bath and North East Somerset","GB")
	if err != nil {
		panic(err)
	}
	fmt.Println(subdivision.Code)
	// Output:
	// BAS
}
