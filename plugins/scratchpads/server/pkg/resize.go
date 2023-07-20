package pkg

import (
	"fmt"
	"github.com/edjubert/hyprland-ipc-go/hyprctl"
	"github.com/edjubert/hyprland-ipc-go/types"
	"strconv"
	"strings"
)

func sizeIsPercent(size string) bool {
	return strings.Contains(size, "%")
}

func sizeIsPx(size string) bool {
	return strings.Contains(size, "px")
}

func convertPercentToPx(size string, monitorSize int) int {
	intSize, err := strconv.Atoi(strings.Split(size, "%")[0])
	if err != nil {
		fmt.Printf("[ERROR] - Could not convert -> %v\n", err)
	}

	clientSize := float64(monitorSize) * (float64(intSize) / 100)

	return int(clientSize)
}

func ResizeClient(client types.HyprlandClient, width, height string) error {
	isValidWidth := width != "" && (sizeIsPercent(width) || sizeIsPx(width))
	isValidHeigth := height != "" && (sizeIsPercent(height) || sizeIsPx(height))

	if !isValidWidth || !isValidHeigth {
		return fmt.Errorf("[ERROR] - Invalid size param")
	}

	getter := hyprctl.Get{}
	monitor, err := getter.MonitorByID(client.Monitor)
	if err != nil {
		return err
	}

	intWidth := 0
	intHeight := 0
	if sizeIsPercent(width) {
		intWidth = convertPercentToPx(width, monitor.Width)
	}
	if sizeIsPercent(height) {
		intHeight = convertPercentToPx(height, monitor.Height)
	}

	if sizeIsPx(width) {
		var err error
		intWidth, err = strconv.Atoi(strings.Split(width, "px")[0])
		if err != nil {
			return err
		}
	}
	if sizeIsPx(height) {
		var err error
		intHeight, err = strconv.Atoi(strings.Split(height, "px")[0])
		if err != nil {
			return err
		}
	}

	dispatch := hyprctl.Dispatch{}
	return dispatch.ResizeWindowExactPixel(client, intWidth, intHeight)
}
