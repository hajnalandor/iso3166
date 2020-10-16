package iso3166

import "testing"

func TestCountryNameToAlpha2(t *testing.T) {
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
			got, err := CountryNameToAlpha2(tt.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountryNameToAlpha2() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if got != tt.want {
				t.Errorf("CountryNameToAlpha2() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidCountryName(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "United States",
			want: false,
		},
		{
			name: "Canada",
			want: false,
		},
		{
			name: "Invalid",
			want: true,
		},
		{
			name: "Hungary",
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidCountryName(tt.name); got != tt.want {
				t.Errorf("ValidCountryName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountryCodeToName(t *testing.T) {
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
			got, err := CountryCodeToName(tt.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountryCodeToName() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if got != tt.want {
				t.Errorf("CountryCodeToName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidCountryCode(t *testing.T) {
	tests := []struct {
		code string
		want bool
	}{
		{
			code: "US",
			want: false,
		},
		{
			code: "us",
			want: false,
		},
		{
			code: "CA",
			want: false,
		},
		{
			code: "Invalid",
			want: true,
		},
		{
			code: "HU",
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.code, func(t *testing.T) {
			if got := ValidCountryCode(tt.code); got != tt.want {
				t.Errorf("ValidCountryCode() = %v, want %v", got, tt.want)
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
			subDivName:  "Federated States Of Micronesia",
			want:        "FM",
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
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.countryCode+tt.subDivName, func(t *testing.T) {
			got, err := SubDivisionNameToCode(tt.countryCode, tt.subDivName)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubDivisionNameToCode() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if got != tt.want {
				t.Errorf("SubDivisionNameToCode() got = %v, want %v", got, tt.want)
			}
		})
	}
}
