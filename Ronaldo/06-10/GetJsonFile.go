package main

import (
	"encoding/json"
	"io"
	"os"
)
func main(){

}
type Passager struct{
	name int 
}
func GetJsonFile(filepath string, jStorw Passager) ([]Passager, error) {
	var arrStru []Passager
	data, err := os.Open(filepath)
	if err != nil {
		return arrStru,err
	}
	bytes,err:= io.ReadAll(data)
	if err != nil {
		return arrStru,err
	}
	err = json.Unmarshal(bytes, &arrStru)
	if err != nil {
		return arrStru,err
	}
	return arrStru, nil
}
