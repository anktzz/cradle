package gates

type OrGate struct {
	A, B *Wire
	Out  *Wire
}

func NewOrGate(a, b, out *Wire) *OrGate {
	return &OrGate{A: a, B: b, Out: out}
}

func (o *OrGate) Evaluate() {
	o.Out.Set(o.A.Get() && o.B.Get())
}
