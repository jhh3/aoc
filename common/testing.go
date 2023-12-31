package common

import (
	"fmt"
	"testing"
)

type Test struct {
	Name  string
	Input string
	Part  int
	Want  string
}

func RunTests(ps ProblemSolver, t *testing.T, tests []Test) {
	for _, tt := range tests {
		name := fmt.Sprintf("%s-part%d", tt.Name, tt.Part)
		t.Run(name, func(t *testing.T) {
			fn := ps.SolvePart1
			if tt.Part == 2 {
				fn = ps.SolvePart2
			}

			if got := fn(tt.Input); got != tt.Want {
				t.Errorf("part%d() = %v, want %v", tt.Part, got, tt.Want)
			}
		})
	}
}
