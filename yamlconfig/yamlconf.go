package yamlconfig

import (
	"errors"
	"flag"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type Conf struct {
	Server ServerConf `yaml:"server"`
}

type ServerConf struct {
	Port    int     `yaml:"port"`
	Host    string  `yaml:"host"`
	TimeOut TimeOut `yaml:"timeout"`
}

type TimeOut struct {
	Read       time.Duration `yaml:"read"`
	ReadHeader time.Duration `yaml:"readHeader"`
	Write      time.Duration `yaml:"write"`
	Idle       time.Duration `yaml:"idle"`
}

func ParseYaml() (Conf, error) {
	fileConf := flag.String("fileconf", "", "name configuration file")

	flag.Parse()

	if *fileConf == "" || fileConf == nil {
		return Conf{}, errors.New("configuration file not defined")
	}

	file, err := os.Open(*fileConf)
	if err != nil {
		return Conf{}, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	cfg := Conf{}

	if err = d.Decode(&cfg); err != nil {
		return Conf{}, err
	}

	return cfg, nil
}
