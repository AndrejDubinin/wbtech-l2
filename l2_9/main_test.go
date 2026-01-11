package main

import (
	"errors"
	"testing"
)

func TestTransform(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr error
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
			name:    "number",
			input:   "45",
			want:    "",
			wantErr: errInvalidInput,
		},
		{
			name:  "empty",
			input: "",
			want:  "",
		},
		{
			name:  "escape",
			input: `qwe\4\5`,
			want:  "qwe45",
		},
		{
			name:  "escape with repeat",
			input: `qwe\45`,
			want:  "qwe44444",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := transform(tt.input)
			if tt.wantErr == nil {
				if result != tt.want {
					t.Errorf("[%s] transform(%s), got: %q, want: %q", tt.name, tt.input, result, tt.want)
				}
				return
			}
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("[%s] transform(%s), got: %v, want: %v", tt.name, tt.input, err, tt.wantErr)
			}
		})
	}
}
