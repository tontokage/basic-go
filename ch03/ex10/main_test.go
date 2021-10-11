package main

import "testing"

func TestComma(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "3桁未満",
			arg:  "123",
			want: "123",
		},
		{
			name: "4桁",
			arg:  "1234",
			want: "1,234",
		},
		{
			name: "5桁",
			arg:  "12345",
			want: "12,345",
		},
		{
			name: "6桁",
			arg:  "123456",
			want: "123,456",
		},
		{
			name: "7桁",
			arg:  "1234567",
			want: "1,234,567",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := comma(tt.arg); got != tt.want {
				t.Errorf("comma() = %v, want %v", got, tt.want)
			}

		})

	}

}
