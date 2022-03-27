package fsfire

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestAbsSuite(t *testing.T) {
	suite.Run(t, &AbsSuite{})
}

type AbsSuite struct {
	suite.Suite
}

func (suite *AbsSuite) Test_AbsPath() {
	grids := []struct {
		dir            string
		create         bool
		expectedExists bool
	}{
		{
			dir:            "test/abs.path",
			create:         true,
			expectedExists: true,
		},
		{
			dir:            "test/abs.path/not.mkdir",
			create:         false,
			expectedExists: false,
		},
	}

	for _, grid := range grids {
		actual, err := AbsPath(grid.dir, grid.create)
		assert.NoError(suite.T(), err)
		assert.Contains(suite.T(), actual, grid.dir)
		if grid.create {
			assert.DirExists(suite.T(), actual)
		} else {
			assert.NoDirExists(suite.T(), actual)
		}
	}
}
