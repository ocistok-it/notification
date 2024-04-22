package bootstrap

import (
	"github.com/ocistok-it/notification/bootstrap/deps"
	"github.com/ocistok-it/notification/bootstrap/initiator"
)

const (
	defaultGracefulTimeout = "15s"
)

type bootstrap struct {
	deps *deps.Deps
}

func setup() *bootstrap {
	dependencies := initiator.New().
		InitBasic().
		InitRepository().
		InitUsecase().
		SetupDeps()

	return &bootstrap{deps: dependencies}
}
