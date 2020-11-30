# Go library providing ISO 3166 data

ISO 3166-1 is part of the ISO 3166 standard published by the International Organization for Standardization (ISO), and defines codes for the names of countries, dependent territories, and special areas of geographical interest. The official name of the standard is Codes for the representation of names of countries and their subdivisions – Part 1: Country codes. It defines three sets of country codes:

ISO 3166-1 alpha-2 – two-letter country codes which are the most widely used of the three, and used most prominently for the Internet's country code top-level domains (with a few exceptions).

ISO 3166-1 alpha-3 – three-letter country codes which allow a better visual association between the codes and the country names than the alpha-2 codes.

ISO 3166-1 numeric – three-digit country codes which are identical to those developed and maintained by the United Nations Statistics Division, with the advantage of script (writing system) independence, and hence useful for people or systems using non-Latin scripts.

ISO 3166-2 defines subdivision name, code and type of the subdivision 

### Usage

```
go get github.com/hajnalandor/iso3166
```

```
import (
    "github.com/hajnalandor/iso3166"
    )
```

#### Country name to alpha2 code

```
iso3166.CountryNameToAlpha2("United States") // output: US
```

#### Country alpha2 code to name

```
iso3166.CountryAlpha2ToName("US") // output: United States
```

#### Validate country name

```
iso3166.ValidateCountryName("United States") // output: true
```

#### Validate country alpha2 code

```
iso3166.ValidateCountryAlpha2("US") // output: true
```

#### Country alpha2 code to common name

```
iso3166.CountryAlpha2ToCommonName("US") // output: "" United States doesn't have common name
```

#### Country alpha2 code to offical name

```
iso3166.CountryAlpha2ToOfficalName("US") // output: "United States of America"
```

#### Subdivision name to code

```
iso3166.SubdivisonNameToCode("GB", "Bath and North East Somerset") // output: BAS
```

#### Subdivision code to name

```
iso3166.SubdivisonCodeToName("GB", "BAS") // output: Bath and North East Somerset
```

#### Validate subdivison code

```
iso3166.ValidateSubdivisionCode("GB", "BAS") // output: true
```
