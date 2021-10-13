package main

import "testing"

func TestReverseArr(t *testing.T) {
	tests := []struct {
		name string
		arg  [4]int
		want [4]int
	}{
		{
			name: "0-3",
			arg:  [4]int{0, 1, 2, 3},
			want: [4]int{3, 2, 1, 0},
		},
		{
			name: "3-0",
			arg:  [4]int{3, 2, 1, 0},
			want: [4]int{0, 1, 2, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reverseArr(&test.arg)
			if test.arg != test.want {
				t.Errorf("reverseArr = %v, want = %v", test.arg, test.want)
			}
		})

	}

}
