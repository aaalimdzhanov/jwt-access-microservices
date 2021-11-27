package main

import (
	"flag"
	"github.com/aaalimdzhanov/jwt-access-microservices/common"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/aaalimdzhanov/jwt-access-microservices/internal/app/apiserver"
)
var (
	configPath string
)

func init(){
	flag.StringVar(&configPath,"config-path","configs/apiserver.toml","path to config file")
}

func main() {
	flag.Parse()
	config := apiserver.NewConfig()
	common.InitKeys()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil{
		log.Fatal(err)
	}
	
	
	if err := apiserver.Start(config); err != nil{
		log.Fatal(err)
	}
}