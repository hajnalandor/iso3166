package iso3166

import (
	"errors"
	"strings"
)

var (
	ErrInvalidCountryName   = errors.New("invalid country name")
	ErrInvalidCountryAlpha2 = errors.New("invalid country alpha2")
	ErrInvalidSubDivName    = errors.New("invalid state name")
	ErrInvalidSubDivCode    = errors.New("invalid state code")
)

// CountryNameToAlpha2 returns the countries alpha2 representation
func CountryNameToAlpha2(name string) (string, error) {
	if alpha2, ok := CountryToAlpha2[name]; ok {
		return alpha2, nil
	} else {
		name = strings.ToUpper(name)
		for _, country := range CountryStates {
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
		for _, country := range CountryStates {
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
	if country, ok := CountryStates[alpha2]; ok {
		return country.Name, nil
	}

	return "", ErrInvalidCountryAlpha2
}

// CountryAlpha2ToOfficalName returns the country's offical name from alpha2 representation
func CountryAlpha2ToOfficialName(alpha2 string) (string, error) {
	alpha2 = strings.ToUpper(alpha2)
	if country, ok := CountryStates[alpha2]; ok {
		return country.OfficialName, nil
	}

	return "", ErrInvalidCountryAlpha2
}

// CountryAlpha2ToCommonName returns the country's common name from alpha2 representation
func CountryAlpha2ToCommonName(alpha2 string) (string, error) {
	alpha2 = strings.ToUpper(alpha2)
	if country, ok := CountryStates[alpha2]; ok {
		return country.CommonName, nil
	}

	return "", ErrInvalidCountryAlpha2
}

// ValidateCountryAlpha2 validates the alpha2 representation
func ValidateCountryAlpha2(alpha2 string) bool {
	alpha2 = strings.ToUpper(alpha2)
	_, ok := CountryStates[alpha2]

	return ok
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
	if c, ok := CountryStates[countryAlpha2].SubDivNameToCode[subdivisionName]; ok {
		return c.Code, nil
	}
	for _, subDiv := range CountryStates[countryAlpha2].SubDivNameToCode {
		if codeWrapper, ok := subDiv.SubDivNameToCode[subdivisionName]; ok {
			return codeWrapper.Code, nil
		}
	}
	for subDivCode, subDivWrapper := range CountryStates[countryAlpha2].SubDivCodeToName {
		if subDivWrapper.Name == subdivisionName || subDivWrapper.LocalName == subdivisionName {
			return subDivCode, nil
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
	if c, ok := CountryStates[countryAlpha2].SubDivCodeToName[subdivisionCode]; ok {
		return c.Name, nil
	}
	for _, subDiv := range CountryStates[countryAlpha2].SubDivCodeToName {
		if codeWrapper, ok := subDiv.SubDivCodeToName[subdivisionCode]; ok {
			return codeWrapper.Name, nil
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
	if _, ok := CountryStates[countryAlpha2].SubDivCodeToName[subdivisionCode]; ok {
		return true
	}
	for _, subDiv := range CountryStates[countryAlpha2].SubDivCodeToName {
		if _, ok := subDiv.SubDivCodeToName[subdivisionCode]; ok {
			return true
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
