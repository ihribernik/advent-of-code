package comparations

import "testing"

func TestContainsIn(t *testing.T) {
	type args struct {
		values          []string
		valuesToCompare []string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "a, b exists in valuesToCompare", args: args{values: []string{"a", "b"}, valuesToCompare: []string{"a", "b", "c"}}, want: true},
		{name: "a, b exists in valuesToCompare", args: args{values: []string{"a", "b"}, valuesToCompare: []string{"c", "d", "e"}}, want: false},
		{name: "a, b exists in valuesToCompare", args: args{values: []string{"a", "b"}, valuesToCompare: []string{"a", "b"}}, want: true},
		{name: "a, b exists in valuesToCompare", args: args{values: []string{"a", "b"}, valuesToCompare: []string{"b", "a"}}, want: true},
		{name: "a, b exists in valuesToCompare", args: args{values: []string{"a", "b"}, valuesToCompare: []string{"b", "a", "c"}}, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasAny(tt.args.values, tt.args.valuesToCompare); got != tt.want {
				t.Errorf("HasAny() = %v, want %v", got, tt.want)
			}
		})
	}
}
