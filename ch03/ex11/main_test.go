package main

import "testing"

func TestComma(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "整数3桁未満",
			arg:  "123",
			want: "123",
		},
		{
			name: "整数4桁",
			arg:  "1234",
			want: "1,234",
		},
		{
			name: "整数5桁",
			arg:  "12345",
			want: "12,345",
		},
		{
			name: "整数6桁",
			arg:  "123456",
			want: "123,456",
		},
		{
			name: "整数7桁",
			arg:  "1234567",
			want: "1,234,567",
		},
		{
			name: "浮動小数点3桁未満",
			arg:  "123.1",
			want: "123.1",
		},
		{
			name: "浮動小数点4桁",
			arg:  "1234.1",
			want: "1,234.1",
		},
		{
			name: "浮動小数点6桁",
			arg:  "123456.1",
			want: "123,456.1",
		},
		{
			name: "浮動小数点7桁",
			arg:  "1234567.1",
			want: "1,234,567.1",
		},
		{
			name: "負の浮動小数点3桁未満",
			arg:  "-123.1",
			want: "-123.1",
		},
		{
			name: "浮動小数点4桁",
			arg:  "-1234.1",
			want: "-1,234.1",
		},
		{
			name: "浮動小数点6桁",
			arg:  "-123456.1",
			want: "-123,456.1",
		},
		{
			name: "浮動小数点7桁",
			arg:  "-1234567.1",
			want: "-1,234,567.1",
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
