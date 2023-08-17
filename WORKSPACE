load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

#
# nix rules and dependencies
#

# load rules_nixpkgs
http_archive(
    name = "io_tweag_rules_nixpkgs",
    sha256 = "ed30665cdd410deaa837f954a86e5bf38b6e69e9641872d47c1d38e247a0285a",
    strip_prefix = "rules_nixpkgs-3fd605cf4c7cd2b4ec7c72cbec87bf49ad5166ea",
    urls = ["https://github.com/tweag/rules_nixpkgs/archive/3fd605cf4c7cd2b4ec7c72cbec87bf49ad5166ea.tar.gz"],
)

# load everything that rules_nixpkgs rules need to work
load("@io_tweag_rules_nixpkgs//nixpkgs:repositories.bzl", "rules_nixpkgs_dependencies")

rules_nixpkgs_dependencies()

# load and configure nixpkgs
load("@io_tweag_rules_nixpkgs//nixpkgs:nixpkgs.bzl", "nixpkgs_local_repository", "nixpkgs_cc_configure")

nixpkgs_local_repository(
    name = "nixpkgs",
    nix_flake_lock_file = "//:flake.lock",
)

# define a cc toolchain provided by nixpkgs, this toolchain is not used by default
nixpkgs_cc_configure(
    name = "nix_config_cc",
    repository = "@nixpkgs",
)

#
# Go rules
#

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "278b7ff5a826f3dc10f04feaf0b70d48b68748ccd512d7f98bf442077f043fe3",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.41.0/rules_go-v0.41.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.41.0/rules_go-v0.41.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "29218f8e0cebe583643cbf93cae6f971be8a2484cdcfa1e45057658df8d54002",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/0.32.0/bazel-gazelle-0.32.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.32.0/bazel-gazelle-v0.32.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
go_rules_dependencies()

# go toolchain provided by io_bazel_rules_go, this toolchain is used by default
go_register_toolchains(version="host")

#
# Gazelle rules
#

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("//:repositories.bzl", "go_repositories")

# gazelle:repository_macro repositories.bzl%go_repositories
go_repositories()

gazelle_dependencies()
