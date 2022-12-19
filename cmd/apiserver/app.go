package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/smakkking/http-rest-api/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse() // вроде как парсим строка

	var err error

	config := apiserver.NewConfig()

	_, err = toml.DecodeFile(configPath, config)

	s := apiserver.New(config)
	err = s.Start()
	if err != nil {
		log.Fatal()
	}
}
