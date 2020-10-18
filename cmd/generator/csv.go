package generator

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/gocarina/gocsv"
)

type SubDiv struct {
	Type               string `csv:"type"`
	Code               string `csv:"code"`
	Name               string `csv:"name"`
	Local              string `csv:"local"`
	LanguageCode       string `csv:"language"`
	RomanizationSystem string `csv:"rsystem"`
	Parent             string `csv:"parent"`
}

func MustParseCsv(fileName string) []SubDiv {
	parseCsv, err := ParseCsv(fileName)
	if err != nil {
		panic(err)
	}

	return parseCsv
}

func ParseCsv(fileName string) ([]SubDiv, error) {
	csvFile, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("couldn't open the csv file because %v", err)
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))
	var lines []SubDiv

	if err := gocsv.UnmarshalCSV(reader, &lines); err != nil {
		return nil, err
	}

	return cleanLines(lines), nil
}

func cleanLines(csvLines []SubDiv) []SubDiv {
	cLines := make([]SubDiv, 0)
	for _, l := range csvLines {
		l.Code = strings.Replace(l.Code, "*", "", -1)
		cLines = append(cLines, l)
	}

	return cLines
}

func mapToSubDivision(csvLines []SubDiv) SubDivisionWrapper {
	sd := make([]SubDivision, 0)
	for _, l := range csvLines {
		sd = append(sd, SubDivision{
			Name:         l.Name,
			LocalName:    l.Local,
			LanguageCode: l.LanguageCode,
			Code:         l.Code,
			Parent:       l.Parent,
			Type:         l.Type,
		})
	}

	return SubDivisionWrapper{SubDivisions: sd}
}
