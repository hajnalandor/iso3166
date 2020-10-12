package generator

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	readFile("../data/iso3166-1.json")
}

func TestParseFile(t *testing.T) {
	var w CountryWrapper

	parseFile("../data/iso3166-1.json", &w)
}

func TestGenerate(t *testing.T) {
	var sdw SubDivisionWrapper
	parseFile("../../data/iso3166-2.json", &sdw)

	var cw CountryWrapper
	parseFile("../../data/iso3166-1.json", &cw)

	countrySlice := SlicesToMap(cw, sdw)
	GenerateCountryStates(countrySlice)
}

func TestGenerateCountryToAlpha2(t *testing.T) {
	var cw CountryWrapper
	parseFile("../../data/iso3166-1.json", &cw)

	data := getCountryNameToAlpha2Map(cw)
	GenerateCountryToAlpha2(data)
}

//func TestCS(t *testing.T) {
//	t.Log(cs["US"])
//}

func TestSlicesToMap(t *testing.T) {
	var sdw SubDivisionWrapper
	parseFile("../data/iso3166-2.json", &sdw)

	var cw CountryWrapper
	parseFile("../data/iso3166-1.json", &cw)

	countrySlice := SlicesToMap(cw, sdw)
	for k, v := range countrySlice["RS"].Child {
		fmt.Println(k, ":")
		if len(v.Child) != 0 {
			fmt.Printf("\t%v\n", v)
		}
	}
}

func TestGetParentStructure(t *testing.T) {
	var sw SubDivisionWrapper
	parseFile("../data/iso3166-2.json", &sw)
	ps := getParentStructure("US", sw)
	for k, v := range ps {
		fmt.Println(k, ":")
		if len(v) != 0 {
			fmt.Printf("\t%v\n", v)
		}
	}
}
