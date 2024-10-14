package problem_solving

import (
	"testing"
)

func TestRotateArray(t *testing.T) {
	type args struct {
		array     []int
		rotations int
		direction string
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test case 1",
			args: args{
				array:     []int{1, 2, 3, 4, 5},
				rotations: 2,
				direction: "left",
			},
			want: []int{3, 4, 5, 1, 2},
		},
		{
			name: "Test case 2",
			args: args{
				array:     []int{1, 2, 3, 4, 5},
				rotations: 4,
				direction: "left",
			},
			want: []int{5, 1, 2, 3, 4},
		},
		{
			name: "Test case 3",
			args: args{
				array:     []int{1, 2, 3, 4, 5},
				rotations: 5,
				direction: "left",
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Test case 4",
			args: args{
				array:     []int{1, 2, 3, 4, 5},
				rotations: 0,
				direction: "left",
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Test case 5",
			args: args{
				array:     []int{1, 2, 3, 4, 5},
				rotations: -1,
				direction: "left",
			},
			want: []int{2, 3, 4, 5, 1},
		},
		{
			name: "Test case 6",
			args: args{
				array:     []int{1, 2, 3, 4, 5},
				rotations: 6,
				direction: "left",
			},
			want: []int{2, 3, 4, 5, 1},
		},
		{
			name: "Test case 7",
			args: args{
				array:     []int{1, 2, 3, 4, 5},
				rotations: 2,
				direction: "right",
			},
			want: []int{4, 5, 1, 2, 3},
		},
		{
			name: "Test case 8",
			args: args{
				array:     []int{1, 2, 3, 4, 5},
				rotations: 6,
				direction: "right",
			},
			want: []int{5, 1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RotateArray(tt.args.array, tt.args.rotations, tt.args.direction)
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Expected %v, got %v", tt.want, got)
				}
			}
		})
	}
}
