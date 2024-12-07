package color

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSprintColor(t *testing.T) {
	assert.Equal(t, SprintColor(FgWhite, BgBlack, "test"), "\x1b[97;50mtest\x1b[0;0m")
	assert.Equal(t, SprintColor(FgLightRed, BgLightYellow, "test"), "\x1b[91;103mtest\x1b[0;0m")
}
