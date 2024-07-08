package problem_solving

import (
	"fmt"
	"testing"
)

func ExampleAppleAndOrange() {
	startPoint := int32(7)
	endPoint := int32(11)
	appleTreeLocation := int32(5)
	orangeTreeLocation := int32(15)
	apples := []int32{-2, 2, 1}
	oranges := []int32{5, -6}
	fmt.Println(AppleAndOrange(startPoint, endPoint, appleTreeLocation, orangeTreeLocation, apples, oranges))
	// Output: 1 1
}

func TestAppleAndOrange(t *testing.T) {
	type args struct {
		startPoint         int32
		endPoint           int32
		appleTreeLocation  int32
		orangeTreeLocation int32
		apples             []int32
		oranges            []int32
	}
	tests := []struct {
		name        string
		args        args
		wantApples  int
		wantOranges int
	}{
		{
			name: "Test Case 1",
			args: args{
				startPoint:         7,
				endPoint:           11,
				appleTreeLocation:  5,
				orangeTreeLocation: 15,
				apples:             []int32{-2, 2, 1},
				oranges:            []int32{5, -6},
			},
			wantApples:  1,
			wantOranges: 1,
		},
		{
			name: "Test Case 2",
			args: args{
				startPoint:         2,
				endPoint:           3,
				appleTreeLocation:  1,
				orangeTreeLocation: 5,
				apples:             []int32{2},
				oranges:            []int32{-2},
			},
			wantApples:  1,
			wantOranges: 1,
		},
		{
			name: "Test Case 3",
			args: args{
				startPoint:         5,
				endPoint:           10,
				appleTreeLocation:  5,
				orangeTreeLocation: 10,
				apples:             []int32{0, 1, 2, 3, 4},
				oranges:            []int32{0, -1, -2, -3, -4},
			},
			wantApples:  5,
			wantOranges: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotApples, gotOranges := AppleAndOrange(tt.args.startPoint, tt.args.endPoint, tt.args.appleTreeLocation, tt.args.orangeTreeLocation, tt.args.apples, tt.args.oranges)
			if gotApples != tt.wantApples {
				t.Errorf("AppleAndOrange() got = %v, want %v", gotApples, tt.wantApples)
			}

			if gotOranges != tt.wantOranges {
				t.Errorf("AppleAndOrange() got1 = %v, want %v", gotOranges, tt.wantOranges)
			}
		})
	}
}
