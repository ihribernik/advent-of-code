package y2015_test

import (
	"testing"

	"github.com/ihribernik/aoc-cli/internal/solutions/y2015"
)

func TestDay10_SolvePart1(t *testing.T) {
	tests := []struct {
		name    string // description of this test case
		input   []string
		want    string
		wantErr bool
	}{
		{name: "1 becomes 11 (1 copy of digit 1)", input: []string{"1"}, want: "11", wantErr: false},
		{name: "11 becomes 21 (2 copies of digit 1).", input: []string{"11"}, want: "21", wantErr: false},
		{name: "21 becomes 1211 (one 2 followed by one 1).", input: []string{"21"}, want: "1211", wantErr: false},
		{name: "1211 becomes 111221 (one 1, one 2, and two 1s)", input: []string{"1211"}, want: "111221", wantErr: false},
		{name: "111221 becomes 312211 (three 1s, two 2s, and one 1).", input: []string{"111221"}, want: "312211", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var d y2015.Day10
			got, gotErr := d.SolvePart1(tt.input)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("SolvePart1() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("SolvePart1() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("SolvePart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay10_SolvePart2(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		input   []string
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var d y2015.Day10
			got, gotErr := d.SolvePart2(tt.input)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("SolvePart2() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("SolvePart2() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("SolvePart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
