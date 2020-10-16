package generator

import "testing"

func TestReadFile(t *testing.T) {
	readFile("../data/iso3166-1.json")
}

func TestParseFile(t *testing.T) {
	var w CountryWrapper

	MustParseJSONFile("../data/iso3166-1.json", &w)
}

func TestGenerateFromJSON(t *testing.T) {
	var sdw SubDivisionWrapper
	MustParseJSONFile("../../data/iso3166-2.json", &sdw)

	var cw CountryWrapper
	MustParseJSONFile("../../data/iso3166-1.json", &cw)

	countrySlice := SlicesToMap(cw, sdw)
	GenerateCountryStates(countrySlice)
}

func TestGenerateCountryToAlpha2FromJSON(t *testing.T) {
	var cw CountryWrapper
	MustParseJSONFile("../../data/iso3166-1.json", &cw)

	data := getCountryNameToAlpha2Map(cw)
	GenerateCountryToAlpha2(data)
}
