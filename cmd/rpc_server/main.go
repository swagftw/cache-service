package main

import (
	"flag"
	"github.com/swagftw/cache-service/internal/rpc"
)

func main() {
	path := flag.String("config", "./utl/config/config.local.yaml", "path to config file")
	flag.Parse()

	rpc.Start(*path)
}
