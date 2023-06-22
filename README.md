# Gophrland
Gophrland is a set of tools to manage windows on Hyprland.
This is a rework of [pyprland](https://github.com/hyprland-community/pyprland)

## Installation
```bash
go build -o gophrland ./cmd/gophrland
mv gophrland ~/.local/bin/gophrland
```

## Configuration
The configuration file is a yaml file where you can activate and configure different plugins:
```yaml
plugins:
  - "scratchpads"
  - "expose"
  - "float"

options:
  float:
    offset: 0.9
    
  expose:
    name: "mySpecial"

  scratchpads:
    - term:
        command: "alacritty --class gophrland-alacritty"
        animation: "fromTop"
        unfocus: "hide"
        margin: 60

    - volume:
        command: "alacritty --class pulsemixer-alacritty -e pulsemixer"
        animation: "fromRight"
        unfocus: "hide"
        margin: 50

    - slack:
        command: "slack"
        animation: "fromRight"
        unfocus: "hide"
        class: "Slack"
        margin: 60

    - whatsdesk:
        command: "whatsdesk"
        animation: "fromBottom"
        unfocus: "hide"
        class: "whatsdesk"
        margin: 40

    - cava:
        command: "alacritty --class cava-alacritty -e cava"
        animation: "fromBottom"
        margin: 10
```

## Running
You must run the daemon to activate Gophrland
```bash
gophrland daemon --config path/to/your/gophrland.yaml
```

## Plugins
### Scratchpads
- Name: `scratchpads`
- CLI:
  - toggle: `gophrland scratchpads toggle [name]`
- Options:
  - command: the command to execute
  - animation: the animation to run
  - unfocus: [Unimplemented] the action when the window is unfocus 
  - margin: the margin from the screen side
  - class: [optional] if you want to get the window client by its class (works well for messaging apps such as Slack, Discord or Whatsdesk)

### Expose
- Name: `expose`
- CLI:
  - toggle: `gophrland expose toggle`
  - show special workspace: `gophrland expose show`
- Options:
  - name: the special workspace name

### Float
- Name: `float`
- CLI:
  - toggle: `gophrland float bring current`
- Options:
  - offset: percentage of window that have to be offscreen to be triggered

## Write a plugin
Write a plugin is easy
Create a directory with the name of your plugin containing two subdirectories:
- client (that will contain code for the client CLI)
- server (with all logics that will be applied when receiving command from client)
```bash
mkdir -p ./plugins/myAwesomePlugin/client ./plugins/myAwesomePlugin/server
```

### Create a plugin server side
```go
// ./plugins/myAwesomePlugin/server/cmd/load_plugin.go
package cmd

type MyAwesomePlugin struct {}

func LoadPlugin() {}
```

You just have to add the call in [apply_config.go](./plugins/apply_config.go)
```go
package plugins

import (
  "fmt"
  "github.com/edjubert/gophrland/pkg/server/pkg/config"
  myAwesomePlugin "github.com/edjubert/gophrland/myAwesomePlugin/expose/server/cmd" // Import your plugin here
)

type Options struct {
  MyAwesomePlugin myAwesomePlugin.MyAwesomePluginOptions `yaml:"my_awesome_plugin"` // Add plugin options linked to yaml here
}

type Config struct {
  Plugins []string `yaml:"plugins"`
  Options Options  `yaml:"options"`
}

const (
  MyAwesomePlugin = "myAwesomePlugin" // Declare your plugin name here
)

func ApplyConfig(config config.Config) {
  for _, plugin := range config.Plugins {
    switch plugin {
    case MyAwesomePlugin: 
		myAwesomePlugin.LoadPlugin() // Load your plugin here
    default:
      fmt.Printf("[WARN] - plugin '%s' is not implemented yet\n", plugin)
    }
  }
}
```

You can write a `CallCommand` function under `./plugins/myAwesomePlugin/server/cmd/call_command.go`
```go
package cmd

import "fmt"

const (
	MyCmd = "my-cmd"
)

func Command(cmd string, opts MyAwesomePluginOptions) error {
	switch cmd {
	case MyCmd:
		return myCmd(opts)
	default:
		return fmt.Errorf("[WARN] - unrecognized command")
	}
}
```

And finally update [process_client.go](./pkg/server/internal/call_command.go)


### Create plugin client side
Create an `Run` function

```go
// ./plugins/myAwesomePlugin/client/cmd/myAwesomePlugin.go
package cmd

import (
    "fmt"
    "github.com/edjubert/hyprland-ipc-go"
    "github.com/spf13/cobra"
)

func MyAwesomeFunction(cmd *cobra.Command, args []string) error {
  conn := IPC.StartUnixConnection()

  // If you command take only one argument
  if len(args) != 1 {
    return cmd.Help()
  }

  if _, err := conn.Write([]byte(fmt.Sprintf("scratchpads toggle %s", args[0]))); err != nil {
    panic(err)
  }

  buffer := make([]byte, 1024)
  _, err := conn.Read(buffer)
  if err != nil {
    fmt.Println("[ERROR] - Error reading:", err.Error())
    panic(err)
  }

  return nil
}


func Run() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "myAwesomePlugin",
		Short: "My awesome plugin CLI handler",
		Long:  "These are tools to use the 'my awesome plugin' plugin from the CLI",
		RunE: MyAwesomeFunction,
	}


	return cmd
}
```

And update [the main.go of the plugin](./plugins/main.go)
```go
package plugins

import (
  myAwesomePlugin "github.com/edjubert/gophrland/plugins/myAwesomePlugin/client/cmd"
  "github.com/spf13/cobra"
)

func AddCommand(cmd *cobra.Command) {
  cmd.AddCommand(myAwesomePlugin.Run())
}
```