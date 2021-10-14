package main

import "testing"

func equal(x, y []string) bool {
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
	tests := []struct {
		name string
		args []string
		want []string
	}{
		{
			name: "隣り合う重複あり",
			args: []string{"a", "a", "b", "c", "d"},
			want: []string{"a", "b", "c", "d"},
		},
		{
			name: "隣り合う重複なし",
			args: []string{"a", "b", "c", "d", "e"},
			want: []string{"a", "b", "c", "d", "e"},
		},
		{
			name: "隣り合う3つの重複",
			args: []string{"a", "a", "a", "b", "c"},
			want: []string{"a", "b", "c"},
		},
		{
			name: "先頭・末尾の重複",
			args: []string{"a", "b", "c", "d", "a"},
			want: []string{"a", "b", "c", "d", "a"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := uniq(test.args)
			if !equal(got, test.want) {
				t.Errorf("uniq = %s, want = %s", got, test.want)
			}
		})
	}

}
