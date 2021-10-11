package main

import "testing"

func TestMain(t *testing.T) {
	type args struct {
		a string
		b string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "アナグラム",
			args: args{a: "アナグラム", b: "グアムナラ"},
			want: true,
		},
		{
			name: "アナグラム 比較対象の入れ替え",
			args: args{a: "グアムナラ", b: "アナグラム"},
			want: true,
		},
		{
			name: "余分な文字1",
			args: args{a: "ア", b: "グアムナラ"},
			want: false,
		},
		{
			name: "余分な文字2",
			args: args{a: "グアムナラ", b: "ア"},
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := isAnagram(test.args.a, test.args.b); got != test.want {
				t.Errorf("isAnagram = %v, want = %v", got, test.want)
			}
		})
	}

}
