package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDepthIncreases(t *testing.T) {
	tests := []struct {
		name   string
		depths []int
		window int
		want   int
	}{
		{
			name:   "day1",
			depths: []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			window: 1,
			want:   7,
		},
		{
			name:   "day2",
			depths: []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			window: 3,
			want:   5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DepthIncreases(tt.depths, tt.window)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}

		})
	}
}
func TestParseFile(t *testing.T) {
	got, err := ParseInputFile("testdata/test.txt")
	if err != nil {
		t.Errorf("failed parsing file: %v", err)
	}

	want := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	if !cmp.Equal(got, want) {
		t.Errorf("got %v, want %v\ndiff: %v", got, want, cmp.Diff(got, want))
	}
}
