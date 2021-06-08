package main

import "testing"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_add(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "both positive",
			args: args{11, 3},
			want: 19,
		},
		{
			name: "first positive",
			args: args{2, -3},
			want: -1,
		},
		{
			name: "second positive",
			args: args{-2, 3},
			want: 1,
		},
		{
			name: "both negative",
			args: args{-2, -3},
			want: -5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subtract(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Both positive",
			args: args{4, 5},
			want: 1,
		},
		{
			name: "B should be  positive",
			args: args{-4, 5},
			want: 9,
		},
		{
			name: "Both negative",
			args: args{-4, -5},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subtract(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}
