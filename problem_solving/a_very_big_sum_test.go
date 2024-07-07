package problem_solving

import (
	"fmt"
	"math"
	"testing"
)

func ExampleAVeryBigSum() {
	array := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	result := AVeryBigSum(array)
	fmt.Println(result)
	// Output: 5000000015
}

func TestAVeryBigSum(t *testing.T) {
	type args struct {
		array []int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Test Case 1",
			args: args{
				array: []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005},
			},
			want: 5000000015,
		},
		{
			name: "Test Case 2",
			args: args{
				array: []int64{1, 2, 3, 4, 5},
			},
			want: 15,
		},
		{
			name: "Test Case 3",
			args: args{
				array: []int64{math.MinInt64, math.MaxInt64},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AVeryBigSum(tt.args.array); got != tt.want {
				t.Errorf("AVeryBigSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
