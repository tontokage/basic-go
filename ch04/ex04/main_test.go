package main

import "testing"

func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func TestMain(t *testing.T) {
	type args struct {
		s   []int
		cnt int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "0を指定",
			args: args{
				s:   []int{0, 1, 2, 3},
				cnt: 0,
			},
			want: []int{0, 1, 2, 3},
		},
		{
			name: "1を指定",
			args: args{
				s:   []int{0, 1, 2, 3},
				cnt: 1,
			},
			want: []int{1, 2, 3, 0},
		},
		{
			name: "末尾を指定",
			args: args{
				s:   []int{0, 1, 2, 3},
				cnt: 3,
			},
			want: []int{3, 0, 1, 2},
		},
		{
			name: "cap外の値を指定",
			args: args{
				s:   []int{0, 1, 2, 3},
				cnt: 4,
			},
			want: []int{0, 1, 2, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rotate(test.args.s, test.args.cnt)
			if !equal(test.args.s, test.want) {
				t.Errorf("rotate = %d, want = %d", test.args.s, test.want)
			}
		})
	}

}
