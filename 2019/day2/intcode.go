package day2

// IntcodeComputer func
type IntcodeComputer struct {
	Intergers []int
	Inputs    []int
	Output    *int
	Offset    int
	Extend    map[int]int
}
type ModeType int

const (
	PositionMode int = 0
	ImediateMode     = 1
	RelativeMode     = 2
)

// Init func
func Init(intergers []int, inputs []int, output *int) *IntcodeComputer {
	return &IntcodeComputer{
		Intergers: intergers,
		Inputs:    inputs,
		Output:    output,
		Offset:    0,
		Extend:    make(map[int]int),
	}
}

// Run func
func (ic *IntcodeComputer) Run() {
	i := 0
	for {
		op, paramsMode := getMode(ic.Get(i, 1))
		switch op {
		case 99:
			return
		case 1:
			ic.Set(i+3, (paramsMode/100)%10, ic.Get(i+1, paramsMode%10)+ic.Get(i+2, (paramsMode/10)%10))
			i += 4
		case 2:
			ic.Set(i+3, (paramsMode/100)%10, ic.Get(i+1, paramsMode%10)*ic.Get(i+2, (paramsMode/10)%10))
			i += 4
		case 3:
			ic.Set(i+1, paramsMode%10, ic.Inputs[0])
			ic.Inputs = ic.Inputs[1:]
			i += 2
		case 4:
			*ic.Output = ic.Get(i+1, paramsMode%10)
			i += 2
		case 5:
			if ic.Get(i+1, paramsMode%10) != 0 {
				i = ic.Get(i+2, (paramsMode/10)%10)

			} else {
				i += 3
			}
		case 6:
			if ic.Get(i+1, paramsMode%10) == 0 {
				i = ic.Get(i+2, (paramsMode/10)%10)
			} else {
				i += 3
			}
		case 7:
			if ic.Get(i+1, paramsMode%10) < ic.Get(i+2, (paramsMode/10)%10) {
				ic.Set(i+3, (paramsMode/100)%10, 1)
			} else {
				ic.Set(i+3, (paramsMode/100)%10, 0)
			}
			i += 4
		case 8:
			if ic.Get(i+1, paramsMode%10) == ic.Get(i+2, (paramsMode/10)%10) {
				ic.Set(i+3, (paramsMode/100)%10, 1)
			} else {
				ic.Set(i+3, (paramsMode/100)%10, 0)
			}
			i += 4
		case 9:
			ic.Offset += ic.Get(i+1, paramsMode%10)
			i += 2
		}
	}
}
func (ic *IntcodeComputer) Get(addr, mode int) int {
	var nidx int
	switch mode {
	case ImediateMode:
		return ic.Intergers[addr]
	case PositionMode:
		nidx = ic.Intergers[addr]
	case RelativeMode:
		nidx = ic.Intergers[addr] + ic.Offset
	}
	if nidx >= 0 && nidx < len(ic.Intergers) {
		return ic.Intergers[nidx]
	}
	return ic.Extend[nidx]
}

func (ic *IntcodeComputer) Set(addr, mode, val int) {
	var nidx int
	switch mode {
	case PositionMode:
		nidx = ic.Intergers[addr]
	case RelativeMode:
		nidx = ic.Intergers[addr] + ic.Offset
	}
	if nidx >= 0 && nidx < len(ic.Intergers) {
		ic.Intergers[nidx] = val
	}
	ic.Extend[nidx] = val
}
