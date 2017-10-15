package utils

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"portchecker/models"
)

func Unmarshall(yamlText []byte) (*models.Config, error) {
	var config models.Config
	var err = yaml.Unmarshal(yamlText, &config)
	if err != nil {
		err_msg := fmt.Sprintf("Error when reading YAML file. Can't create Config Object. Yaml Error: %v\n", err)
		return nil, errors.New(err_msg)
	}

	return &config, nil
}

func UnmarshallFromFile(filePath string) (*models.Config, error) {
	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	config, err := Unmarshall([]byte(data))

	if err != nil {
		return nil, err
	}

	return config, nil
}

func Marshall(in interface{}) string {
	d, err := yaml.Marshal(in)
	result := string(d)
	if err != nil {
		err_msg := fmt.Sprintf("Error when marshalling object {%v}. Err: %v", in, err)
		fmt.Println(err_msg)
		//return nil, errors.New(err_msg)
	}
	return string(result)
}
