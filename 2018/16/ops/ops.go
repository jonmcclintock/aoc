package ops

// Operation is a generic interface for an operation
type Operation interface {
	Execute([]int, []int) []int
	String() string
}

// AllOps contains all of the operations.
var AllOps = []Operation{
	OpAddr{},
	OpAddi{},
	OpMulr{},
	OpMuli{},
	OpBanr{},
	OpBani{},
	OpBorr{},
	OpBori{},
	OpSetr{},
	OpSeti{},
	OpGtir{},
	OpGtri{},
	OpGtrr{},
	OpEqir{},
	OpEqri{},
	OpEqrr{},
}

// OpAddr (add register) stores into register C the result of adding register A and register B.
type OpAddr struct{}

func (o OpAddr) Execute(params, regs []int) []int {
	resultRegs := make([]int, len(regs))
	copy(resultRegs, regs)

	resultRegs[params[2]] = regs[params[0]] + regs[params[1]]
	return resultRegs
}

func (o OpAddr) String() string {
	return "addr"
}

// OpAddi (add immediate) stores into register C the result of adding register A and value B.
type OpAddi struct{}

func (o OpAddi) Execute(params, regs []int) []int {
	resultRegs := make([]int, len(regs))
	copy(resultRegs, regs)

	resultRegs[params[2]] = regs[params[0]] + params[1]
	return resultRegs
}

func (o OpAddi) String() string {
	return "addi"
}

// OpMulr (multiply register) stores into register C the result of multiplying register A and register B.
type OpMulr struct{}

func (o OpMulr) Execute(params, regs []int) []int {
	resultRegs := make([]int, len(regs))
	copy(resultRegs, regs)

	resultRegs[params[2]] = regs[params[0]] * regs[params[1]]
	return resultRegs
}

func (o OpMulr) String() string {
	return "mulr"
}

// OpMuli (multiply immediate) stores into register C the result of multiplying register A and value B.
type OpMuli struct{}

func (o OpMuli) Execute(params, regs []int) []int {
	resultRegs := make([]int, len(regs))
	copy(resultRegs, regs)

	resultRegs[params[2]] = regs[params[0]] * params[1]
	return resultRegs
}

func (o OpMuli) String() string {
	return "muli"
}

// OpBanr (bitwise AND register) stores into register C the result of the bitwise AND of register A and register B.
type OpBanr struct{}

func (o OpBanr) Execute(params, regs []int) []int {
	resultRegs := make([]int, len(regs))
	copy(resultRegs, regs)

	resultRegs[params[2]] = regs[params[0]] & regs[params[1]]
	return resultRegs
}

func (o OpBanr) String() string {
	return "banr"
}

// OpBani (bitwise AND immediate) stores into register C the result of the bitwise AND of register A and value B.
type OpBani struct{}

func (o OpBani) Execute(params, regs []int) []int {
	resultRegs := make([]int, len(regs))
	copy(resultRegs, regs)

	resultRegs[params[2]] = regs[params[0]] & params[1]
	return resultRegs
}

func (o OpBani) String() string {
	return "bani"
}

// OpBorr (bitwise OR register) stores into register C the result of the bitwise OR of register A and register B.
type OpBorr struct{}

func (o OpBorr) Execute(params, regs []int) []int {
	resultRegs := make([]int, len(regs))
	copy(resultRegs, regs)

	resultRegs[params[2]] = regs[params[0]] | regs[params[1]]
	return resultRegs
}

func (o OpBorr) String() string {
	return "borr"
}

// OpBori (bitwise OR immediate) stores into register C the result of the bitwise OR of register A and value B.
type OpBori struct{}

func (o OpBori) Execute(params, regs []int) []int {
	resultRegs := make([]int, len(regs))
	copy(resultRegs, regs)

	resultRegs[params[2]] = regs[params[0]] | params[1]
	return resultRegs
}

func (o OpBori) String() string {
	return "bori"
}

// OpSetr (set register) copies the contents of register A into register C. (Input B is ignored.)
type OpSetr struct{}

func (o OpSetr) Execute(params, regs []int) []int {
	resultRegs := make([]int, len(regs))
	copy(resultRegs, regs)

	resultRegs[params[2]] = regs[params[0]]
	return resultRegs
}

func (o OpSetr) String() string {
	return "setr"
}

// OpSeti (set immediate) stores value A into register C. (Input B is ignored.)
type OpSeti struct{}

func (o OpSeti) Execute(params, regs []int) []int {
	resultRegs := make([]int, len(regs))
	copy(resultRegs, regs)

	resultRegs[params[2]] = params[0]
	return resultRegs
}

func (o OpSeti) String() string {
	return "seti"
}

// OpGtir (greater-than immediate/register) sets register C to 1 if value A is greater than register B. Otherwise, register C is set to 0.
type OpGtir struct{}

func (o OpGtir) Execute(params, regs []int) []int {
	resultRegs := make([]int, len(regs))
	copy(resultRegs, regs)

	if params[0] > regs[params[1]] {
		resultRegs[params[2]] = 1
	} else {
		resultRegs[params[2]] = 0
	}
	return resultRegs
}

func (o OpGtir) String() string {
	return "gtir"
}

// OpGtri (greater-than register/immediate) sets register C to 1 if register A is greater than value B. Otherwise, register C is set to 0.
type OpGtri struct{}

func (o OpGtri) Execute(params, regs []int) []int {
	resultRegs := make([]int, len(regs))
	copy(resultRegs, regs)

	if regs[params[0]] > params[1] {
		resultRegs[params[2]] = 1
	} else {
		resultRegs[params[2]] = 0
	}
	return resultRegs
}

func (o OpGtri) String() string {
	return "gtri"
}

// OpGtrr (greater-than register/register) sets register C to 1 if register A is greater than register B. Otherwise, register C is set to 0.
type OpGtrr struct{}

func (o OpGtrr) Execute(params, regs []int) []int {
	resultRegs := make([]int, len(regs))
	copy(resultRegs, regs)

	if regs[params[0]] > regs[params[1]] {
		resultRegs[params[2]] = 1
	} else {
		resultRegs[params[2]] = 0
	}
	return resultRegs
}

func (o OpGtrr) String() string {
	return "gtrr"
}

// OpEqir (equal immediate/register) sets register C to 1 if value A is equal to register B. Otherwise, register C is set to 0.
type OpEqir struct{}

func (o OpEqir) Execute(params, regs []int) []int {
	resultRegs := make([]int, len(regs))
	copy(resultRegs, regs)

	if params[0] == regs[params[1]] {
		resultRegs[params[2]] = 1
	} else {
		resultRegs[params[2]] = 0
	}
	return resultRegs
}

func (o OpEqir) String() string {
	return "eqir"
}

// OpEqri (equal register/immediate) sets register C to 1 if register A is equal to value B. Otherwise, register C is set to 0.
type OpEqri struct{}

func (o OpEqri) Execute(params, regs []int) []int {
	resultRegs := make([]int, len(regs))
	copy(resultRegs, regs)

	if regs[params[0]] == params[1] {
		resultRegs[params[2]] = 1
	} else {
		resultRegs[params[2]] = 0
	}
	return resultRegs
}

func (o OpEqri) String() string {
	return "eqri"
}

// OpEqrr (equal register/register) sets register C to 1 if register A is equal to register B. Otherwise, register C is set to 0.
type OpEqrr struct{}

func (o OpEqrr) Execute(params, regs []int) []int {
	resultRegs := make([]int, len(regs))
	copy(resultRegs, regs)

	if regs[params[0]] == regs[params[1]] {
		resultRegs[params[2]] = 1
	} else {
		resultRegs[params[2]] = 0
	}
	return resultRegs
}

func (o OpEqrr) String() string {
	return "eqrr"
}
