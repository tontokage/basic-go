package params

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestPack(t *testing.T) {
	type data struct {
		Labels      []string `http:"l"`
		MaxRequests int      `http:"max"`
		Exact       bool     `http:"x"`
	}
	tests := []struct {
		input data
		want  url.Values
	}{
		{
			input: data{[]string{"hello"}, 10, true},
			want: url.Values{
				"l":   []string{"hello"},
				"max": []string{"10"},
				"x":   []string{"true"},
			},
		},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("Pack(%#v)", test.input)
		got := &url.Values{}
		_ = Pack(test.input, got)
		if !reflect.DeepEqual(*got, test.want) {
			//if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
