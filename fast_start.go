package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func processVideoForFastStart(filePath string) (string, error) {
	processedPath := fmt.Sprintf("%s.processing", filePath)

	// Run ffmpeg to move moov atom to start
	// -movflags +faststart: Moves the moov atom to the beginning of the file
	cmd := exec.Command("ffmpeg", "-i", filePath, "-c", "copy", "-movflags", "+faststart", "-f", "mp4", processedPath)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		os.Remove(processedPath)
		return "", fmt.Errorf("failed to process video with ffmpeg: %s: %w", stderr.String(), err)
	}

	return processedPath, nil
}
