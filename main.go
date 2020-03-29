package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	yaml "gopkg.in/yaml.v2"
)

type Yamlconfig struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func main() {
	filename := flag.String("yaml", "yaml.yml", "yaml file to parse")
	flag.Parse()

	yamlFile, err := ioutil.ReadFile(*filename)
	if err != nil {
		fmt.Printf("Error reading the YAML file: %s\n", err)
		return
	}
	var yamlopen Yamlconfig
	err = yaml.Unmarshal(yamlFile, &yamlopen)
	if err != nil {
		fmt.Printf("Error parsing the Yaml file: %s\n", err)
	}
    fmt.Printf("%v:\n", yamlopen)

}