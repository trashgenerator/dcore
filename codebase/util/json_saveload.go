package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func SaveJSONConfig(object fmt.Stringer) {
	data, err := json.MarshalIndent(object, "  ", "\t")
	if err != nil {
		log.Fatal(err.Error())
	}

	file, err := os.OpenFile(object.String(), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func LoadJSONConfig(object fmt.Stringer) (interface{}, error) {
	file, err := os.OpenFile(object.String(), os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, object); err != nil {
		return nil, err
	}

	return object, nil
}
