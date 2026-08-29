package main

import (
	"testing"
)

func TestGetVideoAspectRatio(t *testing.T) {
	tests := []struct {
		name     string
		file     string
		wantType string
	}{
		{
			name:     "horizontal video",
			file:     "samples/boots-video-horizontal.mp4",
			wantType: "landscape",
		},
		{
			name:     "vertical video",
			file:     "samples/boots-video-vertical.mp4",
			wantType: "portrait",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aspectRatio, err := getVideoAspectRatio(tt.file)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			aspectRatioType := getAspectRatioType(aspectRatio)
			if aspectRatioType != tt.wantType {
				t.Errorf("got %s, want %s (aspect ratio was %s)", aspectRatioType, tt.wantType, aspectRatio)
			}
		})
	}
}
