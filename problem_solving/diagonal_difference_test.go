package problem_solving

import (
	"fmt"
	"testing"
)

func ExampleDiagonalDifference() {
	matrix := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}
	result := DiagonalDifference(matrix)
	fmt.Println(result)
	// Output: -15
}

func TestDiagonalDifference(t *testing.T) {
	type args struct {
		matrix [][]int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "Test Case 1",
			args: args{
				matrix: [][]int32{
					{11, 2, 4},
					{4, 5, 6},
					{10, 8, -12},
				},
			},
			want: -15,
		},
		{
			name: "Test Case 2",
			args: args{
				matrix: [][]int32{
					{1, 2},
					{3, 4},
				},
			},
			want: 0,
		},
		{
			name: "Test Case 3",
			args: args{
				matrix: [][]int32{
					{1, 2, 3, 4},
					{5, 6, 7, 8},
					{9, 10, 11, 12},
					{13, 14, 15, 16},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DiagonalDifference(tt.args.matrix); got != tt.want {
				t.Errorf("DiagonalDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
