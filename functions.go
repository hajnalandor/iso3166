package iso3166

import (
	"errors"
	"strings"
)

var (
	ErrInvalidCountryName   = errors.New("invalid country name")
	ErrInvalidCountryAlpha2 = errors.New("invalid country alpha2")
	ErrInvalidCountryAlpha3 = errors.New("invalid country alpha3")
	ErrCountryNotFound      = errors.New("country not found")
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

func ParseSubdivision(countryAlpha2, subdivisionName string) (Subdivision, error) {
	countryAlpha2 = strings.ToUpper(countryAlpha2)
	subdivisionName = strings.ToUpper(subdivisionName)
	validAlpha2 := false

	for _, c := range Countries {
		if c.Alpha2 == countryAlpha2 {
			validAlpha2 = true
			for _, subdivision := range c.Subdivisions {
				if strings.ToUpper(subdivision.Name) == subdivisionName || strings.ToUpper(subdivision.LocalName) == subdivisionName {
					return subdivision, nil
				}
			}
		}
	}
	if !validAlpha2 {
		return Subdivision{},ErrInvalidCountryAlpha2
	}
	return Subdivision{}, ErrSubdivisionNotFound
}
