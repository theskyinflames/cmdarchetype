package cmd

import (
	"time"

	"github.com/theskyinflames/cmdarchetype/config"

	"github.com/sirupsen/logrus"
)

func NewMyCommand(cfg *config.Config, log *logrus.Logger) *MyCommand {
	return &MyCommand{
		cfg: cfg,
		log: log,
	}
}

type MyCommand struct {
	cfg *config.Config
	log *logrus.Logger
}

func (m *MyCommand) DoAction() error {

	// Simulate doing some staff
	time.Sleep(1 * time.Second)

	return nil
}
