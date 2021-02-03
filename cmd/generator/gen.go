package generator

import (
	"bytes"
	"go/format"
	"os"
	"strings"
	"text/template"
)

type countryWrapper struct {
	Countries []country `json:"3166-1"`
}

type country struct {
	Alpha2        string `json:"alpha_2"`
	Alpha3        string `json:"alpha_3"`
	Name          string `json:"name"`
	NameUppercase string `json:"name_uppercase"`
	OfficialName  string `json:"official_name"`
	CommonName    string `json:"common_name"`
	Numeric       string `json:"numeric"`
	Subdivisions  []subDivision
}

type subDivision struct {
	countryAlpha2Code string
	Name              string
	NameUppercase     string
	Code              string
	LocalName         string
	Type              string
	Parent            string
}

func toUpper(cw countryWrapper) {
	for i := 0; i < len(cw.Countries); i++ {
		cw.Countries[i].Alpha2 = strings.ToUpper(cw.Countries[i].Alpha2)
		cw.Countries[i].Alpha3 = strings.ToUpper(cw.Countries[i].Alpha3)
		cw.Countries[i].NameUppercase = strings.ToUpper(cw.Countries[i].Name)
		cw.Countries[i].CommonName = strings.ToUpper(cw.Countries[i].CommonName)
		cw.Countries[i].OfficialName = strings.ToUpper(cw.Countries[i].OfficialName)
	}
}

func getCountryToAlphaMap(cw countryWrapper) map[string]country {
	countryToAlpha := make(map[string]country)
	for _, c := range cw.Countries {
		countryToAlpha[c.Name] = country{
			Alpha2: c.Alpha2,
			Alpha3: c.Alpha3,
		}
	}
	return countryToAlpha
}

func addSubdivisions(countries []country, subDivisions []subDivision) []country {
	for i, _ := range countries {
		countries[i].Subdivisions = make([]subDivision, 0)
		for _, subDiv := range subDivisions {
			if subDiv.countryAlpha2Code == countries[i].Alpha2 {
				countries[i].Subdivisions = append(countries[i].Subdivisions, subDiv)
			}
		}
	}
	return countries
}

func generateCountryMap(data map[string]country) {
	tmpl, err := template.New("country-to-alpha-generator").Parse(countryToAlphaTmpl)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("../../countries.go")
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		panic(err)
	}
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	_, err = f.Write(formatted)
	if err != nil {
		panic(err)
	}
}

func generateCountries(data []country) {
	tmpl, err := template.New("subdivision-generator").Parse(countrySubDivtmpl)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("../../countries.go")
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		panic(err)
	}
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	_, err = f.Write(formatted)
	if err != nil {
		panic(err)
	}
}
