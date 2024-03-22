package utils

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

var pu PathUtils

type UtilsPathTestSuite struct {
	suite.Suite
}

func TestUtilsPathTestSuite(t *testing.T) {
	suite.Run(t, &UtilsPathTestSuite{})
}

func (ups *UtilsPathTestSuite) TestUrlCorrespondsToPathWithoutVariables() {
	pu = NewPathUtils("/test/whaaat")

	ups.True(pu.UrlCorrespondsToPath("/test/whaaat"))
	ups.False(pu.UrlCorrespondsToPath("/nope"))
	ups.False(pu.UrlCorrespondsToPath("/test/nope"))
}

func (ups *UtilsPathTestSuite) TestUrlCorrespondsToPathWithVariables() {
	pu = NewPathUtils("/test/:id")

	ups.True(pu.UrlCorrespondsToPath("/test/1"))
	ups.True(pu.UrlCorrespondsToPath("/test/yup"))
	ups.False(pu.UrlCorrespondsToPath("/nope"))
	ups.False(pu.UrlCorrespondsToPath("/test/1/321"))
}

func (ups *UtilsPathTestSuite) TestUrlCorrespondsToPathWithMultipleVariables() {
	pu = NewPathUtils("/test/:id/:anotherId/justdoit")

	ups.True(pu.UrlCorrespondsToPath("/test/1/321/justdoit"))
	ups.True(pu.UrlCorrespondsToPath("/test/bb/aa/justdoit"))
	ups.False(pu.UrlCorrespondsToPath("/test/1/321"))
}
