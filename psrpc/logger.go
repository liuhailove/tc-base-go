package psrpc

import (
	"github.com/go-logr/logr"

	"github.com/liuhailove/tc-base-go/psrpc/internal/logger"
)

func SetLogger(l logr.Logger) {
	logger.SetLogger(l)
}
