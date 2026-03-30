package y2015

import "testing"

func TestDay09_SolvePart1(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name    string
		d       Day09
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "sample input",
			d:    Day09{},
			args: args{
				input: []string{
					"London to Dublin = 464",
					"London to Belfast = 518",
					"Dublin to Belfast = 141",
				},
			},
			want:    605,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.SolvePart1(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Day09.SolvePart1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Day09.SolvePart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay09_SolvePart2(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name    string
		d       Day09
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "sample input",
			d:    Day09{},
			args: args{
				input: []string{
					"London to Dublin = 464",
					"London to Belfast = 518",
					"Dublin to Belfast = 141",
				},
			},
			want:    982,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.SolvePart2(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Day09.SolvePart2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Day09.SolvePart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
