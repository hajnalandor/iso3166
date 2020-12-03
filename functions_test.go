package iso3166

import "testing"

func TestParseCountryByName(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "United States",
			want:    "US",
			wantErr: false,
		},
		{
			name:    "Canada",
			want:    "CA",
			wantErr: false,
		},
		{
			name:    "Invalid",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Hungary",
			want:    "HU",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			gotCountry, err := ParseCountry(tt.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountryNameToAlpha2() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if gotCountry.Alpha2 != tt.want {
				t.Errorf("CountryNameToAlpha2() got = %v, want %v", gotCountry.Alpha2, tt.want)
			}
		})
	}
}

func TestParseCountryByAlpha2Code(t *testing.T) {
	tests := []struct {
		code    string
		want    string
		wantErr bool
	}{
		{
			code:    "US",
			want:    "United States",
			wantErr: false,
		},
		{
			code:    "us",
			want:    "United States",
			wantErr: false,
		},
		{
			code:    "CA",
			want:    "Canada",
			wantErr: false,
		},
		{
			code:    "gB",
			want:    "United Kingdom",
			wantErr: false,
		},
		{
			code:    "Invalid",
			want:    "",
			wantErr: true,
		},
		{
			code:    "HU",
			want:    "Hungary",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.code, func(t *testing.T) {
			gotCountry, err := ParseCountry(tt.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountryCodeToName() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if gotCountry.Name != tt.want {
				t.Errorf("CountryCodeToName() got = %v, want %v", gotCountry.Name, tt.want)
			}
		})
	}
}

func TestSubDivisionNameToCode(t *testing.T) {
	tests := []struct {
		countryCode string
		subDivName  string
		want        string
		wantErr     bool
		error		error
	}{
		{
			countryCode: "US",
			subDivName:  "Alabama",
			want:        "AL",
			wantErr:     false,
		},
		{
			countryCode: "US",
			subDivName:  "Alaska",
			want:        "AK",
			wantErr:     false,
		},
		{
			countryCode: "US",
			subDivName:  "American Samoa",
			want:        "AS",
			wantErr:     false,
		},
		{
			countryCode: "US",
			subDivName:  "Arizona",
			want:        "AZ",
			wantErr:     false,
		},
		{
			countryCode: "US",
			subDivName:  "Arkansas",
			want:        "AR",
			wantErr:     false,
		},
		{
			countryCode: "US",
			subDivName:  "California",
			want:        "CA",
			wantErr:     false,
		},
		{
			countryCode: "US",
			subDivName:  "Colorado",
			want:        "CO",
			wantErr:     false,
		},
		{
			countryCode: "US",
			subDivName:  "Connecticut",
			want:        "CT",
			wantErr:     false,
		},
		{
			countryCode: "US",
			subDivName:  "Delaware",
			want:        "DE",
			wantErr:     false,
		},
		{
			countryCode: "US",
			subDivName:  "District Of Columbia",
			want:        "DC",
			wantErr:     false,
		},
		{
			countryCode: "US",
			subDivName:  "Florida",
			want:        "FL",
			wantErr:     false,
		},
		{
			countryCode: "US",
			subDivName:  "Georgia",
			want:        "GA",
			wantErr:     false,
		},
		{
			countryCode: "US",
			subDivName:  "Guam",
			want:        "GU",
			wantErr:     false,
		},
		{
			countryCode: "US",
			subDivName:  "Illinois",
			want:        "IL",
			wantErr:     false,
		},
		{
			countryCode: "US",
			subDivName:  "Indiana",
			want:        "IN",
			wantErr:     false,
		},
		{
			countryCode: "US",
			subDivName:  "Maine",
			want:        "ME",
			wantErr:     false,
		},
		{
			countryCode: "US",
			subDivName:  "New Mexico",
			want:        "NM",
			wantErr:     false,
		},
		{
			countryCode: "US",
			subDivName:  "Oklahoma",
			want:        "OK",
			wantErr:     false,
		},
		{
			countryCode: "US",
			subDivName:  "Texas",
			want:        "TX",
			wantErr:     false,
		},
		{
			countryCode: "CA",
			subDivName:  "Alberta",
			want:        "AB",
			wantErr:     false,
		},
		{
			countryCode: "CA",
			subDivName:  "Nunavut",
			want:        "NU",
			wantErr:     false,
		},
		{
			countryCode: "CA",
			subDivName:  "Ontario",
			want:        "ON",
			wantErr:     false,
		},
		{
			countryCode: "CA",
			subDivName:  "Quebec",
			want:        "QC",
			wantErr:     false,
		},
		{
			countryCode: "CA",
			subDivName:  "Yukon",
			want:        "YT",
			wantErr:     false,
		},
		{
			countryCode: "RS",
			subDivName:  "Srednjebanatski okrug",
			want:        "02",
			wantErr:     false,
		},
		{
			countryCode: "GB",
			subDivName:  "West Berkshire",
			want:        "WBK",
			wantErr:     false,
		},
		{
			countryCode: "INVALID",
			subDivName:  "West Berkshire",
			want:        "",
			wantErr:     true,
			error: ErrInvalidCountryAlpha2,
		},
		{
			countryCode: "GB",
			subDivName:  "West Berkshire something",
			want:        "",
			wantErr:     true,
			error: ErrSubdivisionNotFound,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.countryCode+tt.subDivName, func(t *testing.T) {
			gotSubdivision, err := ParseSubdivision(tt.countryCode, tt.subDivName)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubDivisionNameToCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && tt.error != err {
				t.Errorf("SubDivisionNameToCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSubdivision.Code != tt.want {
				t.Errorf("SubDivisionNameToCode() got = %v, want %v", gotSubdivision.Code, tt.want)
			}
		})
	}
}

func BenchmarkSubDivisionNameToCodeCaseInsensitive(b *testing.B) {
	countryCode := "gB"                          // United Kingdom
	subDivName := "batH and noRth east soMerset" // Bath and North East Somerset
	var subdiv Subdivision
	var err error
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		subdiv, err = ParseSubdivision(countryCode, subDivName)
		if err != nil {
			b.Error(err)
		}
	}
	if subdiv.Code != "BAS" {
		b.Error("invalid subdivision code")
	}
}

func BenchmarkSubDivisionNameToCodeCaseSensitive(b *testing.B) {
	countryCode := "GB"                          // United Kingdom
	subDivName := "Bath and North East Somerset" // Bath and North East Somerset
	var subdiv Subdivision
	var err error
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		subdiv, err = ParseSubdivision(countryCode, subDivName)
		if err != nil {
			b.Error(err)
		}
	}
	if subdiv.Code != "BAS" {
		b.Error("invalid subdivision code")
	}
}
