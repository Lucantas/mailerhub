package config

import (
	"fmt"
	"io/ioutil"
	"os"
)

func readJSON(path string) []byte {
	file, err := os.Open(path)

	if err != nil {
		// TODO: properly handle if error on this part
		fmt.Println(err)
	}

	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)

	if err != nil {
		// TODO: properly handle if error on this part
		fmt.Println(err)
	}

	return byteValue
}
