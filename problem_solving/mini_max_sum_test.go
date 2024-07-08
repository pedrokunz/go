package problem_solving

import "testing"

func TestMiniMaxSum(t *testing.T) {
	type args struct {
		array []int32
	}
	tests := []struct {
		name        string
		args        args
		wantMinimum int32
		wantMaximum int32
	}{
		{
			name: "Test Case 1",
			args: args{
				array: []int32{1, 2, 3, 4, 5},
			},
			wantMinimum: 10,
			wantMaximum: 14,
		},
		{
			name: "Test Case 2",
			args: args{
				array: []int32{5, 4, 3, 2, 1},
			},
			wantMinimum: 10,
			wantMaximum: 14,
		},
		{
			name: "Test Case 3",
			args: args{
				array: []int32{1, 1, 1, 1, 1},
			},
			wantMinimum: 4,
			wantMaximum: 4,
		},
		{
			name: "Test Case 4",
			args: args{
				array: []int32{1, 2, 3, -4, -5, -6, 7, 8, 9, 10},
			},
			wantMinimum: 15,
			wantMaximum: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMinimum, gotMaximum := MiniMaxSum(tt.args.array)
			if gotMinimum != tt.wantMinimum {
				t.Errorf("MiniMaxSum() got = %v, want %v", gotMinimum, tt.wantMinimum)
			}
			if gotMaximum != tt.wantMaximum {
				t.Errorf("MiniMaxSum() got1 = %v, want %v", gotMaximum, tt.wantMaximum)
			}
		})
	}
}
