package config

import (
	"gift2grow_backend/utils/logger"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var C = &config{}

func init() {
	yml, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		logger.Log(logrus.Fatal, "UNABLE TO READ YAML CONFIGURATION FILE")
	}
	err = yaml.Unmarshal(yml, C)
	if err != nil {
		logger.Log(logrus.Fatal, "UNABLE TO PARSE YAML CONFIGURATION FILE")
	}
	// Apply configurations
	logrus.SetLevel(logrus.WarnLevel)
}
