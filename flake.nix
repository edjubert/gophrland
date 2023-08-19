{
  description = "A set of tools to manage windows, workspaces, monitors and scratchpads on Hyprland";

  inputs = {
    utils.url = "github:numtide/flake-utils";
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.utils.follows = "utils";
    };
  };

  outputs = { self, utils, nixpkgs, gomod2nix, ... }:
    utils.lib.eachDefaultSystem(
      system:
      let
        pkgs = import nixpkgs {
          inherit system;
          overlays = [ gomod2nix.overlays.default ];
        };
      in {
        devShells.default = import ./shell.nix { inherit pkgs; };
        packages.default = pkgs.callPackage ./. {};
      }
    );
}
