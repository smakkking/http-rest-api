package main

// точка входа в приложение

import (
	"flag"
	"fmt"

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
	toml.DecodeFile(configPath, config)

	err = apiserver.Start(config)
	fmt.Print(err)
	//if err != nil {
	//	//log.Fatal("critical error")
	//}
}
