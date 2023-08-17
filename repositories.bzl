load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_repositories():
    go_repository(
      name = "com_github_edjubert_hyprland_ipc_go",
      importpath = "github.com/edjubert/hyprland-ipc-go",
      sum = "h1:+gG9YCJ2bLouADW4HiFZWSjQypCUSGozAcds7lk0tmc=",
      version = "v0.0.26",
    )
    go_repository(
      name = "com_github_spf13_cobra",
      importpath = "github.com/spf13/cobra",
      sum = "h1:hyqWnYt1ZQShIddO5kBpj3vu05/++x6tJ6dg8EC572I=",
      version = "v1.7.0",
    )
    go_repository(
      name = "com_gihub_spf13_pflag",
      importpath = "github.com/spf13/pflag",
      sum = "h1:iy+VFUOCP1a+8yFto/drg2CJ5u0yRoB7fZw3DKv/JXA=",
      version = "v1.0.5",
    )
    go_repository(
      name = "com_github_spf13_viper",
      importpath = "github.com/spf13/viper",
      sum = "h1:rGGH0XDZhdUOryiDWjmIvUSWpbNqisK8Wk0Vyefw8hc=",
      version = "v1.16.0",
    )
    go_repository(
      name = "com_github_spf13_afero",
      importpath = "github.com/spf13/afero",
      sum = "h1:stMpOSZFs//0Lv29HduCmli3GUfpFoF3Y1Q/aXj/wVM=",
      version = "v1.9.5",
    )
    go_repository(
        name = "com_github_spf13_jwalterweatherman",
	    importpath = "github.com/spf13/jwalterweatherman",
        sum = "h1:ue6voC5bR5F8YxI5S67j9i582FU4Qvo2bmqnqMYADFk=",
	    version = "v1.1.0",
    )
    go_repository(
      name = "com_github_fsnotify_fsnotify",
	  importpath = "github.com/fsnotify/fsnotify",
      sum = "h1:n+5WquG0fcWoWp6xPWfHdbskMCQaFnG6PfBrh1Ky4HY=",
      version = "v1.6.0",
    )
    go_repository(
      name = "com_github_hashicorp_hcl",
	  importpath = "github.com/hashicorp/hcl",
      sum = "h1:0Anlzjpi4vEasTeNFn2mLJgTSwt0+6sfsiTG8qcWGx4=",
      version = "v1.0.0",
    )
    go_repository(
      name = "com_github_inconshreveable_mousetrap",
	  importpath = "github.com/inconshreaveable/mousetrap",
      sum = "salut",
      version = "v1.1.0",
    )
    go_repository(
      name = "com_github_magiconair_properties",
	  importpath = "github.com/magiconair/properties",
      sum = "h1:IeQXZAiQcpL9mgcAe1Nu6cX9LLw6ExEHKjN0VQdvPDY=",
      version = "v1.8.7",
    )
    go_repository(
      name = "com_github_mitchellh_mapstructure",
	  importpath = "github.com/mitchellh/mapstructure",
      sum = "h1:jeMsZIYE/09sWLaz43PL7Gy6RuMjD2eJVyuac5Z2hdY=",
      version = "v1.5.0",
    )
    go_repository(
      name = "com_github_pelletier_go_toml_v2",
	  importpath = "github.com/pelletier/go-toml/v2",
      sum = "h1:0ctb6s9mE31h0/lhu+J6OPmVeDxJn+kYnJc2jZR9tGQ=",
      version = "v2.0.8",
    )
    go_repository(
      name = "org_golang_x_sys",
	  importpath = "golang.org/x/sys",
      sum = "salut",
      version = "v0.8.0",
    )
    go_repository(
      name = "org_golang_x_text",
	  importpath = "golang.org/x/text",
      sum = "h1:2sjJmO8cDvYveuX97RDLsxlyUxLl+GHoLxBiRdHllBE=",
      version = "v0.9.0",
    )
    go_repository(
      name = "in_gopkg_ini_v1",
	  importpath = "gopkg.in/ini.v1",
      sum = "h1:Dgnx+6+nfE+IfzjUEISNeydPJh9AXNNsWbGP9KzCsOA=",
      version = "v1.67.0",
    )
    go_repository(
      name = "com_github_subosito_gotenv",
	  importpath = "github.com/subosito/gotenv",
      sum = "h1:X1TuBLAMDFbaTAChgCBLu3DU3UPyELpnF2jjJ2cz/S8=",
      version = "v1.4.2",
    )
    go_repository(
      name = "com_github_spf13_cast",
	  importpath = "github.com/spf13/cast",
      sum = "h1:R+kOtfhWQE6TVQzY+4D7wJLBgkdVasCEFxSUBYBYIlA=",
      version = "v1.5.1",
    )
    go_repository(
      name = "org_golang_x_sync",
      importpath = "golang.org/x/sync",
      sum = "h1:uVc8UZUe6tr40fFVnUP5Oj+veunVezqYl9z7DYw9xzw=",
      version = "v0.0.0-20220722155255-886fb9371eb4",
    )
    go_repository(
      name = "in_gopkg_yaml_v3",
      importpath = "gopkg.in/yaml.v3",
      sum = "h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=",
      version = "v3.0.1"
    )
