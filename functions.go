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
		name = strings.ToLower(name)
		for _, country := range CountryStates {
			if strings.ToLower(country.Name) == name || strings.ToLower(country.OfficialName) == name || strings.ToLower(country.CommonName) == name {
				return country.Alpha2, nil
			}
		}
	}
	return "", ErrInvalidCountryName
}

func ValidCountryName(name string) bool {
	if _, ok := CountryToAlpha2[name]; ok {
		return true
	} else {
		name = strings.ToLower(name)
		for _, country := range CountryStates {
			if strings.ToLower(country.Name) == name || strings.ToLower(country.OfficialName) == name || strings.ToLower(country.CommonName) == name {
				return true
			}
		}
	}

	return false
}

func CountryCodeToName(code string) (string, error) {
	code = strings.ToUpper(code)
	if country, ok := CountryStates[code]; ok {
		return country.Name, nil
	}

	return "", ErrInvalidCountryCode
}

func CountryCodeToOfficialName(code string) (string, error) {
	code = strings.ToUpper(code)
	if country, ok := CountryStates[code]; ok {
		return country.OfficialName, nil
	}

	return "", ErrInvalidCountryCode
}

func CountryCodeToCommonName(code string) (string, error) {
	code = strings.ToUpper(code)
	if country, ok := CountryStates[code]; ok {
		return country.CommonName, nil
	}

	return "", ErrInvalidCountryCode
}

func ValidCountryCode(code string) bool {
	code = strings.ToUpper(code)
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
	subDivName = strings.ToLower(subDivName)
	for subDivCode, subDivWrapper := range CountryStates[countryCode].SubDivCodeToName {
		if strings.ToLower(subDivWrapper.Name) == subDivName || strings.ToLower(subDivWrapper.LocalName) == subDivName {
			return subDivCode, nil
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
