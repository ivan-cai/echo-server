package pkg

import "testing"

func Test_echo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "succeed",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			echo()
		})
	}
}
