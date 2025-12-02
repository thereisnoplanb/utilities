package utilities

import (
	"reflect"
	"testing"
)

func TestTernary(t *testing.T) {
	type args struct {
		condition   bool
		consequent  int
		alternative int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Condition True",
			args: args{
				condition:   true,
				consequent:  100,
				alternative: -100,
			},
			want: 100,
		},
		{
			name: "Condition False",
			args: args{
				condition:   false,
				consequent:  100,
				alternative: -100,
			},
			want: -100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ternary(tt.args.condition, tt.args.consequent, tt.args.alternative); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ternary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTernaryFunc(t *testing.T) {
	type args struct {
		condition   bool
		consequent  func() int
		alternative func() int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Condition True",
			args: args{
				condition:   true,
				consequent:  func() int { return 100 },
				alternative: func() int { return -100 },
			},
			want: 100,
		},
		{
			name: "Condition False",
			args: args{
				condition:   false,
				consequent:  func() int { return 100 },
				alternative: func() int { return -100 },
			},
			want: -100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TernaryFunc(tt.args.condition, tt.args.consequent, tt.args.alternative); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TernaryFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
