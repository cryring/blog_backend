package main

import (
	"github.com/cryring/blog_backend/internal/config"
	"github.com/cryring/blog_backend/internal/log"
	"github.com/cryring/blog_backend/internal/server"
	"github.com/cryring/blog_backend/internal/system"

	"github.com/alexflint/go-arg"
	"go.uber.org/zap"
)

func main() {
	args := Config{}
	if err := arg.Parse(&args); err != nil {
		log.Errorf("parse args failed: %v", err)
		return
	}

	log.InitLogger("./log/backend.log")

	if _, err := config.Load(args.Config); err != nil {
		log.Errorf("load config file[%s] failed: %v", args.Config, err)
		return
	}

	cfg := server.Config{
		ListenAddr: args.Address,
	}
	srv, err := server.New(cfg)
	if err != nil {
		log.Errorf("http server init failed: %v", err)
		return
	}

	if err := srv.Run(); err != nil {
		log.Errorf("http server start failed: %v", err)
		return
	}

	s := system.WaitForSignal()
	log.Info("signal received, backend closed.", zap.Any("signal", s))
}
