package pkg

import (
	"fmt"
	"github.com/edjubert/hyprland-ipc-go/hyprctl"
	"github.com/edjubert/hyprland-ipc-go/types"
)

const (
	DefaultMargin = 50
	FromLeft      = "fromLeft"
	FromRight     = "fromRight"
	FromTop       = "fromTop"
	FromBottom    = "fromBottom"
	NoAnimation   = ""
)

type AnimationsOptions struct {
	Animation string
	Margin    int
	Width     string
	Height    string
}

func FromAnimation(client types.HyprlandClient, monitor types.HyprlandMonitor, options AnimationsOptions) error {
	switch options.Animation {
	case FromLeft:
		return fromLeft(client, monitor, options.Margin)
	case FromRight:
		return fromRight(client, monitor, options.Margin)
	case FromTop:
		return fromTop(client, monitor, options.Margin)
	case FromBottom, NoAnimation:
		return fromBottom(client, monitor, options.Margin)
	default:
		return fmt.Errorf("[WARN] - animation unrecognized (%s)", options.Animation)
	}
}
func ToAnimation(client types.HyprlandClient, monitor types.HyprlandMonitor, options AnimationsOptions) error {
	if client.Pid == 0 {
		return fmt.Errorf("[ERROR] - no client")
	}

	switch options.Animation {
	case FromLeft:
		return toLeft(client, monitor)
	case FromRight:
		return toRight(client, monitor)
	case FromTop:
		return toTop(client, monitor)
	case FromBottom:
		return toBottom(client, monitor)
	case NoAnimation:
		return nil
	default:
		return fmt.Errorf("[WARN] - animation unrecognized (%s)", options.Animation)
	}
}

func toTop(client types.HyprlandClient, monitor types.HyprlandMonitor) error {
	x := (monitor.Width-client.Size[0])/2 + monitor.X
	y := monitor.Y - client.Size[1] - DefaultMargin

	dispatch := hyprctl.Dispatch{}
	return dispatch.Move.WindowPixelExact(x, y, client.Address)
}

func fromTop(client types.HyprlandClient, monitor types.HyprlandMonitor, margin int) error {
	x := (monitor.Width-client.Size[0])/2 + monitor.X
	y := monitor.Y + margin

	dispatch := hyprctl.Dispatch{}
	return dispatch.Move.WindowPixelExact(x, y, client.Address)
}

func toLeft(client types.HyprlandClient, monitor types.HyprlandMonitor) error {
	x := monitor.X - monitor.Width - client.Size[0]
	y := monitor.Y + DefaultMargin

	dispatch := hyprctl.Dispatch{}
	err := dispatch.Move.WindowPixelExact(x, y, client.Address)
	return err
}

func fromLeft(client types.HyprlandClient, monitor types.HyprlandMonitor, margin int) error {
	x := (monitor.Width-client.Size[0])/2 + monitor.X
	y := monitor.Y + margin

	dispatch := hyprctl.Dispatch{}
	return dispatch.Move.WindowPixelExact(x, y, client.Address)
}

func toRight(client types.HyprlandClient, monitor types.HyprlandMonitor) error {
	x := monitor.X + monitor.Width
	y := monitor.Y + DefaultMargin

	dispatch := hyprctl.Dispatch{}
	return dispatch.Move.WindowPixelExact(x, y, client.Address)
}

func fromRight(client types.HyprlandClient, monitor types.HyprlandMonitor, margin int) error {
	x := (monitor.Width-client.Size[0])/2 + monitor.X
	y := monitor.Y + margin

	dispatch := hyprctl.Dispatch{}
	return dispatch.Move.WindowPixelExact(x, y, client.Address)
}

func toBottom(client types.HyprlandClient, monitor types.HyprlandMonitor) error {
	x := (monitor.Width-client.Size[0])/2 + monitor.X
	y := monitor.Y + client.Size[1] + monitor.Height

	dispatch := hyprctl.Dispatch{}
	return dispatch.Move.WindowPixelExact(x, y, client.Address)
}

func fromBottom(client types.HyprlandClient, monitor types.HyprlandMonitor, margin int) error {
	x := (monitor.Width-client.Size[0])/2 + monitor.X
	y := monitor.Y + margin

	dispatch := hyprctl.Dispatch{}
	return dispatch.Move.WindowPixelExact(x, y, client.Address)
}
