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

// CountryNameToAlpha2 returns the countries alpha2 representation
func CountryNameToAlpha2(name string) (string, error) {
	if alpha2, ok := CountryToAlpha2[name]; ok {
		return alpha2, nil
	} else {
		name = strings.ToUpper(name)
		for _, country := range Countries {
			if strings.ToUpper(country.Name) == name || strings.ToUpper(country.OfficialName) == name || strings.ToUpper(country.CommonName) == name {
				return country.Alpha2, nil
			}
		}
	}
	return "", ErrInvalidCountryName
}

// ValidateCountryName is validate whether the country name is a valid name,
// looking in the alpha2 representation, country's name, country's official
// name and country's common name
func ValidateCountryName(name string) bool {
	if _, ok := CountryToAlpha2[name]; ok {
		return true
	} else {
		name = strings.ToUpper(name)
		for _, country := range Countries {
			if strings.ToUpper(country.Name) == name || strings.ToUpper(country.OfficialName) == name || strings.ToUpper(country.CommonName) == name {
				return true
			}
		}
	}

	return false
}

// CountryAlpha2ToName returns the country's name from alpha2 representation
func CountryAlpha2ToName(alpha2 string) (string, error) {
	alpha2 = strings.ToUpper(alpha2)
	for _, country := range Countries {
		if country.Alpha2 == alpha2 {
			return country.Name, nil
		}
	}

	return "", ErrInvalidCountryAlpha2
}

// CountryAlpha2ToOfficalName returns the country's offical name from alpha2 representation
func CountryAlpha2ToOfficialName(alpha2 string) (string, error) {
	alpha2 = strings.ToUpper(alpha2)
	for _, country := range Countries {
		if country.Alpha2 == alpha2 {
			return country.OfficialName, nil
		}
	}

	return "", ErrInvalidCountryAlpha2
}

// CountryAlpha2ToCommonName returns the country's common name from alpha2 representation
func CountryAlpha2ToCommonName(alpha2 string) (string, error) {
	alpha2 = strings.ToUpper(alpha2)
	for _, country := range Countries {
		if country.Alpha2 == alpha2 {
			return country.CommonName, nil
		}
	}

	return "", ErrInvalidCountryAlpha2
}

// ValidateCountryAlpha2 validates the alpha2 representation
func ValidateCountryAlpha2(alpha2 string) bool {
	alpha2 = strings.ToUpper(alpha2)
	for _, country := range Countries {
		if country.Alpha2 == alpha2 {
			return true
		}
	}

	return false
}

// CountryNameToAlpha3 returns the countries alpha3 representation
func CountryNameToAlpha3(name string) (string, error) {
	if alpha3, ok := CountryToAlpha3[name]; ok {
		return alpha3, nil
	} else {
		name = strings.ToUpper(name)
		for _, country := range Countries {
			if strings.ToUpper(country.Name) == name || strings.ToUpper(country.OfficialName) == name || strings.ToUpper(country.CommonName) == name {
				return country.Alpha3, nil
			}
		}
	}
	return "", ErrInvalidCountryName
}

// CountryAlpha3ToName returns the country's name from alpha3 representation
func CountryAlpha3ToName(alpha3 string) (string, error) {
	alpha3 = strings.ToUpper(alpha3)
	for n, alpha3Variant := range CountryToAlpha3 {
		if alpha3Variant == alpha3 {
			return n, nil
		}
	}

	return "", ErrInvalidCountryAlpha3
}

// SubdivisionNameToCode returns the subdivision's code from it's name
func SubdivisionNameToCode(countryAlpha2, subdivisionName string) (string, error) {
	countryAlpha2 = strings.ToUpper(countryAlpha2)
	subdivisionName = strings.ToUpper(subdivisionName)
	if !ValidateCountryAlpha2(countryAlpha2) {
		var err error
		countryAlpha2, err = CountryNameToAlpha2(countryAlpha2)
		if err != nil {
			return "", err
		}
	}
	for _, country := range Countries {
		if country.Alpha2 == countryAlpha2 {
			for _, subdivision := range country.Subdivisions {
				if strings.ToUpper(subdivision.Name) == subdivisionName || strings.ToUpper(subdivision.LocalName) == subdivisionName {
					return subdivision.Code, nil
				}
			}
		}
	}

	return "", ErrInvalidSubDivName
}

// SubdivisionCodeToName returns the subdivision's name from it's code
func SubdivisionCodeToName(countryAlpha2, subdivisionCode string) (string, error) {
	countryAlpha2 = strings.ToUpper(countryAlpha2)
	if !ValidateCountryAlpha2(countryAlpha2) {
		var err error
		countryAlpha2, err = CountryNameToAlpha2(countryAlpha2)
		if err != nil {
			return "", err
		}
	}

	subdivisionCode = strings.ToUpper(subdivisionCode)
	for _, country := range Countries {
		if country.Alpha2 == countryAlpha2 {
			for _, subdivision := range country.Subdivisions {
				if subdivision.Code == subdivisionCode {
					return subdivision.Name, nil
				}
			}
		}
	}

	return "", ErrInvalidSubDivCode
}

// ValidateSubdivisionCode validate the code of the subdivision
func ValidateSubdivisionCode(countryAlpha2, subdivisionCode string) bool {
	countryAlpha2 = strings.ToUpper(countryAlpha2)
	if !ValidateCountryAlpha2(countryAlpha2) {
		var err error
		countryAlpha2, err = CountryNameToAlpha2(countryAlpha2)
		if err != nil {
			return false
		}
	}

	subdivisionCode = strings.ToUpper(subdivisionCode)
	for _, country := range Countries {
		if country.Alpha2 == countryAlpha2 {
			for _, subdivision := range country.Subdivisions {
				if subdivision.Code == subdivisionCode {
					return true
				}
			}
		}
	}
	return false
}

func ValidateSubdivisionName(countryAlpha2, subdivisionName string) bool {
	if _, err := SubdivisionNameToCode(countryAlpha2, subdivisionName); err != nil {
		return false
	}
	return true
}

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

	for _, c := range Countries {
		if c.Alpha2 == countryAlpha2 {
			for _, subdivision := range c.Subdivisions {
				if strings.ToUpper(subdivision.Name) == subdivisionName || strings.ToUpper(subdivision.LocalName) == subdivisionName {
					return subdivision, nil
				}
			}
		}
	}
	return Subdivision{}, ErrSubdivisionNotFound
}
