package problem_solving

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func ExampleCompareTriplets() {
	tripletA := []int32{5, 6, 7}
	tripletB := []int32{3, 6, 10}
	result := CompareTriplets(tripletA, tripletB)
	fmt.Println(result)
	// Output: [1 1]
}

func TestCompareTriplets(t *testing.T) {
	type args struct {
		tripletA []int32
		tripletB []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "Test Case 1",
			args: args{
				tripletA: []int32{5, 6, 7},
				tripletB: []int32{3, 6, 10},
			},
			want: []int32{1, 1},
		},
		{
			name: "Test Case 2",
			args: args{
				tripletA: []int32{1, 2, 3},
				tripletB: []int32{4, 5, 6},
			},
			want: []int32{0, 3},
		},
		{
			name: "Test Case 3",
			args: args{
				tripletA: []int32{math.MaxInt32, math.MaxInt32, math.MaxInt32},
				tripletB: []int32{math.MinInt32, math.MinInt32, math.MinInt32},
			},
			want: []int32{3, 0},
		},
		{
			name: "Test Case 4",
			args: args{
				tripletA: []int32{1, 2, 3},
				tripletB: []int32{1, 2, 3},
			},
			want: []int32{0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareTriplets(tt.args.tripletA, tt.args.tripletB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompareTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
