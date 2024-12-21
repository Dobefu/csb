package init_env

import (
	"testing"

	"github.com/Dobefu/csb/cmd/logger"
)

func TestInitEnv(t *testing.T) {
	Main("../../.env.test")

	logger.SetExitOnFatal(false)

	Main("../../.env.test.bogus")
}
