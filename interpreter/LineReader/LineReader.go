package LineReader

import (
	"github.com/kostyaBro/SICP/interpreter/interfaces"
	"io"
)

type lineReader struct {
	heep interfaces.IHeep
}

func New(heep interfaces.IHeep) interfaces.ILisp {
	return &lineReader{heep: heep}
}

func (lr *lineReader) Add(args ...interface{}) (out interface{}, err error) {
	panic("implement me")
}

func (lr *lineReader) Rem(args ...interface{}) (out interface{}, err error) {
	panic("implement me")
}

func (lr *lineReader) Mul(args ...interface{}) (out interface{}, err error) {
	panic("implement me")
}

func (lr *lineReader) Div(args ...interface{}) (out interface{}, err error) {
	panic("implement me")
}

func (lr *lineReader) Define(name string, body string) (ok bool, err error) {
	panic("implement me")
}

func (lr *lineReader) ApplicativeRun(writer *io.Writer, logger interfaces.ILogger) (output int, err error) {
	panic("implement me")
}

func (lr *lineReader) NormalRun(writer *io.Writer) (output int, err error) {
	panic("implement me")
}

func (lr *lineReader) Case(condition string, logicMap map[string]string) (out interface{}, err error) {
	panic("implement me")
}

func (lr *lineReader) If(condition string, true string, false string) (out interface{}, err error) {
	panic("implement me")
}
