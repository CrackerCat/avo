package gen

import (
	"github.com/mmcloughlin/avo/internal/inst"
)

type godata struct {
}

func NewGoData() Interface {
	return GoFmt(godata{})
}

func (g godata) Generate(is []*inst.Instruction) ([]byte, error) {
	p := &printer{}

	p.Printf("package inst\n\n")

	p.Printf("var Instructions = []Instruction{\n")

	for _, i := range is {
		p.Printf("{\n")

		p.Printf("Opcode: %#v,\n", i.Opcode)
		p.Printf("Summary: %#v,\n", i.Summary)

		p.Printf("Forms: []Form{\n")
		for _, f := range i.Forms {
			p.Printf("{\n")

			if f.ISA != nil {
				p.Printf("ISA: %#v,\n", f.ISA)
			}

			if f.Operands != nil {
				p.Printf("Operands: []Operand{\n")
				for _, op := range f.Operands {
					p.Printf("{Type: %#v, Action: %#v},\n", op.Type, op.Action)
				}
				p.Printf("},\n")
			}

			if f.ImplicitOperands != nil {
				p.Printf("ImplicitOperands: []ImplicitOperand{\n")
				for _, op := range f.ImplicitOperands {
					p.Printf("{Register: %#v, Action: %#v},\n", op.Register, op.Action)
				}
				p.Printf("},\n")
			}

			p.Printf("},\n")
		}
		p.Printf("},\n")

		p.Printf("},\n")
	}

	p.Printf("}\n")

	return p.Result()
}
