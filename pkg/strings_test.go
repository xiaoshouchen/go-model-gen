package pkg

import "testing"

func TestLineToLowCamel(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "小写驼峰",
			args: args{
				str: "hello_world",
			},
			want: "helloWorld",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LineToLowCamel(tt.args.str); got != tt.want {
				t.Errorf("LineToLowCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLineToUpCamel(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LineToUpCamel(tt.args.str); got != tt.want {
				t.Errorf("LineToUpCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}
