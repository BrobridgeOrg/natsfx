package natsfx

import (
	"context"

	"github.com/nats-io/nats-server/v2/server"
	"go.uber.org/fx"
)

var Module = fx.Module("nats-server",
	fx.Provide(
		New,
	),
	fx.Invoke(StartServer),
)

type Params struct {
	fx.In

	Lifecycle fx.Lifecycle
}

type Result struct {
	fx.Out

	Server *server.Server
}

func New(p Params) (Result, error) {
	var ns *server.Server

	p.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			opts := &server.Options{
				JetStream: true,
				StoreDir:  "./data",
			}

			var err error
			ns, err = server.NewServer(opts)

			if err != nil {
				return err
			}

			ns.ConfigureLogger()
			ns.Start()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			ns.Shutdown()
			return nil
		},
	})

	return Result{Server: ns}, nil
}

func StartServer(server *server.Server) error {

	return nil
}
