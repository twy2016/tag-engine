package main

import (
	"github.com/spf13/pflag"
	"tag-engine/server"
)

var (
	cfg = pflag.StringP("config", "c", "config/config.yaml", "配置文件目录")
)

func main() {
	pflag.Parse()
	engine := server.Init(*cfg)
	server.Run(engine)
}
