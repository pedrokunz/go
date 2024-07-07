package problem_solving

import (
	"fmt"
	"testing"
)

func ExamplePlusMinus() {
	array := []int32{-4, 3, -9, 0, 4}
	fmt.Println(PlusMinus(array))
	// Output:
	// 0.400000
	// 0.400000
	// 0.200000
}

func TestPlusMinus(t *testing.T) {
	type args struct {
		array []int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				array: []int32{-4, 3, -9, 0, 4},
			},
			want: "0.400000\n0.400000\n0.200000\n",
		},
		{
			name: "Test Case 2",
			args: args{
				array: []int32{1, 2, 3, -1, -2, -3, 0, 0},
			},
			want: "0.375000\n0.375000\n0.250000\n",
		},
		{
			name: "Test Case 3",
			args: args{
				array: []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
			want: "1.000000\n0.000000\n0.000000\n",
		},
		{
			name: "Test Case 4",
			args: args{
				array: []int32{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
			},
			want: "0.000000\n1.000000\n0.000000\n",
		},
		{
			name: "Test Case 5",
			args: args{
				array: []int32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			want: "0.000000\n0.000000\n1.000000\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := PlusMinus(tt.args.array); result != tt.want {
				t.Errorf("PlusMinus() = %v, want %v", result, tt.want)
			}
		})
	}
}
