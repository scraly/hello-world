package grpc

import (
	"context"
	"fmt"
	"sync"

	// "github.com/scraly/hello-world/internal/services/pkg/v1"

	"go.uber.org/zap"
	"go.zenithar.org/pkg/log"
	"google.golang.org/grpc"

	"github.com/scraly/hello-world/internal/services/pkg/v1/greeter"

	"github.com/scraly/hello-world/cli/hello-world/config"
)

type application struct {
	cfg    *config.Configuration
	server *grpc.Server
}

var (
	app  *application
	once sync.Once
)

// -----------------------------------------------------------------------------

// New initialize the application
func New(ctx context.Context, cfg *config.Configuration) (*grpc.Server, error) {
	var err error

	once.Do(func() {
		// Initialize application
		app = &application{}

		// Apply configuration
		if err := app.ApplyConfiguration(cfg); err != nil {
			log.For(ctx).Fatal("Unable to initialize server settings", zap.Error(err))
		}

		// Initialize Core components
		greeter := greeter.New()
		app.server, err = grpcServer(ctx, cfg, greeter)
	})

	// Return server
	return app.server, err
}

// -----------------------------------------------------------------------------

// ApplyConfiguration apply the configuration after checking it
func (s *application) ApplyConfiguration(cfg interface{}) error {
	// Check configuration validity
	if err := s.checkConfiguration(cfg); err != nil {
		return err
	}

	// Apply to current component (type assertion done if check)
	s.cfg, _ = cfg.(*config.Configuration)

	// No error
	return nil
}

// -----------------------------------------------------------------------------

func (s *application) checkConfiguration(cfg interface{}) error {
	// Check via type assertion
	_, ok := cfg.(*config.Configuration)
	if !ok {
		return fmt.Errorf("server: invalid configuration")
	}

	// No error
	return nil
}
