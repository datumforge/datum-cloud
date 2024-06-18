package serveropts

import (
	echoprometheus "github.com/datumforge/echo-prometheus/v5"
	echo "github.com/datumforge/echox"
	"github.com/datumforge/echox/middleware"
	"github.com/datumforge/echozap"
	"go.uber.org/zap"

	"github.com/datumforge/datum-cloud/internal/datum"
	"github.com/datumforge/datum-cloud/internal/httpserve/config"

	"github.com/datumforge/datum/pkg/middleware/cachecontrol"
	"github.com/datumforge/datum/pkg/middleware/cors"
	"github.com/datumforge/datum/pkg/middleware/echocontext"
	"github.com/datumforge/datum/pkg/middleware/mime"
	"github.com/datumforge/datum/pkg/middleware/ratelimit"
	"github.com/datumforge/datum/pkg/middleware/redirect"
	"github.com/datumforge/datum/pkg/middleware/sentry"
)

type ServerOption interface {
	apply(*ServerOptions)
}

type applyFunc struct {
	applyInternal func(*ServerOptions)
}

func (fso *applyFunc) apply(s *ServerOptions) {
	fso.applyInternal(s)
}

func newApplyFunc(apply func(option *ServerOptions)) *applyFunc {
	return &applyFunc{
		applyInternal: apply,
	}
}

// WithConfigProvider supplies the config for the server
func WithConfigProvider(cfgProvider config.ConfigProvider) ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		s.ConfigProvider = cfgProvider
	})
}

// WithLogger supplies the logger for the server
func WithLogger(l *zap.SugaredLogger) ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		// Add logger to main config
		s.Config.Logger = l
		// Add logger to the handlers config
		s.Config.Handler.Logger = l
	})
}

// WithDatumClient supplies the datum client for the server
func WithDatumClient() ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		// Add logger to main config
		config, err := datum.NewDefaultConfig()
		if err != nil {
			s.Config.Logger.Fatalw("failed to create datum client", "error", err)
		}

		config.Token = s.Config.Settings.Server.Datum.Token

		s.Config.Handler.DatumClient, err = config.NewClient()
		if err != nil {
			s.Config.Logger.Fatalw("failed to create datum client", "error", err)
		}
	})
}

// WithHTTPS sets up TLS config settings for the server
func WithHTTPS() ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		if !s.Config.Settings.Server.TLS.Enabled {
			// this is set to enabled by WithServer
			// if TLS is not enabled, move on
			return
		}

		s.Config.WithTLSDefaults()

		if !s.Config.Settings.Server.TLS.AutoCert {
			s.Config.WithTLSCerts(s.Config.Settings.Server.TLS.CertFile, s.Config.Settings.Server.TLS.CertKey)
		}
	})
}

// WithMiddleware adds the middleware to the server
func WithMiddleware() ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		// Initialize middleware if null
		if s.Config.DefaultMiddleware == nil {
			s.Config.DefaultMiddleware = []echo.MiddlewareFunc{}
		}

		// default middleware
		s.Config.DefaultMiddleware = append(s.Config.DefaultMiddleware,
			middleware.RequestID(), // add request id
			middleware.Recover(),   // recover server from any panic/fatal error gracefully
			middleware.LoggerWithConfig(middleware.LoggerConfig{
				Format: "remote_ip=${remote_ip}, method=${method}, uri=${uri}, status=${status}, session=${header:Set-Cookie}, host=${host}, referer=${referer}, user_agent=${user_agent}, route=${route}, path=${path}, auth=${header:Authorization}\n",
			}),
			sentry.New(),
			echoprometheus.MetricsMiddleware(),                   // add prometheus metrics
			echozap.ZapLogger(s.Config.Logger.Desugar()),         // add zap logger, middleware requires the "regular" zap logger
			echocontext.EchoContextToContextMiddleware(),         // adds echo context to parent
			cors.New(s.Config.Settings.Server.CORS.AllowOrigins), // add cors middleware
			mime.NewWithConfig(mime.Config{DefaultContentType: echo.MIMEApplicationJSONCharsetUTF8}), // add mime middleware
			cachecontrol.New(),                        // add cache control middleware
			middleware.Secure(),                       // add XSS middleware
			redirect.NewWithConfig(redirect.Config{}), // add redirect middleware
		)
	})
}

// WithRateLimiter sets up the rate limiter for the server
func WithRateLimiter() ServerOption {
	return newApplyFunc(func(s *ServerOptions) {
		if s.Config.Settings.Ratelimit.Enabled {
			s.Config.DefaultMiddleware = append(s.Config.DefaultMiddleware, ratelimit.RateLimiterWithConfig(&s.Config.Settings.Ratelimit))
		}
	})
}
