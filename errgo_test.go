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

// func removeNumbers(s string) string {
// 	rep := strings.NewReplacer(
// 		"1", "_",
// 		"2", "_",
// 		"3", "_",
// 		"4", "_",
// 		"5", "_",
// 		"6", "_",
// 		"7", "_",
// 		"8", "_",
// 		"9", "_",
// 		"0", "_")
// 	return rep.Replace(s)
// }

// func TestMarker(t *testing.T) {
// 	assert := assert.New(t)

// 	m := NewMarker(errorf("BOO"))
// 	assert.Error(m)
// 	assert.Contains(removeNumbers(m.Error()), removeNumbers(MarkHere()))
// }

// func TestErrorMarkf(t *testing.T) {
// 	assert := assert.New(t)
// 	cause := errorf("BOO %d", 10)
// 	em := ErrorMarkf("BOO %d", 10)
// 	if c, ok := em.(interface {
// 		Cause() error
// 	}); ok {
// 		assert.Equal(cause, c.Cause())
// 	}
// 	assert.Contains(removeNumbers(em.Error()), removeNumbers(MarkHere()))
// }
