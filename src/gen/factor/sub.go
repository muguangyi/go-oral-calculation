package factor

import (
	"fmt"
	"gen"
)

func NewSub(a gen.IFactor, b gen.IFactor) gen.IFactor {
	return &sub{a: a, b: b}
}

type sub struct {
	a gen.IFactor
	b gen.IFactor
}

func (f *sub) Result() int {
	return f.a.Result() - f.b.Result()
}

func (f *sub) Format(t gen.FormatType, withResult bool) string {
	formula := ""

	if t == gen.PH_RAND {
		t = gen.RandFormat()
	}

	if t == gen.PH_LEFT {
		if Poss(0.5) {
			formula += fmt.Sprintf("%s - %s", f.a.Format(t, false), f.b.Format(gen.PH_NONE, false))
		} else {
			formula += fmt.Sprintf("%s - %s", f.a.Format(gen.PH_NONE, false), f.b.Format(t, false))
		}
	} else {

		formula += fmt.Sprintf("%s - %s", f.a.Format(gen.PH_NONE, false), f.b.Format(gen.PH_NONE, false))
	}

	if withResult {
		if t == gen.PH_RIGHT {
			formula += " = (   )"
		} else {
			formula += fmt.Sprintf(" = %2d", f.Result())
		}
	}

	return formula
}
