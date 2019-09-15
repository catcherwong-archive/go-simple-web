package config

import (
	"fmt"
	"log"

	"flag"
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

var (
	Cfg *Config
)

var env = flag.String("env", "dev", "env")

func init() {
	flag.Parse()
	f := fmt.Sprintf("config/config.%s.yaml", *env)
	var setting Config
	config, err := ioutil.ReadFile(f)
	if err != nil {
		log.Printf("error when reading file, %s", err.Error())
	}
	yaml.Unmarshal(config, &setting)

	Cfg = &setting

	log.Println(Cfg)
}

type Config struct {
	Env        string           `yaml:"env" json:"env"`
	PostgreSql ConfigPostgreSql `yaml:"postgresql" json:"postgresql"`
}

type ConfigPostgreSql struct {
	Host   string `yaml:"host" json:"host"`
	Port   int    `yaml:"port" json:"port"`
	User   string `yaml:"user" json:"user"`
	Pwd    string `yaml:"pwd" json:"pwd"`
	DbName string `yaml:"dbName" json:"dbName"`
}
