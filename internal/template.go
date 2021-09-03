package internal

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Template struct {
	Game     string            `yaml:"game"`
	Image    string            `yaml:"image"`
	Env      map[string]string `yaml:"env"`
	Cli      map[string]string `yaml:"cli,omitempty"`
	RAM      int               `yaml:"ram,omitempty"`
	CPU      float64           `yaml:"cpu"`
	GamePort int               `yaml:"game_port"`
}

func LoadTemplate(path string) (*Template, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		return nil, err
	}
	template := new(Template)
	err = yaml.Unmarshal(data, template)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return nil, err
	}
	return template, nil
}
