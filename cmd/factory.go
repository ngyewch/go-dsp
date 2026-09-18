package main

import (
	"context"
	"fmt"
	"math"
	"strings"

	"github.com/madelynnblue/go-dsp/window"
	"github.com/urfave/cli/v3"
)

func getStep(ctx context.Context, cmd *cli.Command) int {
	fftSize := cmd.Int(fftSizeFlag.Name)
	overlap := cmd.Float64(overlapFlag.Name)

	step := int(math.Round(float64(fftSize) * (1 - overlap)))
	if step <= 0 {
		step = 1
	}
	if step > fftSize {
		step = fftSize
	}
	return step
}

func getWindowFunction(ctx context.Context, cmd *cli.Command) (func(int) []float64, error) {
	windowFunctionName := cmd.String(windowFunctionFlag.Name)
	switch strings.ToLower(windowFunctionName) {
	case "":
		return nil, nil
	case "bartlett":
		return window.Bartlett, nil
	case "blackman":
		return window.Blackman, nil
	case "flattop":
		return window.FlatTop, nil
	case "hamming":
		return window.Hamming, nil
	case "hann":
		return window.Hann, nil
	case "rectangular":
		return window.Rectangular, nil
	default:
		return nil, fmt.Errorf("unknown window function: %s", windowFunctionName)
	}
}
