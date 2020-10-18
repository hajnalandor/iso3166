package iso3166

import (
	"errors"
	"strings"
)

var (
	ErrInvalidCountryName = errors.New("invalid country name")
	ErrInvalidCountryCode = errors.New("invalid country code")
	ErrInvalidSubDivName  = errors.New("invalid state name")
	ErrInvalidSubDivCode  = errors.New("invalid state code")
)

func CountryNameToAlpha2(name string) (string, error) {
	if alpha2, ok := CountryToAlpha2[name]; ok {
		return alpha2, nil
	} else {
		return "", ErrInvalidCountryName
	}
}

func ValidCountryName(name string) bool {
	_, ok := CountryToAlpha2[name]

	return ok
}

func CountryCodeToName(code string) (string, error) {
	if country, ok := CountryStates[code]; ok {
		return country.Name, nil
	}

	return "", ErrInvalidCountryCode
}

func ValidCountryCode(code string) bool {
	_, ok := CountryStates[code]

	return ok
}

func SubDivisionNameToCode(countryCode, subDivName string) (string, error) {
	countryCode = strings.ToUpper(countryCode)
	if !ValidCountryCode(countryCode) {
		var err error
		countryCode, err = CountryNameToAlpha2(countryCode)
		if err != nil {
			return "", err
		}
	}
	if c, ok := CountryStates[countryCode].SubDivNameToCode[subDivName]; ok {
		return c.Code, nil
	}
	for _, subDiv := range CountryStates[countryCode].SubDivNameToCode {
		if codeWrapper, ok := subDiv.SubDivNameToCode[subDivName]; ok {
			return codeWrapper.Code, nil
		}
	}

	return "", ErrInvalidSubDivName
}

func SubDivisionCodeToName(countryCode, subDivCode string) (string, error) {
	countryCode = strings.ToUpper(countryCode)
	if !ValidCountryCode(countryCode) {
		var err error
		countryCode, err = CountryNameToAlpha2(countryCode)
		if err != nil {
			return "", err
		}
	}
	if c, ok := CountryStates[countryCode].SubDivCodeToName[subDivCode]; ok {
		return c.Name, nil
	}
	for _, subDiv := range CountryStates[countryCode].SubDivCodeToName {
		if codeWrapper, ok := subDiv.SubDivCodeToName[subDivCode]; ok {
			return codeWrapper.Name, nil
		}
	}

	return "", ErrInvalidSubDivCode
}
