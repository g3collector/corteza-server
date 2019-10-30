package external

import (
	"go.uber.org/zap"

	"github.com/cortezaproject/corteza-server/pkg/logger"
	"github.com/cortezaproject/corteza-server/system/service"
)

const (
	OIDC_PROVIDER_PREFIX = "openid-connect."
)

func Init() {
	setupGoth(service.CurrentSettings)
}

func log() *zap.Logger {
	return logger.Default().Named("auth.external")
}
