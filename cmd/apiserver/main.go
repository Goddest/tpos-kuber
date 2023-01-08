package main
//required imports
import (
	"flag"
	"log"
	"github.com/BurntSushi/toml"
	"github.com/Goddest/tpos-kuber/intenal/app/apiserver"
)
var (
	configPath string
)
func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config")
}
func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	//check if decoding config ended with an error
	if err != nil {
		log.Fatal(err)
	}
	//check whether apiserver started normally 
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
