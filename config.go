package logforward

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Port   string
	Module map[string]struct {
		From string
		To   []string
	}
}

var (
	config Config
)

func getOpt() {
	var file = flag.String("c", "conf/conf.toml", "config file path")
	flag.Parse()
	if _, err := toml.DecodeFile(*file, &config); err != nil {
		panic(err)
	}
	log.Printf("%+v\n", config)
}
func (c *Config) getForm(mname string) (f string) {
	if v, ok := c.Module[mname]; ok {
		return v.From
	}
	return
}
func (c *Config) getTo(mname string) (ts []string) {
	if v, ok := c.Module[mname]; ok {
		return v.To
	}
	return
}
