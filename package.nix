{ buildGoModule, lib, pkgs }:
buildGoModule rec {
  name = "gophrlang-${version}";
  version = "0.0.3.5";

  src = pkgs.fetchFromGitHub {
    owner = "edjubert";
    repo = "gophrland";
    rev = "v${version}";
    hash = "sha256-vkwgmKL7sgdvRwsmfF8myHDg+bDjPL1y9FxEvusWCuU=";
  };

  vendorHash = "sha256-+QaB3SpIowrucG0je4wK6hnzS2D5lZqmVJQyGcsCb04=";

  installPhase = ''
    install -m 755 -D gophrland-${version} $out/bin/gophrland
  '';

  meta = with lib; {
    homepage = "https://github.com/edjubert/gophrland";
    description = "A set of tools to manage windows, workspaces, monitors and scratchpads on Hyprland";
    platforms = platforms.linux;
  };
}
