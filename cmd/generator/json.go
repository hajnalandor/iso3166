package generator

import (
	"encoding/json"
	"io/ioutil"
)

func MustParseJSONFile(filename string, s interface{}) {
	file, err := readFile(filename)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(file, &s); err != nil {
		panic(err)
	}
}

func readFile(fileName string) ([]byte, error) {
	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return dat, nil
}
