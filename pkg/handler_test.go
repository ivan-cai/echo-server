package pkg

import "testing"

func Test_echo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "echo succeed",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			echo()
		})
	}
}

func Test_test(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test failed",
			args: args{
				str: "test",
			},
			want: "test1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := test(tt.args.str); got != tt.want {
				t.Errorf("test() = %v, want %v", got, tt.want)
			}
		})
	}
}
