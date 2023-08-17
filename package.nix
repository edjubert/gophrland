{ buildGoModule, lib, pkgs }:
buildGoModule rec {
  name = "gophrlang-${version}";
  version = "0.0.3.4";

  src = pkgs.fetchFromGitHub {
    owner = "edjubert";
    repo = "gophrland";
    rev = "v${version}";
    hash = "sha256-4uqf+8E2Ut7RZf2s5qCJFqu0GgxRDK+TJOcDf91buCE=";
  };

  vendorHash = "sha256-+QaB3SpIowrucG0je4wK6hnzS2D5lZqmVJQyGcsCb04=";

  installPhase = ''
    install -m 755 -D gophrland $out/bin/gophrland
  '';

  meta = with lib; {
    homepage = "https://github.com/edjubert/gophrland";
    description = "A set of tools to manage windows, workspaces, monitors and scratchpads on Hyprland";
    platforms = platforms.linux;
  };
}
