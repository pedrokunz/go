package problem_solving

import (
	"fmt"
	"testing"
)

func ExampleBirthdayCakeCandles() {
	candles := []int32{3, 2, 1, 3}
	result := BirthdayCakeCandles(candles)
	fmt.Println(result)
	// Output: 2
}

func TestBirthdayCakeCandles(t *testing.T) {
	type args struct {
		candles []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "Test Case 1",
			args: args{
				candles: []int32{3, 2, 1, 3},
			},
			want: 2,
		},
		{
			name: "Test Case 2",
			args: args{
				candles: []int32{1, 2, 3, 4, 5},
			},
			want: 1,
		},
		{
			name: "Test Case 3",
			args: args{
				candles: []int32{1, 1, 1, 1, 1},
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BirthdayCakeCandles(tt.args.candles); got != tt.want {
				t.Errorf("BirthdayCakeCandles() = %v, want %v", got, tt.want)
			}
		})
	}
}
