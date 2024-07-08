package problem_solving

import (
	"fmt"
	"testing"
)

func ExampleStaircase() {
	result := Staircase(6)
	fmt.Println(result)
	// Output:
	//      #
	//     ##
	//    ###
	//   ####
	//  #####
	// ######

}

func TestStaircase(t *testing.T) {
	type args struct {
		size int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				size: 6,
			},
			want: "     #\n    ##\n   ###\n  ####\n #####\n######\n",
		},
		{
			name: "Test Case 2",
			args: args{
				size: 2,
			},
			want: " #\n##\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Staircase(tt.args.size); got != tt.want {
				t.Errorf("Staircase() = %v, want %v", got, tt.want)
			}
		})
	}
}
