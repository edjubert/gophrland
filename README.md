# Gophrland
Gophrland is a set of tools to manage windows on Hyprland.
This is a rework of [pyprland](https://github.com/hyprland-community/pyprland)

## Installation
### Nix
#### Home manager flake
```nix
{
  inputs = {
    nxipkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    home-manager = {
      url = "github:nix-community/home-manager";
      inputs.nixpkgs.follows = "nixpkgs";
    };
    gophrland = {
      url = "github:edjubert/gophrland";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };
  
  outputs = { nixpkgs, gophrland, ... } @ inputs:
  let
    system = "x86_64-linux";
    pkgs = inputs.nixpkgs.legacyPackages.${system};
  in {
    # Don't forget to change `user` by your username
    homeConfigurations."user" = home-manager.lib.homeManagerConfiguration {
      inherit pkgs;
      
      modules = [
        {
          home.packages = [
            gophrland.packages.${system}.default
          ];
        }
      ];
    };
  };
}
```

### Local
#### Bazel
You can have needed dependencies (like gcc) with nix shell if you have direnv setted up
```bash
direnv allow
```

Then run
```bash
bazel build //cmd/gophrland
```

#### Nix Flake
To install using the `flake.nix` file, just run
```
nix build .
```

The output result is located at `./result/bin/gophrland` which is a symlink to your generated `/nix/store/<hash>/bin/gophrland`

### From source
```bash
git clone https://github.com/edjubert/gophrland
cd gophrland
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
  - "monitors"

options:
  float:
    offset: 0.9
    
  expose:
    name: "mySpecial"

  scratchpads:
    - term:
        command: "alacritty --class gophrland-alacritty"
        unfocus: "hide"
        float: true
        floatOpts:
          animation: "fromTop"
          margin: 60

    - volume:
        command: "alacritty --class pulsemixer-alacritty -e pulsemixer"
        unfocus: "hide"
        float: true
        floatOpts:
          animation: "fromRight"
          margin: 50

    - slack:
        command: "slack"
        float: false
        class: "Slack"

    - whatsdesk:
        command: "whatsdesk"
        float: false
        class: "whatsdesk"

    - cava:
        command: "alacritty --class cava-alacritty -e cava"
        float: true
        floatOpts:
          animation: "fromBottom"
          margin: 10
          width: "100%"
          height: "10%"
```

## Running
You must run the daemon to activate Gophrland
```bash
gophrland daemon --config path/to/your/gophrland.yaml
```

## Plugins
### Scratchpads
![scratchpads](https://github.com/edjubert/gophrland/assets/16240724/b2df1475-5528-40d4-bb54-6078f301dc9c)
- **Name**: `scratchpads`
- **CLI**:
  - **scratchpads**:
    - **toggle**:
      - **[name]**: `gophrland scratchpads toggle [name]` - Show/hide scratchpad
- **Variables**:
  - **Animation**:
    - **Type**: `string`
    - **Values**:
      - fromTop
      - fromBottom
      - fromLeft
      - fromRight
  - **Unfocus**:
    - **Type**: `string`
    - **Values**:
      - hide
- **Options**:
  - **command**: `string` the command to execute
  - **float**: `bool` put the window in floating mode
  - **floatOpts**:
    - **animation**: `Animation` the animation to run (if `float: true`)
    - **margin**:    `int` the margin from the screen side (if `float: true`)
    - **width**:     `string` the client width, in `%` or `px` 
    - **height**:    `string` the client height, in `%` or `px`
  - **unfocus**: `Unfocus` the action when the window is unfocused
  - **class**: *optional* `string` if you want to get the window client by its class (works well for messaging apps such as Slack, Discord or Whatsdesk)

#### Example
```yaml
options:
  scratchpads:
    - term:
        command: "alacritty --class gophrland-alacritty"
        float: true
        floatOpts:
          animation: "fromTop"
          margin: 60
        unfocus: "hide"
        
    - volume:
        command: "alacritty --class pulsemixer-alacritty -e pulsemixer"
        float: false
        floatOpts: # floatOpts won't have any impact as float is false
          animation: "fromRight"
          margin: 50

    - slack:
        command: "slack"
        class: "Slack"
        float: false
```

### Expose
![expose](https://github.com/edjubert/gophrland/assets/16240724/6a37881f-2892-4636-99f6-1093af005275)
- **Name**: `expose`
- **CLI**:
  - **expose**:
    - **toggle**: `gophrland expose toggle` - Move current window to special or current workspace
    - **show**: `gophrland expose show` - Focus expose special workspace
- **Options**:
  - **name**: *optional* `string` the special workspace name

#### Example
```yaml
options:
  expose:
    name: "my_awesome_special_workspace_name"
```

### Float
![float](https://github.com/edjubert/gophrland/assets/16240724/0b490a24-fa25-420f-b812-c0c992d0d42e)
- **Name**: `float`
- **CLI**:
  - **float**:
    - **bring**:
      - **current**: `gophrland float bring current` - Bring all offscreen floating windows to current workspace
- **Options**:
  - **offset**: percentage of window that have to be offscreen to be triggered

#### Example
```yaml
options:
  float:
    offset: 0.1
```

### Monitors
- **Name**: `monitors`
- **CLI**:
  - **float**:
    - **focus**:
      - **next**: `gophrland float focus next` - Focus next monitor
      - **prev**: `gophrland float focus prev` - Focus previous monitor
    - **move**: 
      - **next** : `gophrland move focus next` - Move window to next monitor
      - **prev** : `gophrland move focus prev` - Move window to prev monitor

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
    "github.com/edjubert/hyprland-ipc-go/ipc"
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
