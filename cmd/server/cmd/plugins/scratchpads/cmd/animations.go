package cmd

import (
	"fmt"
	"gophrland/cmd/server/cmd/IPC"
)

const (
	DEFAULT_MARGIN = 50
	FROM_LEFT      = "fromLeft"
	FROM_RIGHT     = "fromRight"
	FROM_TOP       = "fromTop"
	FROM_BOTTOM    = "fromBottom"
	NO_ANIMATION   = ""
)

type AnimationsOptions struct {
	Animation string
	Margin    int
}

func FromAnimation(client IPC.HyprlandClient, monitor IPC.HyprlandMonitor, options AnimationsOptions) error {
	switch options.Animation {
	case FROM_LEFT:
		return fromLeft(client, monitor, options.Margin)
	case FROM_RIGHT:
		return fromRight(client, monitor, options.Margin)
	case FROM_TOP:
		return fromTop(client, monitor, options.Margin)
	case FROM_BOTTOM, NO_ANIMATION:
		return fromBottom(client, monitor, options.Margin)
	default:
		return fmt.Errorf("[WARN] - animation unrecognized (%s)", options.Animation)
	}
}
func ToAnimation(client IPC.HyprlandClient, monitor IPC.HyprlandMonitor, options AnimationsOptions) error {
	if client.Pid == 0 {
		return fmt.Errorf("[ERROR] - no client")
	}

	switch options.Animation {
	case FROM_LEFT:
		return toLeft(client, monitor)
	case FROM_RIGHT:
		return toRight(client, monitor)
	case FROM_TOP:
		return toTop(client, monitor)
	case FROM_BOTTOM:
		return toBottom(client, monitor)
	case NO_ANIMATION:
		return nil
	default:
		return fmt.Errorf("[WARN] - animation unrecognized (%s)", options.Animation)
	}
}

func toTop(client IPC.HyprlandClient, monitor IPC.HyprlandMonitor) error {
	x := (monitor.Width-client.Size[0])/2 + monitor.X
	y := monitor.Y - client.Size[1] - DEFAULT_MARGIN

	return IPC.MoveWindowPixelExact(x, y, client.Address)
}

func fromTop(client IPC.HyprlandClient, monitor IPC.HyprlandMonitor, margin int) error {
	x := (monitor.Width-client.Size[0])/2 + monitor.X
	y := monitor.Y + margin

	return IPC.MoveWindowPixelExact(x, y, client.Address)
}

func toLeft(client IPC.HyprlandClient, monitor IPC.HyprlandMonitor) error {
	x := monitor.X - monitor.Width - client.Size[0]
	y := monitor.Y + DEFAULT_MARGIN

	err := IPC.MoveWindowPixelExact(x, y, client.Address)
	return err
}

func fromLeft(client IPC.HyprlandClient, monitor IPC.HyprlandMonitor, margin int) error {
	x := (monitor.Width-client.Size[0])/2 + monitor.X
	y := monitor.Y + margin

	return IPC.MoveWindowPixelExact(x, y, client.Address)
}

func toRight(client IPC.HyprlandClient, monitor IPC.HyprlandMonitor) error {
	x := monitor.X + monitor.Width
	y := monitor.Y + DEFAULT_MARGIN

	return IPC.MoveWindowPixelExact(x, y, client.Address)
}

func fromRight(client IPC.HyprlandClient, monitor IPC.HyprlandMonitor, margin int) error {
	x := (monitor.Width-client.Size[0])/2 + monitor.X
	y := monitor.Y + margin

	return IPC.MoveWindowPixelExact(x, y, client.Address)
}

func toBottom(client IPC.HyprlandClient, monitor IPC.HyprlandMonitor) error {
	x := (monitor.Width-client.Size[0])/2 + monitor.X
	y := monitor.Y + client.Size[1] + monitor.Height

	return IPC.MoveWindowPixelExact(x, y, client.Address)
}

func fromBottom(client IPC.HyprlandClient, monitor IPC.HyprlandMonitor, margin int) error {
	x := (monitor.Width-client.Size[0])/2 + monitor.X
	y := monitor.Y + margin

	return IPC.MoveWindowPixelExact(x, y, client.Address)
}
