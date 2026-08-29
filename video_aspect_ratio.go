package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"os/exec"
)

type ffprobeStream struct {
	CodecType string `json:"codec_type"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
}

type ffprobeOutput struct {
	Streams []ffprobeStream `json:"streams"`
}

func getVideoAspectRatio(filePath string) (string, error) {
	ffprobePath := "ffprobe"
	if os.Getenv("FFPROBE_PATH") != "" {
		ffprobePath = os.Getenv("FFPROBE_PATH")
	}

	cmd := exec.Command(
		ffprobePath,
		"-v", "error",
		"-print_format", "json",
		"-show_streams",
		filePath,
	)

	var out bytes.Buffer
	cmd.Stdout = &out

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("ffprobe: %s: %w", stderr.String(), err)
	}

	var data ffprobeOutput
	err = json.Unmarshal(out.Bytes(), &data)
	if err != nil {
		return "", fmt.Errorf("failed to parse ffprobe json: %w", err)
	}

	for _, stream := range data.Streams {
		if stream.CodecType == "video" {
			if stream.Width == 0 || stream.Height == 0 {
				return "other", nil
			}
			ratio := float64(stream.Width) / float64(stream.Height)
			const tolerance = 0.05
			if math.Abs(ratio-(16.0/9.0)) < tolerance {
				return "16:9", nil
			}
			if math.Abs(ratio-(9.0/16.0)) < tolerance {
				return "9:16", nil
			}
			return "other", nil
		}
	}
	return "", fmt.Errorf("no video stream found")
}

func getAspectRatioType(aspectRatio string) string {
	switch aspectRatio {
	case "16:9":
		return "landscape"
	case "9:16":
		return "portrait"
	default:
		return "other"
	}
}
