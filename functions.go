package iso3166

import (
	"errors"
	"strings"
)

var (
	ErrInvalidCountry       = errors.New("invalid country")
	ErrInvalidCountryName   = errors.New("invalid country name")
	ErrInvalidCountryAlpha2 = errors.New("invalid country alpha2")
	ErrInvalidCountryAlpha3 = errors.New("invalid country alpha3")
	ErrCountryNotFound      = errors.New("country not found")
	ErrInvalidState         = errors.New("invalid state")
	ErrInvalidSubDivName    = errors.New("invalid state name")
	ErrInvalidSubDivCode    = errors.New("invalid state code")
	ErrSubdivisionNotFound  = errors.New("subdivision not found")
)

func ParseCountry(name string) (Country, error) {
	name = strings.ToUpper(name)
	for _, c := range Countries {
		if strings.ToUpper(c.Name) == name || strings.ToUpper(c.OfficialName) == name ||
			strings.ToUpper(c.CommonName) == name || c.Alpha2 == name || c.Alpha3 == name {
			return c, nil
		}
	}
	return Country{}, ErrCountryNotFound
}

func ParseSubdivision(subdivisionName string, country ...string) (Subdivision, error) {
	var countryName string
	if len(country) == 1 {
		countryName = country[0]
	}
	if subdivisionName == "" {
		return Subdivision{}, ErrInvalidState
	}
	countryName = strings.ToUpper(countryName)
	subdivisionName = strings.ToUpper(subdivisionName)
	validAlpha2 := false

	for _, c := range Countries {
		if countryName == "" || strings.ToUpper(c.Alpha2) == countryName || strings.ToUpper(c.Alpha3) == countryName || strings.ToUpper(c.Name) == countryName {
			validAlpha2 = true
			for _, subdivision := range c.Subdivisions {
				if strings.ToUpper(subdivision.Name) == subdivisionName ||
					strings.ToUpper(subdivision.LocalName) == subdivisionName ||
					strings.ToUpper(subdivision.Code) == subdivisionName {
					return subdivision, nil
				}
			}
		}
	}
	if !validAlpha2 {
		return Subdivision{}, ErrInvalidCountryAlpha2
	}
	return Subdivision{}, ErrSubdivisionNotFound
}

func LookupSubdivision(subdivisionName string, country ...string) ([]Subdivision, error) {
	var countryName string
	if len(country) == 1 {
		countryName = country[0]
	}
	if subdivisionName == "" {
		return []Subdivision{}, ErrInvalidState
	}
	countryName = strings.ToUpper(countryName)
	subdivisionName = strings.ToUpper(subdivisionName)
	validAlpha2 := false

	var subdivisionList []Subdivision
	for _, c := range Countries {
		if countryName == "" || strings.ToUpper(c.Alpha2) == countryName || strings.ToUpper(c.Alpha3) == countryName || strings.ToUpper(c.Name) == countryName {
			validAlpha2 = true
			for _, subdivision := range c.Subdivisions {
				if strings.Contains(strings.ToUpper(subdivision.Name), subdivisionName) ||
					strings.Contains(strings.ToUpper(subdivision.LocalName), subdivisionName) ||
					strings.Contains(strings.ToUpper(subdivision.Code), subdivisionName) {
					subdivisionList = append(subdivisionList, subdivision)
				}
			}
		}
	}
	if !validAlpha2 {
		return []Subdivision{}, ErrInvalidCountryAlpha2
	}
	return subdivisionList, nil
}
