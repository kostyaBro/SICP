package LineReader

import (
	"github.com/kostyaBro/SICP/interpreter/interfaces"
)

type processor struct {
}

func NewProcessor(lisp interfaces.ILisp) interfaces.IProcessor {
	return new(processor)
}

func (p *processor) Process() {
	// todo switch for run p
}
