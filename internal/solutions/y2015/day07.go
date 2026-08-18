package y2015

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type Day07 struct{}

type operationKind int

const (
	operationAnd operationKind = iota
	operationOr
	operationLShift
	operationRShift
	operationNot
	operationAssign
)

type gate struct {
	kind operationKind
	x    string
	y    string
}

func New(kind operationKind, x,y string) gate{
	return gate {
		kind: kind,
		x: x,
		y: y,
	}
}

var (
	andRe    = regexp.MustCompile("^([^ ]+) AND ([^ ]+) -> ([^ ]+)$")
	assignRe = regexp.MustCompile("^([^ ]+) -> ([^ ]+)$")
	lshiftRe = regexp.MustCompile("^([^ ]+) LSHIFT ([^ ]+) -> ([^ ]+)$")
	notRe    = regexp.MustCompile("^(NOT) ([^ ]+) -> ([^ ]+)$")
	orRe     = regexp.MustCompile("^([^ ]+) OR ([^ ]+) -> ([^ ]+)$")
	rshiftRe = regexp.MustCompile("^([^ ]+) RSHIFT ([^ ]+) -> ([^ ]+)$")
)

func generateCircuit(input []string) (map[string]gate, error) {
	wires := make(map[string]gate)

	for i, line := range input {
		switch {
		case andRe.MatchString(line):
			m := andRe.FindStringSubmatch(line)
			wires[m[3]] = gate{kind: operationAnd, x: m[1], y: m[2]}
		case assignRe.MatchString(line):
			m := assignRe.FindStringSubmatch(line)
			wires[m[2]] = gate{kind: operationAssign, x: m[1]}
		case lshiftRe.MatchString(line):
			m := lshiftRe.FindStringSubmatch(line)
			wires[m[3]] = gate{kind: operationLShift, x: m[1], y: m[2]}
		case notRe.MatchString(line):
			m := notRe.FindStringSubmatch(line)
			wires[m[3]] = gate{kind: operationNot, x: m[2]}
		case orRe.MatchString(line):
			m := orRe.FindStringSubmatch(line)
			wires[m[3]] = gate{kind: operationOr, x: m[1], y: m[2]}
		case rshiftRe.MatchString(line):
			m := rshiftRe.FindStringSubmatch(line)
			wires[m[3]] = gate{kind: operationRShift, x: m[1], y: m[2]}
		default:
			return nil, newInputError(7, i+1, line, errors.New("operation not matched"))
		}
	}
	return wires, nil
}

func calculateWireValue(wire string, wires map[string]gate, memoization map[string]uint16) (uint16, error) {
	if v, ok := memoization[wire]; ok {
		return v, nil
	}

	if n, err := strconv.ParseUint(wire, 10, 16); err == nil {
		return uint16(n), nil
	}

	g, ok := wires[wire]
	if !ok {
		return 0, newInputError(7, 0, wire, errors.New("wire not found"))
	}

	var result uint16

	switch g.kind {
	case operationAnd:
		a, err := calculateWireValue(g.x, wires, memoization)
		if err != nil {
			return 0, err
		}

		b, err := calculateWireValue(g.y, wires, memoization)
		if err != nil {
			return 0, err
		}
		result = a & b
	case operationOr:
		a, err := calculateWireValue(g.x, wires, memoization)
		if err != nil {
			return 0, err
		}

		b, err := calculateWireValue(g.y, wires, memoization)
		if err != nil {
			return 0, err
		}
		result = a | b
	case operationLShift:
		a, err := calculateWireValue(g.x, wires, memoization)
		if err != nil {
			return 0, err
		}

		b, err := calculateWireValue(g.y, wires, memoization)
		if err != nil {
			return 0, err
		}
		result = a << b
	case operationRShift:
		a, err := calculateWireValue(g.x, wires, memoization)
		if err != nil {
			return 0, err
		}

		b, err := calculateWireValue(g.y, wires, memoization)
		if err != nil {
			return 0, err
		}
		result = a >> b
	case operationNot:
		v, err := calculateWireValue(g.x, wires, memoization)
		if err != nil {
			return 0, err
		}
		result = ^v
	case operationAssign:
		v, err := calculateWireValue(g.x, wires, memoization)
		if err != nil {
			return 0, err
		}
		result = v
	default:
		return 0, fmt.Errorf("unknown operation kind %d", g.kind)
	}
	result &= 0xFFFF

	memoization[wire] = result
	return result, nil
}

// SolvePart1 implements solutions.Solver.
func (d Day07) SolvePart1(input []string) (int, error) {
	wires, err := generateCircuit(input)
	if err != nil {
		return 0, err
	}

	memoization := make(map[string]uint16)

	value, err := calculateWireValue("a", wires, memoization)
	if err != nil {
		return 0, err
	}

	return int(value), nil
}

// SolvePart2 implements solutions.Solver.
func (d Day07) SolvePart2(input []string) (int, error) {
	wires, err := generateCircuit(input)
	if err != nil {
		return 0, err
	}

	memoization := make(map[string]uint16)
	valueA, err := calculateWireValue("a", wires, memoization)
	if err != nil {
		return 0, err
	}

	wires, err = generateCircuit(input)
	if err != nil {
		return 0, err
	}

	wires["b"] = gate{kind: operationAssign, x: fmt.Sprintf("%d", valueA)}
	memoization = make(map[string]uint16)
	value, err := calculateWireValue("a", wires, memoization)
	if err != nil {
		return 0, err
	}

	return int(value), nil
}
