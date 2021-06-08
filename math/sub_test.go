package math

import "testing"

func TestDiv(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "both positive",
			args: args{11, 3},
			want: 3.6666666666666665,
		},
		{
			name: "first positive",
			args: args{2, -3},
			want: -0.6666666666666666,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Div(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Div() = %v, want %v", got, tt.want)
			}
		})
	}
}
