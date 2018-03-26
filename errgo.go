package errgo

import (
	"fmt"
	"path/filepath"
	"runtime"
)

//-----------------------------------------------------------------------------

// Here info about code location with a string representation formatted as
// <dir>/<file>.go@<line>:<package>.<function>()
func Here(skip ...int) Loc {
	sk := 1
	if len(skip) > 0 && skip[0] > 1 {
		sk = skip[0]
	}
	pc, fileName, fileLine, ok := runtime.Caller(sk)
	fn := runtime.FuncForPC(pc)
	var res Loc
	defer func() {
		if res.str != "" {
			return
		}
		res.str = res.FuncName
	}()
	if !ok {
		res.FuncName = "N/A"
		return res
	}
	res.FileName = fileName
	res.FileLine = fileLine
	res.FuncName = fn.Name()
	fileName = filepath.Join(filepath.Base(filepath.Dir(fileName)), filepath.Base(fileName))
	res.str = fmt.Sprintf("%s@%d:%s()", fileName, res.FileLine, res.FuncName)
	return res
}

// Loc info about code location with a string representation formatted as
// <dir>/<file>.go@<line>:<package>.<function>()
type Loc struct {
	FuncName string
	FileName string
	FileLine int

	str string
}

func (l Loc) String() string {
	return l.str
}

//-----------------------------------------------------------------------------

type marker struct {
	loc Loc
	err error
}

func (m *marker) Error() string {
	var cause string
	if m.err != nil {
		cause = m.err.Error()
	} else {
		cause = "CAUSE:N/A"
	}
	return m.loc.String() + " " + cause
}

// Cause implements Causer interface
func (m *marker) Cause() error { return m.err }

// Mark ...
func Mark(cause error) error {
	if cause == nil {
		return nil
	}
	return &marker{
		loc: Here(2),
		err: cause,
	}
}

//-----------------------------------------------------------------------------

// Markf ...
func Markf(format string, v ...interface{}) error {
	return fmt.Errorf(Here(2).String()+" "+format, v...)
}
