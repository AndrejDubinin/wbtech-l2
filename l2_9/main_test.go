package main

import (
	"testing"
)

func TestTransform(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "normal",
			input: "a4bc2d5e",
			want:  "aaaabccddddde",
		},
		{
			name:  "equal",
			input: "abcd",
			want:  "abcd",
		},
		{
			name:  "number",
			input: "45",
			want:  "",
		},
		{
			name:  "empty",
			input: "",
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := transform(tt.input)
			if result != tt.want {
				t.Errorf("[%q] transform(%q) = %q, want %q", tt.name, tt.input, result, tt.want)
			}
		})
	}
}
