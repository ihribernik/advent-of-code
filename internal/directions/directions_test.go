package directions

import (
	"reflect"
	"testing"
)

func TestNewDirection(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want Direction
	}{
		{name: "x: 1, y: 2", args: args{x: 1, y: 2}, want: NewDirection(1, 2)},
		{name: "x: -1, y: 2", args: args{x: -1, y: 2}, want: NewDirection(-1, 2)},
		{name: "x: 1, y: -2", args: args{x: 1, y: -2}, want: NewDirection(1, -2)},
		{name: "x: -1, y: -2", args: args{x: -1, y: -2}, want: NewDirection(-1, -2)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDirection(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirection_LessThan(t *testing.T) {
	type args struct {
		otherDir Direction
	}
	tests := []struct {
		name string
		dir  Direction
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dir.LessThan(tt.args.otherDir); got != tt.want {
				t.Errorf("Direction.LessThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirection_GreaterThan(t *testing.T) {
	type args struct {
		otherDir Direction
	}
	tests := []struct {
		name string
		dir  Direction
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dir.GreaterThan(tt.args.otherDir); got != tt.want {
				t.Errorf("Direction.GreaterThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirection_NewPositionWith(t *testing.T) {
	type args struct {
		otherDir Direction
	}
	tests := []struct {
		name string
		dir  Direction
		args args
		want Direction
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dir.NewPositionWith(tt.args.otherDir); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Direction.NewPositionWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirection_NewPosition(t *testing.T) {
	type args struct {
		otherDir Direction
		n        int
	}
	tests := []struct {
		name string
		dir  Direction
		args args
		want Direction
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dir.NewPosition(tt.args.otherDir, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Direction.NewPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirection_String(t *testing.T) {
	tests := []struct {
		name string
		dir  Direction
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dir.String(); got != tt.want {
				t.Errorf("Direction.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
