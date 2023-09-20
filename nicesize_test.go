package nicesize

import (
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		input        uint64
		output       string
		spacedOutput string
	}{
		{500, "500B", "500 B"},
		{1024, "1KB", "1 KB"},
		{2 * 1024 * 1024, "2MB", "2 MB"},
		{5e9, "4.66GB", "4.66 GB"},
		{5e12, "4.55TB", "4.55 TB"},
		{5e15, "4.44PB", "4.44 PB"},
		{5e18, "4.34EB", "4.34 EB"},
		{8000_000_000_000_000_000, "6.94EB", "6.94 EB"},
	}

	for _, tt := range testCases {
		t.Run(tt.output, func(t *testing.T) {
			got := Format(tt.input)
			got2 := FormatWithSpacer(tt.input, " ")

			if got != tt.output || got2 != tt.spacedOutput {
				t.Errorf("Expected %s but got %s", tt.output, got)
			}

			if got2 != tt.spacedOutput {
				t.Errorf("Expected %s but got %s", tt.output, got2)
			}
		})
	}
}
