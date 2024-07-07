package problem_solving

import (
	"fmt"
	"math"
	"testing"
)

func ExampleSimpleArraySum() {
	array := []int32{1, 2, 3, 4, 5}
	result := SimpleArraySum(array)
	fmt.Println(result)
	// Output: 15
}

func TestSimpleArraySum(t *testing.T) {
	type args struct {
		ar []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "Test Case 1",
			args: args{
				ar: []int32{1, 10, 100, 1000, 10000},
			},
			want: 11111,
		},
		{
			name: "Test Case 2",
			args: args{
				ar: []int32{1, 2, 3, 4, 5},
			},
			want: 15,
		},
		{
			name: "Test Case 3",
			args: args{
				ar: []int32{-123, 0, 123},
			},
			want: 0,
		},
		{
			name: "Test Case 4",
			args: args{
				ar: []int32{math.MinInt32, math.MaxInt32},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SimpleArraySum(tt.args.ar); got != tt.want {
				t.Errorf("SimpleArraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
