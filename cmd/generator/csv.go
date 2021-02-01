package generator

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/gocarina/gocsv"
)

type csvLine struct {
	Type   string `csv:"type"`
	Code   string `csv:"code"`
	Name   string `csv:"name"`
	Local  string `csv:"local"`
	Parent string `csv:"parent"`
}

func mustParseCsv(fileName string) []csvLine {
	parseCsv, err := parseCsv(fileName)
	if err != nil {
		panic(err)
	}

	return parseCsv
}

func parseCsv(fileName string) ([]csvLine, error) {
	csvFile, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("couldn't open the csv file because %v", err)
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))
	var lines []csvLine

	if err := gocsv.UnmarshalCSV(reader, &lines); err != nil {
		return nil, err
	}

	return cleanLines(lines), nil
}

func cleanLines(csvLines []csvLine) []csvLine {
	lines := make([]csvLine, 0)
	for _, l := range csvLines {
		l.Code = strings.Replace(l.Code, "*", "", -1)
		lines = append(lines, l)
	}

	return lines
}

func csvToSubdivision(csvLines []csvLine) []subDivision {
	csvLines = cleanLines(csvLines)

	sd := make([]subDivision, 0)
	for _, l := range csvLines {
		splittedSubdivisionCode := strings.Split(l.Code, "-")
		parentCode := ""
		if s := strings.Split(l.Parent, "-"); len(s) == 2 {
			parentCode = s[1]
		}
		sd = append(sd, subDivision{
			Name:              l.Name,
			countryAlpha2Code: splittedSubdivisionCode[0],
			Code:              splittedSubdivisionCode[1],
			LocalName:         l.Local,
			Type:              l.Type,
			Parent:            parentCode,
		})
	}
	return sd
}
