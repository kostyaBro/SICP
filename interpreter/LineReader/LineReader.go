package LineReader

import (
	"fmt"
	"github.com/kostyaBro/SICP/interpreter/interfaces"
	"github.com/pkg/errors"
	"io"
	"reflect"
)

type lineReader struct {
	heep interfaces.IHeep
}

func New(heep interfaces.IHeep) interfaces.ILisp {
	return &lineReader{heep: heep}
}

func (lr *lineReader) Add(args ...interface{}) (out interface{}, err error) {
	if len(args) == 0 {
		return 0, nil
	}
	if reflect.ValueOf(args[0]).Kind() >= reflect.Int && reflect.ValueOf(args[0]).Kind() <= reflect.Uint64 {
		var answer int64
		for _, arg := range args {
			answer += reflect.ValueOf(arg).Int()
		}
		return answer, nil
	}
	if reflect.ValueOf(args[0]).Kind() == reflect.Float32 || reflect.ValueOf(args[0]).Kind() == reflect.Float64 {
		var answer float64
		for _, arg := range args {
			answer += reflect.ValueOf(arg).Float()
		}
		return answer, nil
	}
	return 0, errors.New(fmt.Sprintf("args not work %+v", args))
}

func (lr *lineReader) Rem(args ...interface{}) (out interface{}, err error) {
	if len(args) == 0 {
		return 0, nil
	}
	if reflect.ValueOf(args[0]).Kind() >= reflect.Int && reflect.ValueOf(args[0]).Kind() <= reflect.Uint64 {
		var answer int64
		for _, arg := range args {
			answer -= reflect.ValueOf(arg).Int()
		}
		return answer, nil
	}
	if reflect.ValueOf(args[0]).Kind() == reflect.Float32 || reflect.ValueOf(args[0]).Kind() == reflect.Float64 {
		var answer float64
		for _, arg := range args {
			answer -= reflect.ValueOf(arg).Float()
		}
		return answer, nil
	}
	return 0, errors.New(fmt.Sprintf("args not work %+v", args))
}

func (lr *lineReader) Mul(args ...interface{}) (out interface{}, err error) {
	if len(args) == 0 {
		return 0, nil
	}
	if reflect.ValueOf(args[0]).Kind() >= reflect.Int && reflect.ValueOf(args[0]).Kind() <= reflect.Uint64 {
		var answer int64
		for _, arg := range args {
			answer *= reflect.ValueOf(arg).Int()
		}
		return answer, nil
	}
	if reflect.ValueOf(args[0]).Kind() == reflect.Float32 || reflect.ValueOf(args[0]).Kind() == reflect.Float64 {
		var answer float64
		for _, arg := range args {
			answer *= reflect.ValueOf(arg).Float()
		}
		return answer, nil
	}
	return 0, errors.New(fmt.Sprintf("args not work %+v", args))
}

func (lr *lineReader) Div(args ...interface{}) (out interface{}, err error) {
	if len(args) == 0 {
		return 0, nil
	}
	if reflect.ValueOf(args[0]).Kind() >= reflect.Int && reflect.ValueOf(args[0]).Kind() <= reflect.Uint64 {
		var answer int64
		for _, arg := range args {
			answer /= reflect.ValueOf(arg).Int()
		}
		return answer, nil
	}
	if reflect.ValueOf(args[0]).Kind() == reflect.Float32 || reflect.ValueOf(args[0]).Kind() == reflect.Float64 {
		var answer float64
		for _, arg := range args {
			answer /= reflect.ValueOf(arg).Float()
		}
		return answer, nil
	}
	return 0, errors.New(fmt.Sprintf("args not work %+v", args))
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
