package generator

import (
	"encoding/json"
	"io/ioutil"
)

func parseJSONFile(filename string, s interface{}) {
	if err := json.Unmarshal(readFile(filename), &s); err != nil {
		panic(err)
	}
}

func readFile(fileName string) []byte {
	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return dat
}
