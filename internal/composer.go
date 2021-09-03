package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type ComposeFile struct {
	Version string `yaml:"version"`

	Services map[string]Service `yaml:"services"`
}

type Service struct {
	Image         string            `yaml:"image"`
	ContainerName string            `yaml:"container_name"`
	Ports         []string          `yaml:"ports,omitempty"`
	Environment   map[string]string `yaml:"environment,omitempty"`
	Tty           bool              `yaml:"tty"`
	SdtinOpen     bool              `yaml:"stdin_open"`
	Restart       string            `yaml:"restart"`
	RAM           string            `yaml:"mem_limit"`
	CPU           float64           `yaml:"cpus"`
}

func ConvertToComposeFile(gs *GameServer) error {
	// Mapping
	cmp := new(ComposeFile)
	cmp.Version = "3"
	cmp.Services = make(map[string]Service)
	cmp.Services[gs.Name] = Service{
		Image:         gs.Template.Image,
		ContainerName: gs.Name,
		Ports:         []string{fmt.Sprintf("%d:%d", gs.Template.GamePort, gs.Template.GamePort)},
		Environment:   gs.Template.Env,
		Tty:           true,
		SdtinOpen:     true,
		Restart:       "unless-stopped",
		RAM:           fmt.Sprint(gs.Template.RAM) + "m",
		CPU:           gs.Template.CPU,
	}

	// Convert to yaml
	data, err := yaml.Marshal(cmp)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Save file
	err = ioutil.WriteFile(os.Getenv("SERVER_DIR")+gs.Name+".compose.yaml", data, 0)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
