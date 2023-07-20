package cmd

import (
	"fmt"
	"github.com/edjubert/hyprland-ipc-go/types"
)

func clientIsLost(client types.HyprlandClient, monitor types.HyprlandMonitor, opts FloatOptions) (LostClient, error) {
	offset := OFFSET
	if opts.Offset > 0 && opts.Offset <= 1 {
		offset = opts.Offset
	}

	fmt.Println(opts.Offset, offset)
	clientX := float64(client.At[0])
	clientWidth := float64(client.Size[0])
	clientY := float64(client.At[1])
	clientHeight := float64(client.Size[1])

	offsetX := clientWidth * offset
	offsetY := clientHeight * offset

	monitorX := float64(monitor.X)
	monitorY := float64(monitor.Y)
	monitorWidth := float64(monitor.Width)
	monitorHeight := float64(monitor.Height)

	onLeft := clientX+clientWidth-offsetX < monitorX
	onRight := clientX+clientWidth-offsetX > monitorX+monitorWidth
	onTop := clientY+clientHeight-offsetY < monitorY
	onBottom := clientY > monitorY+monitorHeight-offsetY

	return LostClient{
		Client: client,
		Left:   onLeft,
		Right:  onRight,
		Top:    onTop,
		Bottom: onBottom,
	}, nil
}
