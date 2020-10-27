package main

import "testing"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "Always failing to test CI"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// main()
			t.Error(tt.name)
		})
	}
}
