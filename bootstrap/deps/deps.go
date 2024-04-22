package deps

import (
	"context"
	"github.com/ocistok-it/notification/internal/infrastructure/config"
	"github.com/rs/zerolog/log"
	"sync"
)

type (
	Deps struct {
		Config  *config.Main
		Basic   *Basic
		Usecase *Usecase
	}

	StopperFn func(ctx context.Context) error
)

func (d *Deps) Close(ctx context.Context) {
	stoppers := []StopperFn{
		func(ctx context.Context) error {
			return d.Basic.Consumer.Close()
		},
		//func(ctx context.Context) error {
		//	return d.Basic.Mailer.Close()
		//},
	}

	d.stopper(ctx, stoppers)
}

func (d *Deps) stopper(ctx context.Context, resources []StopperFn) {
	wg := sync.WaitGroup{}
	wg.Add(len(resources))

	for i := range resources {
		stopper := resources[i]
		go func() {
			defer wg.Done()
			if err := stopper(ctx); err != nil {
				log.Err(err).Msg("failed to stop dependencies")
			}
		}()
	}

	wg.Wait()
}
