package problem_solving

import (
	"fmt"
	"testing"
)

func ExampleTimeConversion() {
	result := TimeConversion("07:05:45PM")
	fmt.Println(result)
	// Output: 19:05:45
}

func TestTimeConversion(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				s: "07:05:45PM",
			},
			want: "19:05:45",
		},
		{
			name: "Test Case 2",
			args: args{
				s: "12:00:00AM",
			},
			want: "00:00:00",
		},
		{
			name: "Test Case 3",
			args: args{
				s: "12:00:00PM",
			},
			want: "12:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeConversion(tt.args.s); got != tt.want {
				t.Errorf("TimeConversion() = %v, want %v", got, tt.want)
			}
		})
	}
}
