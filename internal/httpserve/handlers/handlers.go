package handlers

import (
	"go.uber.org/zap"

	"github.com/datumforge/datum/pkg/datumclient"
)

// Handler contains configuration options for handlers
type Handler struct {
	// IsTest is a flag to determine if the application is running in test mode and will mock external calls
	IsTest bool
	// Logger provides the zap logger to do logging things from the handlers
	Logger *zap.SugaredLogger
	// ReadyChecks is a set of checkFuncs to determine if the application is "ready" upon startup
	ReadyChecks Checks
	// DatumClient is the client to interact with the Datum API
	DatumClient *datumclient.DatumClient
}
