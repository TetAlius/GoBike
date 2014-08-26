package util

import (
	"encoding/json"
	"io/ioutil"
)

var Routes struct {
	Routes []RouteType
}

type RouteType struct {
	Name      string
	Kilometer float64
	Comments  []string
}

func Parse(filePath string) (e error) {
	file, e := ioutil.ReadFile(filePath)
	if e != nil {
		return e
	}
	json.Unmarshal(file, &Routes)
	return
}
