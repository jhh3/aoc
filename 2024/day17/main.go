package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/jhh3/aoc/common"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	common.RunFromSolver(&solver{}, input)
}

//--------------------------------------------------------------------
// Solution
//--------------------------------------------------------------------

type solver struct{}

func (s *solver) SolvePart1(input string) string {
	problemInput := parseInput(input)
	problemInput.PrettyPrint()
	outputs := problemInput.Run()
	commaSeperateStr := ""
	for i, output := range outputs {
		if i > 0 {
			commaSeperateStr += ","
		}
		commaSeperateStr += strconv.Itoa(output)
	}

	return commaSeperateStr
}

func (s *solver) SolvePart2(input string) string {
	pi := parseInput(input)
	result := pi.RecoverCorruptedRegisterASmart()
	return strconv.Itoa(result)
}

type Instruction struct {
	value int
}

type ProblemInput struct {
	RegisterA int
	RegisterB int
	RegisterC int

	InstructionPointer int

	Program []Instruction
}

func (pi *ProblemInput) getComboOperandValue(operand int) int {
	if operand <= 3 {
		return operand
	}
	switch operand {
	case 4:
		return pi.RegisterA
	case 5:
		return pi.RegisterB
	case 6:
		return pi.RegisterC
	default:
		panic(fmt.Sprintf("invalid combo operand: %d", operand))
	}
}

// func (pi *ProblemInput) RecoverCorruptedRegisterA() int {
// 	targetProgram := make([]int, len(pi.Program))
// 	for i, inst := range pi.Program {
// 		targetProgram[i] = inst.value
// 	}
//
// 	// Number of worker goroutines
// 	numWorkers := 1
// 	results := make(chan int)
// 	done := make(chan bool)
//
// 	// Launch workers to search different ranges
// 	for worker := 0; worker < numWorkers; worker++ {
// 		go func(startAt int) {
// 			// Each worker starts at their own offset and increments by numWorkers
// 			for regA := startAt; regA <= 100000000000; regA += numWorkers {
// 				// Print progress every 1000000
// 				if regA%1000000 == 0 {
// 					fmt.Printf("Worker %d: %d\n", worker+1, regA)
// 				}
// 				testPi := &ProblemInput{
// 					RegisterA:          regA,
// 					RegisterB:          pi.RegisterB,
// 					RegisterC:          pi.RegisterC,
// 					InstructionPointer: 0,
// 					Program:            make([]Instruction, len(pi.Program)),
// 				}
// 				copy(testPi.Program, pi.Program)
//
// 				output := testPi.Run()
// 				outputValues := strings.Split(output, ",")
//
// 				if len(outputValues) == len(targetProgram) {
// 					match := true
// 					for i := 0; i < len(outputValues); i++ {
// 						val := common.MustAtoi(outputValues[i])
// 						if val != targetProgram[i] {
// 							match = false
// 							break
// 						}
// 					}
// 					if match {
// 						results <- regA
// 						return
// 					}
// 				}
//
// 				// Quick check if we're done
// 				select {
// 				case <-done:
// 					return
// 				default:
// 				}
// 			}
// 		}(worker + 1) // Start at 1, 2, 3, etc.
// 	}
//
// 	// Wait for first result
// 	result := <-results
// 	close(done) // Signal all workers to stop
//
// 	return result
// }

func (pi *ProblemInput) Run() []int {
	outputs := make([]int, 0)

	for pi.InstructionPointer < len(pi.Program) {
		opcode := pi.Program[pi.InstructionPointer]
		operand := pi.Program[pi.InstructionPointer+1].value

		switch opcode.value {
		case 0b000: // adv
			denominator := 1 << pi.getComboOperandValue(operand)
			pi.RegisterA = pi.RegisterA / denominator

		case 0b001: // bxl
			pi.RegisterB = pi.RegisterB ^ operand

		case 0b010: // bst
			pi.RegisterB = pi.getComboOperandValue(operand) % 8

		case 0b011: // jnz
			if pi.RegisterA != 0 {
				pi.InstructionPointer = operand
				continue // Skip the normal increment
			}

		case 0b100: // bxc
			pi.RegisterB = pi.RegisterB ^ pi.RegisterC

		case 0b101: // out
			value := pi.getComboOperandValue(operand) % 8
			outputs = append(outputs, value)

		case 0b110: // bdv
			denominator := 1 << pi.getComboOperandValue(operand)
			pi.RegisterB = pi.RegisterA / denominator

		case 0b111: // cdv
			denominator := 1 << pi.getComboOperandValue(operand)
			pi.RegisterC = pi.RegisterA / denominator
		}

		pi.InstructionPointer += 2
	}

	return outputs
}

func (pi *ProblemInput) PrettyPrint() {
	fmt.Printf("Register A: %d\n", pi.RegisterA)
	fmt.Printf("Register B: %d\n", pi.RegisterB)
	fmt.Printf("Register C: %d\n", pi.RegisterC)
	fmt.Printf("Program: %v\n", pi.Program)
}

func parseInput(input string) *ProblemInput {
	lines := common.ReadAsLines(input)
	pi := &ProblemInput{
		Program: make([]Instruction, 0),
	}
	for _, line := range lines {
		println(line)
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		if strings.HasPrefix(line, "Register A") {
			val := strings.TrimPrefix(line, "Register A: ")
			pi.RegisterA = common.MustAtoi(val)
		} else if strings.HasPrefix(line, "Register B") {
			val := strings.TrimPrefix(line, "Register B: ")
			pi.RegisterB = common.MustAtoi(val)
		} else if strings.HasPrefix(line, "Register C") {
			val := strings.TrimPrefix(line, "Register C: ")
			pi.RegisterC = common.MustAtoi(val)
		} else if strings.HasPrefix(line, "Program: ") {
			valuesStr := strings.TrimPrefix(line, "Program: ")
			values := strings.Split(valuesStr, ",")
			for _, valStr := range values {
				val := common.MustAtoi(valStr)
				pi.Program = append(pi.Program, Instruction{value: val})
			}
		}
	}
	return pi
}

func (pi *ProblemInput) cloneWithA(aVal int) *ProblemInput {
	newPi := &ProblemInput{
		RegisterA:          aVal,
		RegisterB:          pi.RegisterB,
		RegisterC:          pi.RegisterC,
		InstructionPointer: 0,
		Program:            make([]Instruction, len(pi.Program)),
	}
	copy(newPi.Program, pi.Program)
	return newPi
}

func (pi *ProblemInput) ComputeOutputFormulas() []string {
	// We'll run a symbolic simulation of the program.
	// Initialize registers as symbolic variables
	A := "A"
	B := "B"
	C := "C"

	outputFormulas := []string{}

	// Helper to get symbolic value of combo operand
	getSymbolicOperand := func(operand int) string {
		if operand <= 3 {
			// Just a number as a string
			return fmt.Sprintf("%d", operand)
		}
		switch operand {
		case 4:
			return A
		case 5:
			return B
		case 6:
			return C
		default:
			panic(fmt.Sprintf("invalid combo operand: %d", operand))
		}
	}

	// Make a copy of the program counter
	ip := 0
	for ip < len(pi.Program) {
		opcode := pi.Program[ip].value
		operand := pi.Program[ip+1].value

		switch opcode {
		case 0b000: // adv: A = A / (2^(comboOperand))
			x := getSymbolicOperand(operand)
			// x should be numeric (0-3 or known register?), ideally we handle only numeric here
			// If x is register-based (like A,B,C), that complicates. Assume operand ≤ 3 or a known reg that we must handle.
			// For simplicity, assume operand is always ≤3 for division shift. If not, you'd need additional logic.
			A = fmt.Sprintf("(%s/(2^(%s)))", A, x)

		case 0b001: // bxl: B = B ^ operand
			// operand here is a raw value used for XOR (0–...?)
			B = fmt.Sprintf("(%s ^ %d)", B, operand)

		case 0b010: // bst: B = (comboOperand % 8)
			val := getSymbolicOperand(operand)
			B = fmt.Sprintf("((%s) %% 8)", val)

		case 0b011: // jnz: if A != 0 jump
			// Symbolically, we don't know if A != 0. Handling branches symbolically is complex.
			// For a real solution, you'd represent both paths or have constraints.
			// Here, we just do nothing and assume no jump for simplicity.
			// A full solution would require more complex logic.
			ip += 2
			continue

		case 0b100: // bxc: B = B ^ C
			B = fmt.Sprintf("(%s ^ %s)", B, C)

		case 0b101: // out: output = (comboOperand % 8)
			val := getSymbolicOperand(operand)
			outExpr := fmt.Sprintf("((%s) %% 8)", val)
			outputFormulas = append(outputFormulas, outExpr)

		case 0b110: // bdv: B = A / (2^(comboOperand))
			x := getSymbolicOperand(operand)
			B = fmt.Sprintf("(%s/(2^(%s)))", A, x)

		case 0b111: // cdv: C = A / (2^(comboOperand))
			x := getSymbolicOperand(operand)
			C = fmt.Sprintf("(%s/(2^(%s)))", A, x)
		}

		ip += 2
	}

	return outputFormulas
}

// runProgram is a helper that runs the program given A, B, C, and returns the output as a slice of ints.
func runProgram(A, B, C int, program []int) []int {
	pi := &ProblemInput{
		RegisterA: A,
		RegisterB: B,
		RegisterC: C,
		Program:   make([]Instruction, len(program)),
	}
	for i, val := range program {
		pi.Program[i] = Instruction{value: val}
	}
	return pi.Run()
}

func (pi *ProblemInput) RecoverCorruptedRegisterASmart() int {
	progValues := make([]int, len(pi.Program))
	for i, inst := range pi.Program {
		progValues[i] = inst.value
	}

	// Try to find a RegisterA that produces the program as output (quine)
	val, found := getBestQuineInput(progValues, len(progValues)-1, 0)
	if !found {
		return -1
	}
	return val
}

func getBestQuineInput(program []int, cursor int, sofar int) (int, bool) {
	for candidate := 0; candidate < 8; candidate++ {
		testA := sofar*8 + candidate
		output := runProgram(testA, 0, 0, program)
		if slicesEqual(output, program[cursor:]) {
			if cursor == 0 {
				return testA, true
			}
			ret, found := getBestQuineInput(program, cursor-1, sofar*8+candidate)
			if found {
				return ret, true
			}
		}
	}
	return 0, false
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
