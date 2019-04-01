package interfaces

import (
	"io"
)

type ILisp interface {
	IExpression
	INaming
	ICalculate
	IBranching
}

type IExpression interface {
	Add(args ...interface{}) (out interface{}, err error)
	Rem(args ...interface{}) (out interface{}, err error)
	Mul(args ...interface{}) (out interface{}, err error)
	Div(args ...interface{}) (out interface{}, err error)
}

type INaming interface {
	Define(name string, body string) (ok bool, err error)
}

type ICalculate interface {
	ApplicativeRun(writer *io.Writer, logger ILogger) (output int, err error)
	NormalRun(writer *io.Writer) (output int, err error)
}

type IBranching interface {
	Case(condition string, logicMap map[string]string) (out interface{}, err error)
	// true or false can be null
	If(condition string, true string, false string) (out interface{}, err error)
}
