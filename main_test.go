package main

import "testing"

func TestConvertCode(t *testing.T) {
	type args struct {
		numberSet string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{"333"},
			want: "f",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertCode(tt.args.numberSet); got != tt.want {
				t.Errorf("ConvertCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
