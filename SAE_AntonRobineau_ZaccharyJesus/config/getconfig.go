package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

/* 
Get retrieves the contents of the config.json file and stores it in the variable 
General of the config package. Normally you should never change this function. 
*/
func Get(fileName string) {

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error when opening config file: ", err)
	}

	err = json.Unmarshal(content, &General)
	if err != nil {
		log.Fatal("Error when reading config file: ", err)
	}
}
