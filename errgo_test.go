package errgo

import (
	// "strings"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func sampleFunc() Loc {
	return Here()
}

func TestHere(t *testing.T) {
	assert := assert.New(t)

	s1 := sampleFunc().String()
	s2 := Here().String()

	assert.Contains(s1, "errgo/errgo_test.go")
	assert.Contains(s1, "github.com/dc0d/errgo.sampleFunc()")
	assert.Contains(s2, "errgo/errgo_test.go")
	assert.Contains(s2, "github.com/dc0d/errgo.TestHere()")
}

func TestMarker(t *testing.T) {
	assert := assert.New(t)

	err := Mark(fmt.Errorf("TEST"))
	s1 := err.Error()

	assert.Contains(s1, "errgo/errgo_test.go")
	assert.Contains(s1, "github.com/dc0d/errgo.TestMarker() TEST")

	c, ok := err.(interface {
		Cause() error
	})
	assert.True(ok)
	assert.Contains(c.Cause().Error(), "TEST")
}

func TestMarkerNil(t *testing.T) {
	assert := assert.New(t)

	err := Mark(nil)
	assert.Nil(err)
}

func TestMarkerf(t *testing.T) {
	assert := assert.New(t)

	s1 := Markf("cnt %v", 10).Error()

	assert.Contains(s1, "errgo/errgo_test.go")
	assert.Contains(s1, "github.com/dc0d/errgo.TestMarkerf() cnt 10")
}
