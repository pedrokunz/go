package problem_solving

import (
	"fmt"
	"testing"
)

func ExampleKangaroo() {
	kangaroo1StartPosition := int32(0)
	kangaroo1Speed := int32(3)
	kangaroo2StartPosition := int32(4)
	kangaroo2Speed := int32(2)
	result := Kangaroo(kangaroo1StartPosition, kangaroo1Speed, kangaroo2StartPosition, kangaroo2Speed)

	fmt.Println(result)
	// Output: YES
}

func TestKangaroo(t *testing.T) {
	type args struct {
		kangaroo1StartPosition int32
		kangaroo1Speed         int32
		kangaroo2StartPosition int32
		kangaroo2Speed         int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				kangaroo1StartPosition: 0,
				kangaroo1Speed:         3,
				kangaroo2StartPosition: 4,
				kangaroo2Speed:         2,
			},
			want: "YES",
		},
		{
			name: "Test Case 2",
			args: args{
				kangaroo1StartPosition: 0,
				kangaroo1Speed:         2,
				kangaroo2StartPosition: 5,
				kangaroo2Speed:         3,
			},
			want: "NO",
		},
		{
			name: "Test Case 3",
			args: args{
				kangaroo1StartPosition: 0,
				kangaroo1Speed:         2,
				kangaroo2StartPosition: 0,
				kangaroo2Speed:         2,
			},
			want: "YES",
		},
		{
			name: "Test Case 4",
			args: args{
				kangaroo1StartPosition: 0,
				kangaroo1Speed:         4,
				kangaroo2StartPosition: 2,
				kangaroo2Speed:         2,
			},
			want: "YES",
		},
		{
			name: "Test Case 5",
			args: args{
				kangaroo1StartPosition: -1,
				kangaroo1Speed:         2,
				kangaroo2StartPosition: -4,
				kangaroo2Speed:         2,
			},
			want: "NO",
		},
		{
			name: "Test Case 6",
			args: args{
				kangaroo1StartPosition: 0,
				kangaroo1Speed:         0,
				kangaroo2StartPosition: 5,
				kangaroo2Speed:         0,
			},
			want: "NO",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Kangaroo(tt.args.kangaroo1StartPosition, tt.args.kangaroo1Speed, tt.args.kangaroo2StartPosition, tt.args.kangaroo2Speed); got != tt.want {
				t.Errorf("Kangaroo() = %v, want %v", got, tt.want)
			}
		})
	}
}
