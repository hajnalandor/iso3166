package generator

import "testing"

func TestGenerateFromCSV(t *testing.T) {
	csv := ParseCsv("../../data/iso3166-2.csv")
	csv = cleanLines(csv)
	subDivisionWrapper := mapToSubDivision(csv)

	var cw CountryWrapper
	parseJSONFile("../../data/iso3166-1.json", &cw)

	countrySlice := SlicesToMap(cw, subDivisionWrapper)
	GenerateCountryStates(countrySlice)
}
