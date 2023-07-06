package main

import (
	"context"
	"log"

	"github.com/dasiyes/gorel/src/config"
	"github.com/dasiyes/gorel/src/http/server"
	"github.com/dasiyes/gorel/src/router"
)

func main() {

	cnfg, err := config.InitConfig("config.yaml")
	if err != nil {
		log.Fatalf("Error loading config. %vs", err)
	}

	log.Printf("configured Title: %s", cnfg.GetTitle())

	var (
		rt  = router.New(cnfg)
		ctx = context.Background()
		sc  = cnfg.GetSrvCfg()
	)

	_ = server.RunServer(ctx, sc, rt)
}
