package fsfire

import (
	"runtime"
	"strings"
	"testing"

	"github.com/coolstina/expression"
	"github.com/stretchr/testify/assert"
)

func TestGetVersion(t *testing.T) {
	replaced, err := expression.RegularReplaceMatch(`go`, runtime.Version(), "")
	assert.NoError(t, err)
	assert.NotEmpty(t, replaced)

	version, err := GetVersion(replaced)
	assert.NoError(t, err)
	assert.NotNil(t, version)
}

func TestGetFullVersion(t *testing.T) {
	version, err := GetVersion(runtime.Version())
	assert.NoError(t, err)
	assert.NotNil(t, version)

	fullVersion := version.GetFullVersion()
	assert.Len(t, strings.Split(fullVersion, "."), 3)
}

func TestGetMinorVersion(t *testing.T) {
	version, err := GetVersion(runtime.Version())
	assert.NoError(t, err)
	assert.NotNil(t, version)

	minorVersion := version.GetMinorVersion()
	assert.Len(t, strings.Split(minorVersion, "."), 2)
}
