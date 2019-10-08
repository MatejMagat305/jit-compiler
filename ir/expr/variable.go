package expr

import (
	"github.com/bspaans/jit/asm"
	. "github.com/bspaans/jit/ir/shared"
)

type IR_Variable struct {
	*BaseIRExpression
	Value string
}

func NewIR_Variable(v string) *IR_Variable {
	return &IR_Variable{
		BaseIRExpression: NewBaseIRExpression(Variable),
		Value:            v,
	}
}

func (i *IR_Variable) ReturnType(ctx *IR_Context) Type {
	return ctx.VariableTypes[i.Value]
}

func (i *IR_Variable) String() string {
	return i.Value
}

func (i *IR_Variable) Encode(ctx *IR_Context, target *asm.Register) ([]asm.Instruction, error) {
	var reg *asm.Register
	var result []asm.Instruction
	if i.ReturnType(ctx) == TFloat64 {
		reg = asm.GetFloatingPointRegisterByIndex(ctx.VariableMap[i.Value])
		result = []asm.Instruction{&asm.MOVSD{reg, target}}
	} else {
		reg = asm.Get64BitRegisterByIndex(ctx.VariableMap[i.Value])
		result = []asm.Instruction{&asm.MOV{reg, target}}
	}
	ctx.AddInstructions(result)
	return result, nil
}
