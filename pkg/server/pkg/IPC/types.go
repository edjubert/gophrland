package IPC

type HyprlandWorkspace struct {
	Id      int    `json:"id"`
	Windows int    `json:"windows"`
	Name    string `json:"name"`
}

type HyprlandMonitor struct {
	Id              int               `json:"id"`
	Name            string            `json:"name"`
	Description     string            `json:"description"`
	Make            string            `json:"make"`
	Model           string            `json:"model"`
	Serial          string            `json:"serial"`
	Width           int               `json:"width"`
	Height          int               `json:"height"`
	RefreshRate     float64           `json:"refreshRate"`
	X               int               `json:"x"`
	Y               int               `json:"y"`
	ActiveWorkspace HyprlandWorkspace `json:"activeWorkspace"`
	Reserved        []int             `json:"reserved"`
	Scale           float64           `json:"scale"`
	Transform       int               `json:"transform"`
	Focused         bool              `json:"focused"`
	DpmsStatus      bool              `json:"dpmsStatus"`
	Vrr             bool              `json:"vrr"`
}

type HyprlandClient struct {
	Address        string            `json:"address,omitempty"`
	Mapped         bool              `json:"mapped,omitempty"`
	Hidden         bool              `json:"hidden,omitempty"`
	At             []int             `json:"at,omitempty"`
	Size           []int             `json:"size,omitempty"`
	Workspace      HyprlandWorkspace `json:"workspace,omitempty"`
	Floating       bool              `json:"floating,omitempty"`
	Monitor        int               `json:"monitor,omitempty"`
	Class          string            `json:"class,omitempty"`
	InitialClass   string            `json:"initialClass,omitempty"`
	Title          string            `json:"title,omitempty"`
	InitialTitle   string            `json:"initialTitle,omitempty"`
	Pid            int               `json:"pid,omitempty"`
	XWayland       bool              `json:"xwayland,omitempty"`
	Pinned         bool              `json:"pinned,omitempty"`
	Fullscreen     bool              `json:"fullscreen,omitempty"`
	FullscreenMode int               `json:"fullscreenMode,omitempty"`
	FakeFullscreen bool              `json:"fakeFullscreen,omitempty"`
}
